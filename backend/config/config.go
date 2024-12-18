package config

import (
	"log"

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
}

var AppConfig *Config

func ReadYaml() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	// Should be changed to "./" when run test
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error when reading config file: %v", err)
	}

	AppConfig = &Config{}

	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Error when decoding config file: %v", err)
	}
}

func InitConfig() error {
	ReadYaml()

	if err := InitDB(); err != nil {
		log.Fatalf("Error when reading config file: %v", err)
	}

	return nil
}
