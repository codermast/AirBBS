package models

type Server struct {
	Name string `yaml:"name" json:"name"`
	Host string `yaml:"host" json:"host"`
	Port string `yaml:"port" json:"port"`
}
