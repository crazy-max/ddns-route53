package config

import (
	"fmt"
	"strings"
	"testing"

	"github.com/crazy-max/ddns-route53/v2/pkg/utl"
	"github.com/crazy-max/gonfig/env"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadFile(t *testing.T) {
	cases := []struct {
		name     string
		cli      Cli
		wantData *Config
		wantErr  bool
	}{
		{
			name:    "Failed on non-existing file",
			wantErr: true,
		},
		{
			name: "Fail on wrong file format",
			cli: Cli{
				Cfgfile: "./fixtures/config.invalid.yml",
			},
			wantErr: true,
		},
		{
			name: "Success",
			cli: Cli{
				Cfgfile: "./fixtures/config.test.yml",
			},
			wantData: &Config{
				Cli: Cli{
					Cfgfile: "./fixtures/config.test.yml",
				},
				Credentials: &Credentials{
					AccessKeyID:     "ABCDEFGHIJKLMNO123456",
					SecretAccessKey: "abcdefgh123456IJKLMN+OPQRS7890+ABCDEFGH",
				},
				Route53: &Route53{
					HostedZoneID: "ABCEEFG123456789",
					RecordsSet: RecordsSet{
						RecordSet{
							Name: "ddns.example.com.",
							Type: "A",
							TTL:  300,
						},
						RecordSet{
							Name: "ddns.example.com.",
							Type: "AAAA",
							TTL:  300,
						},
						RecordSet{
							Name: "another.example2.com.",
							Type: "A",
							TTL:  600,
						},
					},
					HandleIPv4: utl.NewTrue(),
					HandleIPv6: utl.NewTrue(),
				},
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := Load(tt.cli)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantData, cfg)
			if cfg != nil {
				assert.NotEmpty(t, cfg.String())
			}
		})
	}
}

func TestLoadEnv(t *testing.T) {
	testCases := []struct {
		desc     string
		cli      Cli
		environ  []string
		expected interface{}
		wantErr  bool
	}{
		{
			desc:     "no env vars",
			environ:  nil,
			expected: nil,
			wantErr:  true,
		},
		{
			desc: "one record set",
			environ: []string{
				"DDNSR53_CREDENTIALS_ACCESSKEYIDFILE=./fixtures/run_secrets_akid",
				"DDNSR53_CREDENTIALS_SECRETACCESSKEYFILE=./fixtures/run_secrets_sak",
				"DDNSR53_ROUTE53_HOSTEDZONEID=ABCEEFG123456789",
				"DDNSR53_ROUTE53_RECORDSSET_0_NAME=ddns.example.com.",
				"DDNSR53_ROUTE53_RECORDSSET_0_TYPE=A",
				"DDNSR53_ROUTE53_RECORDSSET_0_TTL=300",
			},
			expected: &Config{
				Credentials: &Credentials{
					AccessKeyIDFile:     "./fixtures/run_secrets_akid",
					SecretAccessKeyFile: "./fixtures/run_secrets_sak",
				},
				Route53: &Route53{
					HostedZoneID: "ABCEEFG123456789",
					RecordsSet: RecordsSet{
						RecordSet{
							Name: "ddns.example.com.",
							Type: "A",
							TTL:  300,
						},
					},
					HandleIPv4: utl.NewTrue(),
					HandleIPv6: utl.NewFalse(),
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			if tt.environ != nil {
				for _, environ := range tt.environ {
					n := strings.SplitN(environ, "=", 2)
					t.Setenv(n[0], n[1])
				}
			}

			cfg, err := Load(tt.cli)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expected, cfg)
		})
	}
}

func TestValidation(t *testing.T) {
	cases := []struct {
		name string
		cli  Cli
	}{
		{
			name: "Success",
			cli: Cli{
				Cfgfile: "./fixtures/config.test.yml",
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := Load(tt.cli)
			require.NoError(t, err)

			dec, err := env.Encode("DDNSR53_", cfg)
			require.NoError(t, err)
			for _, value := range dec {
				fmt.Printf(`%s=%s\n`, value.Name, value.Default)
			}
		})
	}
}
