package config

import (
	"fmt"
	"os"
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
				Ipprovider: "identme",
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
	defer UnsetEnv("DDNSR53_")

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
				"DDNSR53_IPPROVIDER=ipify",
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
				Ipprovider: "ipify",
			},
			wantErr: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			UnsetEnv("DDNSR53_")

			if tt.environ != nil {
				for _, environ := range tt.environ {
					n := strings.SplitN(environ, "=", 2)
					os.Setenv(n[0], n[1])
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
				fmt.Println(fmt.Sprintf(`%s=%s`, value.Name, value.Default))
			}
		})
	}
}

func UnsetEnv(prefix string) (restore func()) {
	before := map[string]string{}

	for _, e := range os.Environ() {
		if !strings.HasPrefix(e, prefix) {
			continue
		}

		parts := strings.SplitN(e, "=", 2)
		before[parts[0]] = parts[1]

		os.Unsetenv(parts[0])
	}

	return func() {
		after := map[string]string{}

		for _, e := range os.Environ() {
			if !strings.HasPrefix(e, prefix) {
				continue
			}

			parts := strings.SplitN(e, "=", 2)
			after[parts[0]] = parts[1]

			// Check if the envar previously existed
			v, ok := before[parts[0]]
			if !ok {
				// This is a newly added envar with prefix, zap it
				os.Unsetenv(parts[0])
				continue
			}

			if parts[1] != v {
				// If the envar value has changed, set it back
				os.Setenv(parts[0], v)
			}
		}

		// Still need to check if there have been any deleted envars
		for k, v := range before {
			if _, ok := after[k]; !ok {
				// k is not present in after, so we set it.
				os.Setenv(k, v)
			}
		}
	}
}
