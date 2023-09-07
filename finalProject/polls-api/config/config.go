package config

import (
	"os"
	"log"
)

var (
	RedisHost     = getEnv("REDIS_HOST", "localhost:6379")
	RedisPassword = getEnv("REDIS_PASSWORD", "")
	RedisDB       = getEnvAsInt("REDIS_DB", 0)
)

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
		log.Printf("Failed to parse env variable %s: %v. Using default value: %d", key, err, defaultValue)
	}
	return defaultValue
}
