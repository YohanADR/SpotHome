package app

import (
	"github.com/YohanADR/SpotHome/infrastructure/config"
	"github.com/YohanADR/SpotHome/pkg/logger"
)

type Application struct {
	Config *config.Config
	Logger logger.Logger
}

// InitApp initialise la configuration et le logger pour l'application
func InitApp() (*Application, error) {
	// Initialiser le logger
	log := logger.InitLogger()

	// Charger la configuration depuis le fichier .yaml
	config, err := config.LoadConfig("config/config.yaml", log)
	if err != nil {
		log.Fatal("Erreur lors du chargement de la configuration", "error", err)
		return nil, err
	}

	log.Info("Application initialisée avec succès")

	// Retourner l'application avec la configuration et le logger
	return &Application{
		Config: config,
		Logger: log,
	}, nil
}
