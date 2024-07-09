package models

// Config 配置结构体
type Config struct {
	Database Database `yaml:"database"`
	JWT      JWT      `yaml:"jwt"`
}
