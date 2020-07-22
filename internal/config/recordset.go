package config

// RecordsSet holds slice of record set configuration
type RecordsSet []RecordSet

// RecordSet holds data necessary for record set configuration
type RecordSet struct {
	Name string `yaml:"name,omitempty" json:"name,omitempty" validate:"required"`
	Type string `yaml:"type,omitempty" json:"type,omitempty" validate:"required,oneof=A AAAA"`
	TTL  int64  `yaml:"ttl,omitempty" json:"ttl,omitempty" validate:"required,min=1"`
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
