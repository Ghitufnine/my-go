package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string

	PostgresHost string
	PostgresPort string
	PostgresUser string
	PostgresPass string
	PostgresDB   string

	MongoURI string
	MongoDB  string

	RedisAddr     string
	RedisPassword string
	RedisDB       int

	RabbitURL      string
	RabbitExchange string
}

func Load() *Config {

	// Load .env file — try project root first, then two levels up (for cmd/api runs)
	if err := godotenv.Load(".env"); err != nil {
		if err2 := godotenv.Load("../../.env"); err2 != nil {
			log.Println("Warning: .env file not found, using defaults")
		}
	}
	return &Config{
		AppPort: getEnv("APP_PORT", "8080"),

		PostgresHost: getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort: getEnv("POSTGRES_PORT", "5432"),
		PostgresUser: getEnv("POSTGRES_USER", "postgres"),
		PostgresPass: getEnv("POSTGRES_PASSWORD", "postgres"),
		PostgresDB:   getEnv("POSTGRES_DB", "clean_arch"),

		MongoURI: getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDB:  getEnv("MONGO_DB", "clean_arch_logs"),

		RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       0,

		RabbitURL:      getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
		RabbitExchange: getEnv("RABBITMQ_EXCHANGE", "events"),
	}
}

func getEnv(key, fallback string) string {

	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
