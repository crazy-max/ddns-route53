package config

import r53types "github.com/aws/aws-sdk-go-v2/service/route53/types"

// RecordsSet holds slice of record set configuration
type RecordsSet []RecordSet

// RecordSet holds data necessary for record set configuration
type RecordSet struct {
	Name     string          `yaml:"name,omitempty" json:"name,omitempty" validate:"required"`
	Type     r53types.RRType `yaml:"type,omitempty" json:"type,omitempty" validate:"required,oneof=A AAAA SRV"`
	TTL      int64           `yaml:"ttl,omitempty" json:"ttl,omitempty" validate:"required,min=1"`
	Priority int             `yaml:"priority,omitempty" json:"priority,omitempty" validate:"min=0,max=65535"`
	Weight   int             `yaml:"weight,omitempty" json:"weight,omitempty" validate:"min=0,max=65535"`
	Port     int             `yaml:"port,omitempty" json:"port,omitempty" validate:"required_if=Type SRV,min=0,max=65535"`
	Target   string          `yaml:"target,omitempty" json:"target,omitempty" validate:"required_if=Type SRV"`
}

// GetDefaults gets the default values
func (s *RecordSet) GetDefaults() *RecordSet {
	n := &RecordSet{}
	n.SetDefaults()
	return n
}

// SetDefaults sets the default values
func (s *RecordSet) SetDefaults() {
	// noop
}
