//go:build !windows

package logger

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/QuantumNous/new-api/common"
)

type syslogClient struct {
	mu       sync.Mutex
	conn     net.Conn
	network  string
	addr     string
	tag      string
	hostname string
	instance string
	local    bool
	pid      int
}

var syslogUnsafeWriter *syslogClient
var syslogConnectedKey string

const syslogDialTimeout = 3 * time.Second

func refreshSyslog() {
	if !common.SyslogEnabled {
		closeSyslog()
		syslogConnectedKey = syslogConfigKey()
		return
	}

	network, addr, tag := resolveSyslogTarget(common.SyslogNetwork, common.SyslogAddr, common.SyslogTag)
	hostname := resolveSyslogHostname(common.SyslogHostname, tag)
	key := syslogConfigKey()
	if syslogUnsafeWriter != nil && syslogConnectedKey == key {
		return
	}

	closeSyslog()

	client := &syslogClient{
		network:  network,
		addr:     addr,
		tag:      tag,
		hostname: hostname,
		instance: syslogInstanceID(),
		local:    network == "",
		pid:      os.Getpid(),
	}
	if err := client.connect(); err != nil {
		log.Printf("[SYS] syslog connect failed (network=%s addr=%s tag=%s hostname=%s): %v", network, addr, tag, hostname, err)
		syslogUnsafeWriter = nil
		syslogConnectedKey = ""
		return
	}
	syslogUnsafeWriter = client
	syslogConnectedKey = key
	log.Printf("[SYS] syslog connected (network=%s addr=%s tag=%s hostname=%s instance=%s)", network, addr, tag, hostname, client.instance)
}

func syslogConfigKey() string {
	network, addr, tag := resolveSyslogTarget(common.SyslogNetwork, common.SyslogAddr, common.SyslogTag)
	hostname := resolveSyslogHostname(common.SyslogHostname, tag)
	return fmt.Sprintf("%t|%s|%s|%s|%s", common.SyslogEnabled, network, addr, tag, hostname)
}

func getSyslogWriter() io.Writer {
	if syslogUnsafeWriter == nil {
		return nil
	}
	return syslogUnsafeWriter
}

func closeSyslog() {
	if syslogUnsafeWriter != nil {
		_ = syslogUnsafeWriter.Close()
		syslogUnsafeWriter = nil
	}
	syslogConnectedKey = ""
}

func (w *syslogClient) Write(p []byte) (int, error) {
	if w == nil {
		return 0, io.ErrClosedPipe
	}
	msg := prefixSyslogInstance(string(p), w.instance, w.hostname)
	lines := strings.Split(strings.ReplaceAll(msg, "\r\n", "\n"), "\n")
	w.mu.Lock()
	defer w.mu.Unlock()
	for _, line := range lines {
		if line == "" {
			continue
		}
		if err := w.writeAndRetry(line); err != nil {
			return 0, err
		}
	}
	return len(p), nil
}

func (w *syslogClient) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.closeLocked()
}

func (w *syslogClient) writeAndRetry(line string) error {
	if w.conn != nil {
		if err := w.writeLine(line); err == nil {
			return nil
		}
	}
	if err := w.connectLocked(); err != nil {
		return err
	}
	return w.writeLine(line)
}

func (w *syslogClient) writeLine(line string) error {
	var payload string
	if w.local {
		payload = formatLocalSyslog(time.Now(), w.tag, w.pid, line)
	} else {
		payload = formatRemoteSyslog(time.Now(), w.hostname, w.tag, w.pid, line)
	}
	_, err := io.WriteString(w.conn, payload)
	return err
}

func (w *syslogClient) connect() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.connectLocked()
}

func (w *syslogClient) connectLocked() error {
	_ = w.closeLocked()
	if w.network == "" {
		conn, err := dialLocalSyslog()
		if err != nil {
			return err
		}
		w.conn = conn
		w.local = true
		return nil
	}
	if w.addr == "" {
		return fmt.Errorf("syslog address is empty")
	}
	conn, err := dialSyslogConn(w.network, w.addr)
	if err != nil {
		return err
	}
	w.conn = conn
	w.local = false
	return nil
}

func (w *syslogClient) closeLocked() error {
	if w.conn == nil {
		return nil
	}
	err := w.conn.Close()
	w.conn = nil
	return err
}

func dialSyslogConn(network, addr string) (net.Conn, error) {
	d := net.Dialer{Timeout: syslogDialTimeout}
	return d.Dial(network, addr)
}

func dialLocalSyslog() (net.Conn, error) {
	networks := []string{"unixgram", "unix"}
	paths := []string{"/dev/log", "/var/run/syslog", "/var/run/log"}
	var lastErr error
	for _, network := range networks {
		for _, path := range paths {
			conn, err := dialSyslogConn(network, path)
			if err == nil {
				return conn, nil
			}
			lastErr = err
		}
	}
	if lastErr == nil {
		lastErr = fmt.Errorf("unix syslog delivery error")
	}
	return nil, lastErr
}
