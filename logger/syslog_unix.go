//go:build !windows

package logger

import (
	"fmt"
	"io"
	"log"
	"log/syslog"
	"time"

	"github.com/QuantumNous/new-api/common"
)

var syslogUnsafeWriter *syslog.Writer
var syslogConnectedKey string

const syslogDialTimeout = 3 * time.Second

func refreshSyslog() {
	if !common.SyslogEnabled {
		closeSyslog()
		syslogConnectedKey = syslogConfigKey()
		return
	}

	network, addr, tag := resolveSyslogTarget(common.SyslogNetwork, common.SyslogAddr, common.SyslogTag)
	key := syslogConfigKey()
	if syslogUnsafeWriter != nil && syslogConnectedKey == key {
		return
	}

	closeSyslog()

	var (
		w   *syslog.Writer
		err error
	)
	if network != "" {
		if addr == "" {
			err = fmt.Errorf("syslog address is empty")
		} else {
			w, err = dialSyslogWithTimeout(network, addr, tag)
		}
	} else {
		w, err = syslog.New(syslog.LOG_INFO, tag)
	}
	if err != nil {
		log.Printf("[SYS] syslog connect failed (network=%s addr=%s tag=%s): %v", network, addr, tag, err)
		syslogUnsafeWriter = nil
		syslogConnectedKey = ""
		return
	}
	syslogUnsafeWriter = w
	syslogConnectedKey = key
	log.Printf("[SYS] syslog connected (network=%s addr=%s tag=%s)", network, addr, tag)
}

func syslogConfigKey() string {
	network, addr, tag := resolveSyslogTarget(common.SyslogNetwork, common.SyslogAddr, common.SyslogTag)
	return fmt.Sprintf("%t|%s|%s|%s", common.SyslogEnabled, network, addr, tag)
}

func dialSyslogWithTimeout(network, addr, tag string) (*syslog.Writer, error) {
	type dialResult struct {
		w   *syslog.Writer
		err error
	}
	done := make(chan dialResult, 1)
	go func() {
		w, err := syslog.Dial(network, addr, syslog.LOG_INFO, tag)
		done <- dialResult{w: w, err: err}
	}()
	select {
	case res := <-done:
		return res.w, res.err
	case <-time.After(syslogDialTimeout):
		go func() {
			res := <-done
			if res.w != nil {
				_ = res.w.Close()
			}
		}()
		return nil, fmt.Errorf("syslog dial to %s://%s timed out after %s", network, addr, syslogDialTimeout)
	}
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
