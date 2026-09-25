package testdb

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/QuantumNous/new-api/common"
	sqlmysql "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Options struct {
	Models  []any
	MaxOpen int
}

func Open(t testing.TB, opts ...Options) *gorm.DB {
	t.Helper()
	opt := options(opts)
	kind, dsn := resolve(t)
	db, _ := openIsolated(t, kind, dsn)
	if opt.MaxOpen > 0 {
		sqlDB, err := db.DB()
		require.NoError(t, err)
		sqlDB.SetMaxOpenConns(opt.MaxOpen)
	}
	if len(opt.Models) > 0 {
		require.NoError(t, db.AutoMigrate(opt.Models...))
	}
	return db
}

func OpenBound(t testing.TB, main, log **gorm.DB, initCols func(), opts ...Options) *gorm.DB {
	t.Helper()
	db := Open(t, opts...)
	Bind(t, db, main, log, initCols)
	return db
}

func OpenMain(models ...any) (*gorm.DB, common.DatabaseType, func()) {
	kind, dsn, err := resolveErr()
	if err != nil {
		panic(err)
	}
	db, isolatedDSN, err := openIsolatedErr(kind, dsn)
	if err != nil {
		panic(fmt.Errorf("testdb: open isolated %s database %s: %w", kind, isolatedDSN, err))
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(1)
	if len(models) > 0 {
		if err := db.AutoMigrate(models...); err != nil {
			panic("testdb: migrate: " + err.Error())
		}
	}
	return db, databaseType(kind), func() { _ = sqlDB.Close() }
}

func Bind(t testing.TB, db *gorm.DB, main, log **gorm.DB, initCols func()) {
	t.Helper()
	previousMain, previousLog := *main, *log
	previousMainType, previousLogType := common.MainDatabaseType(), common.LogDatabaseType()
	typ := Type(t)
	*main, *log = db, db
	common.SetDatabaseTypes(typ, typ)
	if initCols != nil {
		initCols()
	}
	t.Cleanup(func() {
		*main, *log = previousMain, previousLog
		common.SetDatabaseTypes(previousMainType, previousLogType)
		if initCols != nil {
			initCols()
		}
	})
}

func Dialector(t testing.TB) (gorm.Dialector, common.DatabaseType) {
	t.Helper()
	kind, dsn := resolve(t)
	_, isolatedDSN := openIsolated(t, kind, dsn)
	if kind == "mysql" {
		return mysql.Open(isolatedDSN), common.DatabaseTypeMySQL
	}
	return postgres.New(postgres.Config{DSN: isolatedDSN, PreferSimpleProtocol: true}), common.DatabaseTypePostgreSQL
}

func Kind(t testing.TB) string {
	t.Helper()
	kind, _ := resolve(t)
	return kind
}

func Type(t testing.TB) common.DatabaseType {
	t.Helper()
	return databaseType(Kind(t))
}

func VersionQuery(kind string) string {
	if kind == "mysql" {
		return "SELECT VERSION()"
	}
	return "SELECT version()"
}

func resolve(t testing.TB) (string, string) {
	t.Helper()
	kind, dsn, err := resolveErr()
	if err != nil {
		t.Skip(err.Error())
	}
	return kind, dsn
}

func resolveErr() (string, string, error) {
	loadDotEnv()
	kind := strings.ToLower(strings.TrimSpace(os.Getenv("TEST_DB_DIALECT")))
	if kind == "sqlite" {
		return "", "", fmt.Errorf("SQLite is not supported; set TEST_DB_DIALECT to mysql or postgres")
	}
	if kind == "" {
		if dsn := firstEnv("TEST_POSTGRES_DSN", "SQL_DSN"); looksLikePostgres(dsn) {
			return "postgres", dsn, nil
		}
		if dsn := firstEnv("TEST_MYSQL_DSN"); dsn != "" {
			return "mysql", dsn, nil
		}
		if dsn := strings.TrimSpace(os.Getenv("SQL_DSN")); dsn != "" && !looksLikePostgres(dsn) && !looksLikeClickHouse(dsn) {
			return "mysql", dsn, nil
		}
		return "", "", fmt.Errorf("set TEST_POSTGRES_DSN, TEST_MYSQL_DSN, or SQL_DSN to a MySQL or PostgreSQL database")
	}
	switch kind {
	case "mysql", "postgres", "postgresql":
		if kind == "postgresql" {
			kind = "postgres"
		}
		env := "TEST_MYSQL_DSN"
		if kind == "postgres" {
			env = "TEST_POSTGRES_DSN"
		}
		dsn := firstEnv(env, "SQL_DSN")
		if dsn == "" {
			return "", "", fmt.Errorf("%s or SQL_DSN is not configured", env)
		}
		return kind, dsn, nil
	default:
		return "", "", fmt.Errorf("unsupported TEST_DB_DIALECT %q; use mysql or postgres", kind)
	}
}

func openIsolated(t testing.TB, kind, dsn string) (*gorm.DB, string) {
	t.Helper()
	db, isolatedDSN, err := openIsolatedErr(kind, dsn)
	require.NoError(t, err)
	t.Logf("isolated database: %s (%s)", isolatedDSN, kind)
	t.Cleanup(func() {
		if sqlDB, err := db.DB(); err == nil {
			_ = sqlDB.Close()
		}
	})
	return db, isolatedDSN
}

func openIsolatedErr(kind, dsn string) (*gorm.DB, string, error) {
	name := fmt.Sprintf("newapi_test_%d", time.Now().UnixNano())
	var original, isolated gorm.Dialector
	var isolatedDSN string
	switch kind {
	case "mysql":
		config, err := sqlmysql.ParseDSN(dsn)
		if err != nil {
			return nil, "", err
		}
		if config.Net != "tcp" {
			return nil, "", fmt.Errorf("mysql tests require tcp DSN")
		}
		host, _, err := net.SplitHostPort(config.Addr)
		if err != nil {
			return nil, "", err
		}
		if !isLoopbackHost(host) {
			return nil, "", fmt.Errorf("database tests only permit loopback instances")
		}
		original = mysql.Open(dsn)
		config.DBName = name
		isolatedDSN = config.FormatDSN()
		isolated = mysql.Open(isolatedDSN)
	case "postgres":
		parsed, err := url.Parse(dsn)
		if err != nil {
			return nil, "", err
		}
		if !isLoopbackHost(parsed.Hostname()) {
			return nil, "", fmt.Errorf("database tests only permit loopback instances")
		}
		parsed.Path = "/" + name
		isolatedDSN = parsed.String()
		original = postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true})
		isolated = postgres.New(postgres.Config{DSN: isolatedDSN, PreferSimpleProtocol: true})
	default:
		return nil, "", fmt.Errorf("unsupported test dialect %q", kind)
	}
	admin, err := gorm.Open(original, &gorm.Config{})
	if err != nil {
		return nil, "", err
	}
	createSQL := "CREATE DATABASE " + name
	if kind == "mysql" {
		createSQL += " CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"
	}
	if err := admin.Exec(createSQL).Error; err != nil {
		return nil, "", err
	}
	sqlDB, err := admin.DB()
	if err != nil {
		return nil, "", err
	}
	_ = sqlDB.Close()
	db, err := gorm.Open(isolated, &gorm.Config{})
	if err != nil {
		return nil, "", err
	}
	return db, isolatedDSN, nil
}

