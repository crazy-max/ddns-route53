package config

// WanIP holds WAN IP source configuration.
type WanIP struct {
	Providers *WanIPProviders `yaml:"providers,omitempty" json:"providers,omitempty" validate:"omitempty"`
}

// WanIPProviders holds configurable WAN IP lookup providers.
type WanIPProviders struct {
	IPv4 []string `yaml:"ipv4,omitempty" json:"ipv4,omitempty" validate:"omitempty,dive,required"`
	IPv6 []string `yaml:"ipv6,omitempty" json:"ipv6,omitempty" validate:"omitempty,dive,required"`
}

// GetDefaults gets the default values.
func (s *WanIP) GetDefaults() *WanIP {
	n := &WanIP{}
	n.SetDefaults()
	return n
}

// SetDefaults sets the default values.
func (s *WanIP) SetDefaults() {
	// noop
}
