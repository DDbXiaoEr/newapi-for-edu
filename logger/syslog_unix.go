//go:build !windows

package logger

import (
	"io"
	"log/syslog"

	"github.com/QuantumNous/new-api/common"
)

var syslogUnsafeWriter *syslog.Writer

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
		syslogUnsafeWriter, err = syslog.Dial(common.SyslogNetwork, common.SyslogAddr, syslog.LOG_INFO, tag)
	} else {
		syslogUnsafeWriter, err = syslog.New(syslog.LOG_INFO, tag)
	}
	if err != nil {
		syslogUnsafeWriter = nil
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
