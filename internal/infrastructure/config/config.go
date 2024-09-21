package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config structure pour stocker les configurations
type Config struct {
	Port        string
	Environment string
}

// LoadCharge le fichier .env et retourne la configuration
func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aucun fichier .env trouvé, utilisation des valeurs par défaut")
	}

	return &Config{
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("ENVIRONMENT", "developpement"),
	}
}

// getEnv retourne la valeur d'une variable d'environnement ou une valeur par défaut
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
