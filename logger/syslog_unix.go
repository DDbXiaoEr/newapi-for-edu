//go:build !windows

package logger

import (
	"fmt"
	"io"
	"log/syslog"
	"time"

	"github.com/QuantumNous/new-api/common"
)

var syslogUnsafeWriter *syslog.Writer

// syslogDialTimeout 限制 syslog 连接耗时，避免目标地址不可达时阻塞启动/配置同步
const syslogDialTimeout = 3 * time.Second

func refreshSyslog() {
	if syslogUnsafeWriter != nil {
		syslogUnsafeWriter.Close()
		syslogUnsafeWriter = nil
	}
	if !common.SyslogEnabled {
		return
	}
	var err error
	tag := common.SyslogTag
	if tag == "" {
		tag = "newapi"
	}
	if common.SyslogNetwork != "" {
		syslogUnsafeWriter, err = dialSyslogWithTimeout(common.SyslogNetwork, common.SyslogAddr, tag)
	} else {
		syslogUnsafeWriter, err = syslog.New(syslog.LOG_INFO, tag)
	}
	if err != nil {
		syslogUnsafeWriter = nil
	}
}

// dialSyslogWithTimeout 带超时的 syslog.Dial：超过 syslogDialTimeout 未连接成功即放弃
func dialSyslogWithTimeout(network, addr, tag string) (*syslog.Writer, error) {
	var (
		w   *syslog.Writer
		err error
	)
	done := make(chan struct{})
	go func() {
		defer close(done)
		w, err = syslog.Dial(network, addr, syslog.LOG_INFO, tag)
	}()
	select {
	case <-done:
		return w, err
	case <-time.After(syslogDialTimeout):
		return nil, fmt.Errorf("syslog dial to %s timed out after %s", addr, syslogDialTimeout)
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
		syslogUnsafeWriter.Close()
		syslogUnsafeWriter = nil
	}
}
