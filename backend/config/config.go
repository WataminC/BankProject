package config

import (
	"log"
	"os"
	"bank/global"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	DataBase struct {
		Dsn           string
		MaxIdleConns  int
		MaxOpennConns int
	}
	Redis struct {
		Addr     string
		Port     string
		PassWord string
		DB       int
	}
}

var AppConfig *Config

func ReadYaml(filePath string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Println("Error:", err)
		return
	}
	log.Println("Current Directory:", dir)

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	// Should be changed to "./" when run test
	viper.AddConfigPath(filePath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error when reading config file: %v", err)
	}

	AppConfig = &Config{}

	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Error when decoding config file: %v", err)
	}
}

func InitConfig(filePath string) error {
	ReadYaml(filePath)

	if err := InitDB(); err != nil {
		log.Fatalf("Error when reading config file: %v", err)
	}

	rdb, err := InitRedis()
	if err != nil {
		log.Fatalf("Error when reading config file: %v", err)
	}

	global.RDB = rdb

	return nil
}
