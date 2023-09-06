package config

import (
	"log"
	"os"
)

type Configuration struct {
	ServerPort string
	RedisAddr  string
	RedisPass  string
	Environment string
}

var Config *Configuration

func LoadConfig() {
	Config = &Configuration{
		ServerPort: getEnv("SERVER_PORT", "8080"),       
		RedisAddr:  getEnv("REDIS_ADDR", "localhost:6379"), 
		RedisPass:  getEnv("REDIS_PASS", ""),         
		Environment: getEnv("ENV", "development"),     
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Environment variable %s not found. Using the default value %s", key, defaultValue)
		return defaultValue
	}
	return value
}
