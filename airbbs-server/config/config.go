package config

import (
	"codermast.com/airbbs/models"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sync"
)

var (
	config     *models.Config
	onceConfig sync.Once
)

func AirBbsSystemConfig() {
	// 仅加载一次
	onceConfig.Do(func() {
		// 1. 加载配置文件
		file, err := ioutil.ReadFile("config.yaml")
		if err != nil {
			log.Printf("Error reading config file: %v", err)
			return
		}

		// 2. 解析配置文件到全局变量 config
		err = yaml.Unmarshal(file, &config)
		if err != nil {
			log.Printf("Error parsing config file: %v", err)
			return
		}
	})
}

// GetServerConfig 获取系统配置
func GetServerConfig() *models.Server {
	AirBbsSystemConfig() // 确保配置加载完成
	return &config.Server
}

// GetDatabaseConfig 获取数据库配置信息
func GetDatabaseConfig() *models.Database {
	AirBbsSystemConfig() // 确保配置加载完成
	return &config.Database
}

// GetJWTConfig 获取 JWT 配置信息
func GetJWTConfig() *models.JWT {
	AirBbsSystemConfig() // 确保配置加载完成
	return &config.JWT
}
