package config

import (
	"codermast.com/airblog/models"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var config *models.Config

func AirBlogSystemConfig() {
	// 1. 加载配置文件
	file, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}
}

func GetDatabaseConfig() *models.Database {
	return &config.Database
}
