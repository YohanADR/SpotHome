package config

import (
	"time"

	"github.com/YohanADR/SpotHome/pkg/logger"
	"github.com/spf13/viper"
)

// Config struct pour stocker les configurations du projet
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
}

// ServerConfig contient les configurations du serveur HTTP
type ServerConfig struct {
	Port         string        `mapstructure:"port"`
	Environment  string        `mapstructure:"environment"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

// DatabaseConfig contient les informations de connexion à la base de données
type DatabaseConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	SSLMode  string `mapstructure:"sslmode"`
}

// RedisConfig contient les informations de connexion à Redis
type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	DB   int    `mapstructure:"db"`
}

// LoadConfig lit les configurations depuis le fichier .yaml
func LoadConfig(configPath string, log logger.Logger) (*Config, error) {
	var config Config

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	// Lire le fichier de configuration
	if err := viper.ReadInConfig(); err != nil {
		log.Error("Erreur lors de la lecture du fichier de configuration", "error", err)
		return nil, err
	}

	// Unmarshal les données dans la struct Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Error("Erreur lors de l'analyse de la configuration", "error", err)
		return nil, err
	}

	log.Info("Configuration chargée avec succès")
	return &config, nil
}
