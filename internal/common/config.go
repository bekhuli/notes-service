package common

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	PublicHost string
	Port       string
	User       string
	Password   string
	Name       string
	SSLMode    string
}

var (
	DBEnv = initDBConfig()
)

func initDBConfig() DBConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	return DBConfig{
		PublicHost: getEnv("DB_HOST"),
		Port:       getEnv("DB_PORT"),
		User:       getEnv("DB_USER"),
		Password:   getEnv("DB_PASSWORD"),
		Name:       getEnv("DB_NAME"),
		SSLMode:    getEnv("DB_SSL_MODE"),
	}
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Environment variable %s not set", key)
	}

	return value
}
