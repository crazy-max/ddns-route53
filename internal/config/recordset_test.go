package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateSRVRecordSet(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		priority int
		weight   int
		port     int
		target   string
		wantErr  bool
	}{
		{name: "zero priority and weight", port: 2032, target: "host1.example.org."},
		{name: "maximum values", priority: 65535, weight: 65535, port: 65535, target: "host1.example.org."},
		{name: "missing port", target: "host1.example.org.", wantErr: true},
		{name: "missing target", port: 2032, wantErr: true},
		{name: "negative priority", priority: -1, port: 2032, target: "host1.example.org.", wantErr: true},
		{name: "priority overflow", priority: 65536, port: 2032, target: "host1.example.org.", wantErr: true},
		{name: "negative weight", weight: -1, port: 2032, target: "host1.example.org.", wantErr: true},
		{name: "weight overflow", weight: 65536, port: 2032, target: "host1.example.org.", wantErr: true},
		{name: "negative port", port: -1, target: "host1.example.org.", wantErr: true},
		{name: "port overflow", port: 65536, target: "host1.example.org.", wantErr: true},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			cfg := &Config{
				Route53: &Route53{
					HostedZoneID: "ZTEST123",
					RecordsSet: RecordsSet{{
						Name:     "_ssh._tcp.host1.example.org.",
						Type:     "SRV",
						TTL:      300,
						Priority: tt.priority,
						Weight:   tt.weight,
						Port:     tt.port,
						Target:   tt.target,
					}},
				},
			}
			if tt.wantErr {
				require.Error(t, cfg.validate())
			} else {
				require.NoError(t, cfg.validate())
			}
		})
	}
}
