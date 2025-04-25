package config

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerCfg ServerConfig
	DbCfg     DatabaseConfig
	JwtCfg    JwtConfig
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSLMode  string
	TimeZone string
}

type JwtConfig struct {
	SecretKey   string
	TimeAccess  time.Duration
	TimeRefresh time.Duration
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		config, err := load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
		instance = config
	})
	return instance
}

func load() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	serverConfig := ServerConfig{
		Port: parseInt(getEnv("SERVER_PORT", "8000")),
	}

	dbConfig := DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "user"),
		Password: getEnv("DB_PASSWORD", "password"),
		DbName:   getEnv("DB_NAME", "core_user"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
		TimeZone: getEnv("DB_TIMEZONE", "UTC"),
	}

	jwtConfig := JwtConfig{
		SecretKey:   getEnv("SECRET_KEY", "123123123123123123123123112312312"),
		TimeAccess:  parseDuration(getEnv("TIME_ACCESS", "1000001h")),
		TimeRefresh: parseDuration(getEnv("TIME_REFRESH", "1000001h")),
	}

	return &Config{
		ServerCfg: serverConfig,
		DbCfg:     dbConfig,
		JwtCfg:    jwtConfig,
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func parseInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return intValue
}

func parseDuration(value string) time.Duration {
	duration, err := time.ParseDuration(value)
	if err != nil {
		return 1 * time.Hour
	}
	return duration
}
