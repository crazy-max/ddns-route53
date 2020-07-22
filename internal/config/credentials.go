package config

// Credentials holds data necessary for AWS configuration
type Credentials struct {
	AccessKeyID         string `yaml:"accessKeyID,omitempty" json:"accessKeyID,omitempty" validate:"omitempty"`
	AccessKeyIDFile     string `yaml:"accessKeyIDFile,omitempty" json:"accessKeyIDFile,omitempty" validate:"omitempty,file"`
	SecretAccessKey     string `yaml:"secretAccessKey,omitempty" json:"secretAccessKey,omitempty" validate:"omitempty"`
	SecretAccessKeyFile string `yaml:"secretAccessKeyFile,omitempty" json:"secretAccessKeyFile,omitempty" validate:"omitempty,file"`
}

// GetDefaults gets the default values
func (s *Credentials) GetDefaults() *Credentials {
	n := &Credentials{}
	n.SetDefaults()
	return n
}

// SetDefaults sets the default values
func (s *Credentials) SetDefaults() {
	// noop
}
