package config

import (
	"SpotHome/internal/infrastructure/logger"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config structure pour stocker les configurations
type Config struct {
	Port        string
	Environment string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     int
	DB_SSLMODE  string
	REDIS_HOST  string
	REDIS_PORT  int
	REDIS_DB    int
}

// LoadCharge le fichier .env et retourne la configuration
func Load(log logger.Logger) *Config {
	err := godotenv.Load()
	if err != nil {
		log.Error("Aucun fichier .env trouvé, utilisation des valeurs par défaut", logger.FieldError(err))
	}

	dbPort, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))       // Valeur par défaut pour le port
	redisPort, _ := strconv.Atoi(getEnv("REDIS_PORT", "6379")) // Valeur par défaut pour le port
	redisDb, _ := strconv.Atoi(getEnv("REDIS_DB", "0"))        // Valeur par défaut pour le port

	return &Config{
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("ENVIRONMENT", "developpement"),
		DB_USER:     getEnv("DB_USER", "SpotAdm"),
		DB_PASSWORD: getEnv("DB_PASSWORD", "password123"),
		DB_NAME:     getEnv("DB_NAME", "spothome"),
		DB_HOST:     getEnv("DB_HOST", "localhost"),
		DB_PORT:     dbPort,
		DB_SSLMODE:  getEnv("DB_SSLMODE", "disable"),
		REDIS_HOST:  getEnv("REDIS_HOST", "redis"),
		REDIS_PORT:  redisPort,
		REDIS_DB:    redisDb,
	}
}

// getEnv retourne la valeur d'une variable d'environnement ou une valeur par défaut
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
