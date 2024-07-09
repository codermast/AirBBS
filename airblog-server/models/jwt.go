package models

import "time"

type JWT struct {
	SigningKey  string        `yaml:"signing_key" json:"signing_key"`
	Header      string        `yaml:"header" json:"header"`
	Secret      string        `yaml:"secret" json:"secret"`
	Issuer      string        `yaml:"issuer" json:"issuer"`
	ExpireHours time.Duration `yaml:"expire_hours" json:"expire_hours"`
	Subject     string        `yaml:"subject" json:"subject"`
}
