package common

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// LogWriterMu protects concurrent access to gin.DefaultWriter/gin.DefaultErrorWriter
// during log file rotation. Acquire RLock when reading/writing through the writers,
// acquire Lock when swapping writers and closing old files.
var LogWriterMu sync.RWMutex

func syslogWriter() io.Writer {
	w := gin.DefaultWriter
	if w == nil {
		return os.Stdout
	}
	return w
}

func syslogErrorWriter() io.Writer {
	w := gin.DefaultErrorWriter
	if w == nil {
		return os.Stderr
	}
	return w
}

func SysLog(s string) {
	t := time.Now()
	LogWriterMu.RLock()
	_, _ = fmt.Fprintf(syslogWriter(), "[SYS] %v | %s \n", t.Format("2006/01/02 - 15:04:05"), s)
	LogWriterMu.RUnlock()
}

func SysError(s string) {
	t := time.Now()
	LogWriterMu.RLock()
	_, _ = fmt.Fprintf(syslogErrorWriter(), "[SYS] %v | %s \n", t.Format("2006/01/02 - 15:04:05"), s)
	LogWriterMu.RUnlock()
}

func FatalLog(v ...any) {
	t := time.Now()
	LogWriterMu.RLock()
	_, _ = fmt.Fprintf(syslogErrorWriter(), "[FATAL] %v | %v \n", t.Format("2006/01/02 - 15:04:05"), v)
	LogWriterMu.RUnlock()
	os.Exit(1)
}

func LogModInitTime(module string, startTime time.Time) {
	duration := time.Since(startTime)
	durationMs := duration.Milliseconds()
	LogWriterMu.RLock()
	defer LogWriterMu.RUnlock()
	_, _ = fmt.Fprintf(syslogWriter(), "[SYS] %v | %-30s %6d ms\n", time.Now().Format("2006/01/02 - 15:04:05"), module, durationMs)
}

func LogStartupSuccess(startTime time.Time, port string) {
	duration := time.Since(startTime)
	durationMs := duration.Milliseconds()

	// Get network IPs
	networkIps := GetNetworkIps()

	LogWriterMu.RLock()
	defer LogWriterMu.RUnlock()

	w := syslogWriter()

	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "  \033[32m%s %s\033[0m  ready in %d ms\n", SystemName, Version, durationMs)
	fmt.Fprintf(w, "\n")

	if !IsRunningInContainer() {
		fmt.Fprintf(w, "  ➜  \033[1mLocal:\033[0m   http://localhost:%s/\n", port)
	}

	for _, ip := range networkIps {
		fmt.Fprintf(w, "  ➜  \033[1mNetwork:\033[0m http://%s:%s/\n", ip, port)
	}

	fmt.Fprintf(w, "\n")
}
