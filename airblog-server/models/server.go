package models

type Server struct {
	Name string `yaml:"name" json:"name"`
	Port string `yaml:"port" json:"port"`
}
