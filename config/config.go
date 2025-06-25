package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ServerConfig struct{
	Host string
	Port string
	Debug bool
}

type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type Config struct{
	Server *ServerConfig
	DB * DbConfig
}

func GetConfig() *Config {
	envInit()

	return &Config{
		Server: &ServerConfig{
			Host: getStringFromEnv("SERVER_HOST", "localhost"),
			Port: getStringFromEnv("SERVER_PORT", "3000"),
			Debug: getBoolFromEnv("DB_DEBUG", true),
		},
		DB: &DbConfig{
			Host: getStringFromEnv("DB_HOST", "localhost"),
			Port: getIntFromEnv("DB_PORT", 5437),
			User: getStringFromEnv("DB_USER", "postgres"),
			Password: getStringFromEnv("DB_PASS", "postgres"),
			Name: getStringFromEnv("DB_NAME", "PURPLE_TEST"),
		},
	}

}

func envInit() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Load .env file error")
		return
	}
	log.Println("Load .env ok")
}

func getStringFromEnv(key string, defaultValue string) string {
	envValueStr := os.Getenv(key)
	if envValueStr == "" {
		return  defaultValue
	}
	return envValueStr
}

func getIntFromEnv(key string, defaultValue int) int {
	envValueStr := os.Getenv(key)
	if envValueStr == "" {
		return  defaultValue
	}
	envValueInt, err := strconv.Atoi(envValueStr)
	if err != nil {
		return defaultValue
	}
	return envValueInt
}

func getBoolFromEnv(key string, defaultValue bool) bool {
	envValueStr := os.Getenv(key)
	if envValueStr == "" {
		return  defaultValue
	}
	envValueBool, err := strconv.ParseBool(envValueStr)
	if err != nil {
		return defaultValue
	}
	return envValueBool
}