func options(opts []Options) Options {
	if len(opts) == 0 {
		return Options{}
	}
	return opts[0]
}

func databaseType(kind string) common.DatabaseType {
	if kind == "mysql" {
		return common.DatabaseTypeMySQL
	}
	return common.DatabaseTypePostgreSQL
}

func loadDotEnv() {
	dir, err := os.Getwd()
	if err != nil {
		return
	}
	for {
		path := filepath.Join(dir, ".env")
		if _, err := os.Stat(path); err == nil {
			_ = godotenv.Load(path)
			return
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return
		}
		dir = parent
	}
}

func firstEnv(names ...string) string {
	for _, name := range names {
		if dsn := strings.TrimSpace(os.Getenv(name)); dsn != "" {
			return dsn
		}
	}
	return ""
}

func looksLikePostgres(dsn string) bool {
	return strings.HasPrefix(dsn, "postgres://") || strings.HasPrefix(dsn, "postgresql://")
}

func looksLikeClickHouse(dsn string) bool {
	return strings.HasPrefix(dsn, "clickhouse://") || strings.HasPrefix(dsn, "tcp://") ||
		strings.HasPrefix(dsn, "http://") || strings.HasPrefix(dsn, "https://")
}

func isLoopbackHost(host string) bool {
	switch strings.ToLower(host) {
	case "localhost", "127.0.0.1", "::1":
		return true
	}
	ip := net.ParseIP(host)
	return ip != nil && ip.IsLoopback()
}
