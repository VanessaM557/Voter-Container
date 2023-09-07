package config

import (
	"log"
	"os"
)

type Config struct {
	RedisHost     string
	RedisPort     string
	RedisPassword string
	APIPort       string
}

func LoadConfig() *Config {
	config := &Config{
		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnv("REDIS_PORT", "6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		APIPort:       getEnv("API_PORT", ":8080"),
	}

	return config
}
func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Environment variable %s not set, using default: %s", key, defaultValue)
		return defaultValue
	}
	return value
}
