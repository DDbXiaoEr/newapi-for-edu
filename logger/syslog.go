package logger

import (
	"net"
	"strings"
)

const syslogDefaultPort = "514"

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

func isSyslogIPNetwork(network string) bool {
	switch network {
	case "tcp", "tcp4", "tcp6", "udp", "udp4", "udp6":
		return true
	default:
		return false
	}
}
