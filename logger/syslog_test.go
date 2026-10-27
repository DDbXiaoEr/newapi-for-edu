package logger

import (
	"strings"
	"testing"
	"time"
)

func TestResolveSyslogTarget(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		network     string
		addr        string
		tag         string
		wantNetwork string
		wantAddr    string
		wantTag     string
	}{
		{
			name:    "empty uses local socket defaults",
			wantTag: "newapi",
		},
		{
			name:        "tcp without port gets 514",
			network:     "tcp",
			addr:        "172.16.74.141",
			tag:         "newapi_edu",
			wantNetwork: "tcp",
			wantAddr:    "172.16.74.141:514",
			wantTag:     "newapi_edu",
		},
		{
			name:        "addr without network defaults to udp",
			addr:        "172.16.74.141",
			wantNetwork: "udp",
			wantAddr:    "172.16.74.141:514",
			wantTag:     "newapi",
		},
		{
			name:        "existing port is kept",
			network:     "TCP",
			addr:        "172.16.74.141:1514",
			wantNetwork: "tcp",
			wantAddr:    "172.16.74.141:1514",
			wantTag:     "newapi",
		},
		{
			name:        "ipv6 without port",
			network:     "udp",
			addr:        "2001:db8::1",
			wantNetwork: "udp",
			wantAddr:    "[2001:db8::1]:514",
			wantTag:     "newapi",
		},
		{
			name:        "ipv6 with brackets and no port",
			network:     "tcp",
			addr:        "[2001:db8::1]",
			wantNetwork: "tcp",
			wantAddr:    "[2001:db8::1]:514",
			wantTag:     "newapi",
		},
		{
			name:        "whitespace is trimmed",
			network:     " udp ",
			addr:        " 192.168.1.1 ",
			tag:         "  tagged  ",
			wantNetwork: "udp",
			wantAddr:    "192.168.1.1:514",
			wantTag:     "tagged",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotNetwork, gotAddr, gotTag := resolveSyslogTarget(tt.network, tt.addr, tt.tag)
			if gotNetwork != tt.wantNetwork || gotAddr != tt.wantAddr || gotTag != tt.wantTag {
				t.Fatalf("resolveSyslogTarget(%q, %q, %q) = (%q, %q, %q), want (%q, %q, %q)",
					tt.network, tt.addr, tt.tag, gotNetwork, gotAddr, gotTag, tt.wantNetwork, tt.wantAddr, tt.wantTag)
			}
		})
	}
}

func TestResolveSyslogHostname(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		hostname string
		tag      string
		want     string
	}{
		{name: "empty uses tag", tag: "newapi_edu", want: "newapi_edu"},
		{name: "empty tag falls back", want: "newapi"},
		{name: "explicit hostname wins", hostname: "shared-host", tag: "newapi", want: "shared-host"},
		{name: "whitespace hostname is ignored", hostname: "  ", tag: "app", want: "app"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := resolveSyslogHostname(tt.hostname, tt.tag); got != tt.want {
				t.Fatalf("resolveSyslogHostname(%q, %q) = %q, want %q", tt.hostname, tt.tag, got, tt.want)
			}
		})
	}
}

func TestPrefixSyslogInstance(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		msg      string
		instance string
		hostname string
		want     string
	}{
		{
			name:     "prefixes each line with container identity",
			msg:      "[INFO] hello\n[ERR] boom\n",
			instance: "newapi-0",
			hostname: "newapi",
			want:     "[newapi-0] [INFO] hello\n[newapi-0] [ERR] boom\n",
		},
		{
			name:     "skips prefix when instance matches shared hostname",
			msg:      "[INFO] hello",
			instance: "newapi",
			hostname: "newapi",
			want:     "[INFO] hello",
		},
		{
			name:     "does not double prefix",
			msg:      "[newapi-0] [INFO] hello",
			instance: "newapi-0",
			hostname: "newapi",
			want:     "[newapi-0] [INFO] hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := prefixSyslogInstance(tt.msg, tt.instance, tt.hostname); got != tt.want {
				t.Fatalf("prefixSyslogInstance() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestFormatRemoteSyslogUsesSharedHostname(t *testing.T) {
	t.Parallel()

	ts := time.Date(2026, 10, 27, 8, 30, 0, 0, time.UTC)
	got := formatRemoteSyslog(ts, "newapi", "newapi", 42, "[newapi-abc] [INFO] hello")
	if !strings.Contains(got, " newapi newapi[42]: ") {
		t.Fatalf("shared hostname missing in %q", got)
	}
	if !strings.Contains(got, "[newapi-abc] [INFO] hello") {
		t.Fatalf("instance identity missing in %q", got)
	}
	if strings.Contains(got, " newapi-abc ") {
		t.Fatalf("instance identity must not be used as syslog hostname: %q", got)
	}
}
