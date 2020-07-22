package config

import "github.com/crazy-max/ddns-route53/v2/pkg/utl"

// Route53 holds AWS Route53 data
type Route53 struct {
	HostedZoneID string     `yaml:"hostedZoneID,omitempty" json:"hostedZoneID,omitempty" validate:"required"`
	RecordsSet   RecordsSet `yaml:"recordsSet,omitempty" json:"recordsSet,omitempty" validate:"dive"`
	HandleIPv4   *bool      `yaml:"-" json:"-" label:"-" file:"-"`
	HandleIPv6   *bool      `yaml:"-" json:"-" label:"-" file:"-"`
}

// GetDefaults gets the default values
func (s *Route53) GetDefaults() *Route53 {
	n := &Route53{}
	n.SetDefaults()
	return n
}

// SetDefaults sets the default values
func (s *Route53) SetDefaults() {
	s.HandleIPv4 = utl.NewFalse()
	s.HandleIPv6 = utl.NewFalse()
}
