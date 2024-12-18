package config_test

import (
	"bank/config"
	"log"
	"testing"
)

func TestInitDB(t *testing.T) {
	// 调用 InitDB 并检查是否返回错误
	config.ReadYaml()
	err := config.InitDB()
	if err != nil {
		// 如果初始化失败，打印错误信息
		t.Fatalf("Error initializing database: %v", err)
	} else {
		// 如果初始化成功，输出成功消息
		log.Println("Database initialized successfully")
	}
}
