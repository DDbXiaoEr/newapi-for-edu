//go:build windows

package logger

import "io"

func refreshSyslog() {}

func getSyslogWriter() io.Writer { return nil }

func closeSyslog() {}
