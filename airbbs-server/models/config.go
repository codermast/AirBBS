package models

import "time"

// Config 配置结构体
type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	JWT      JWT      `yaml:"jwt"`
	Redis    Redis    `yaml:"redis"`
}

// Database 数据库配置信息结构体
type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBName   string `json:"dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type JWT struct {
	SigningKey  string        `yaml:"signing_key" json:"signing_key"`
	Header      string        `yaml:"header" json:"header"`
	Secret      string        `yaml:"secret" json:"secret"`
	Issuer      string        `yaml:"issuer" json:"issuer"`
	ExpireHours time.Duration `yaml:"expire_hours" json:"expire_hours"`
	Subject     string        `yaml:"subject" json:"subject"`
}

// Redis 配置信息结构体
type Redis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}
