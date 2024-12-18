package config_test

import (
	"bank/config"
	"log"
	"testing"
)

func TestReadYaml(t *testing.T) {
	// 初始化配置
	config.ReadYaml()

	// 输出整个 AppConfig 配置
	log.Printf("AppConfig: %v\n", config.AppConfig)
}

func TestInitConfig(t *testing.T) {
	if err := config.InitConfig(); err != nil {
		log.Fatalf("Error when InitConfig: %v", err)
	}
}
