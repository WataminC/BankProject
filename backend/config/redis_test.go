package config_test

import (
	"bank/config"
	"log"
	"testing"
)

func TestRedisConnection(t *testing.T) {
	// Call InitRedis and check if it returns an error
	config.ReadYaml("./")
	rdb, err := config.InitRedis()
	if err != nil {
		// If initialization fails, print the error message
		t.Fatalf("Error initializing Redis: %v", err)
	} else {
		// If initialization is successful, output a success message
		log.Println("Redis initialized successfully")
	}

	// Close the connection after the test
	defer rdb.Close()
}
