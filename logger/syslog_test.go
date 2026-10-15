package logger

import "testing"

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
