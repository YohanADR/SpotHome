package app

import (
	"github.com/YohanADR/SpotHome/infrastructure/config"
	"github.com/YohanADR/SpotHome/infrastructure/db/postgis"
	"github.com/YohanADR/SpotHome/infrastructure/db/redis"
	"github.com/YohanADR/SpotHome/pkg/logger"
)

type Application struct {
	Config        *config.Config
	Logger        logger.Logger
	RedisClient   *redis.RedisClient
	PostGISClient *postgis.PostGISClient
}

// InitApp initialise la configuration, le logger, Redis, et PostGIS
func InitApp() (*Application, error) {
	// Initialiser le logger
	log := logger.InitLogger()

	// Charger la configuration depuis le fichier .yaml
	cfg, err := config.LoadConfig("config/config.yaml", log)
	if err != nil {
		log.Fatal("Erreur lors du chargement de la configuration", "error", err)
		return nil, err
	}

	// Initialiser Redis avec la configuration
	redisClient := redis.NewRedisClient(cfg.Redis, log)

	// Initialiser PostGIS avec la configuration
	postgisClient, err := postgis.NewPostGISClient(cfg.Database, log)
	if err != nil {
		log.Fatal("Erreur lors de l'initialisation de PostGIS", "error", err)
		return nil, err
	}

	log.Info("Application initialisée avec succès")

	return &Application{
		Config:        cfg,
		Logger:        log,
		RedisClient:   redisClient,
		PostGISClient: postgisClient,
	}, nil
}
