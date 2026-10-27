package logger

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/QuantumNous/new-api/common"
)

const syslogDefaultPort = "514"
const syslogPriority = 6

func resolveSyslogTarget(network, addr, tag string) (string, string, string) {
	network = strings.ToLower(strings.TrimSpace(network))
	addr = strings.TrimSpace(addr)
	tag = strings.TrimSpace(tag)
	if tag == "" {
		tag = "newapi"
	}
	if network == "" && addr != "" {
		network = "udp"
	}
	if isSyslogIPNetwork(network) && addr != "" {
		if _, _, err := net.SplitHostPort(addr); err != nil {
			addr = net.JoinHostPort(strings.Trim(addr, "[]"), syslogDefaultPort)
		}
	}
	return network, addr, tag
}

func resolveSyslogHostname(hostname, tag string) string {
	hostname = strings.TrimSpace(hostname)
	if hostname != "" {
		return hostname
	}
	tag = strings.TrimSpace(tag)
	if tag == "" {
		return "newapi"
	}
	return tag
}

func syslogInstanceID() string {
	for _, v := range []string{os.Getenv("POD_NAME"), os.Getenv("HOSTNAME")} {
		if s := strings.TrimSpace(v); s != "" {
			return s
		}
	}
	if s := strings.TrimSpace(common.NodeName); s != "" {
		return s
	}
	h, _ := os.Hostname()
	return strings.TrimSpace(h)
}

func prefixSyslogInstance(msg, instance, hostname string) string {
	instance = strings.TrimSpace(instance)
	if instance == "" || instance == strings.TrimSpace(hostname) {
		return msg
	}
	prefix := "[" + instance + "] "
	lines := strings.Split(msg, "\n")
	for i, line := range lines {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, prefix) {
			continue
		}
		lines[i] = prefix + line
	}
	return strings.Join(lines, "\n")
}

func formatRemoteSyslog(ts time.Time, hostname, tag string, pid int, msg string) string {
	nl := ""
	if !strings.HasSuffix(msg, "\n") {
		nl = "\n"
	}
	return fmt.Sprintf("<%d>%s %s %s[%d]: %s%s",
		syslogPriority, ts.Format(time.RFC3339), hostname, tag, pid, msg, nl)
}

func formatLocalSyslog(ts time.Time, tag string, pid int, msg string) string {
	nl := ""
	if !strings.HasSuffix(msg, "\n") {
		nl = "\n"
	}
	return fmt.Sprintf("<%d>%s %s[%d]: %s%s",
		syslogPriority, ts.Format(time.Stamp), tag, pid, msg, nl)
}

func isSyslogIPNetwork(network string) bool {
	switch network {
	case "tcp", "tcp4", "tcp6", "udp", "udp4", "udp6":
		return true
	default:
		return false
	}
}
