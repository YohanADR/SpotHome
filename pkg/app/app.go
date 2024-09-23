package app

import (
	"github.com/YohanADR/SpotHome/infrastructure/config"
	"github.com/YohanADR/SpotHome/infrastructure/db/postgis"
	"github.com/YohanADR/SpotHome/infrastructure/db/redis"
	"github.com/YohanADR/SpotHome/infrastructure/logger"
	"github.com/YohanADR/SpotHome/infrastructure/messaging/kafka"
	"github.com/YohanADR/SpotHome/infrastructure/server/router"
	"github.com/YohanADR/SpotHome/pkg/events"
	"github.com/YohanADR/SpotHome/pkg/jwt"
	"github.com/YohanADR/SpotHome/pkg/transport"
)

type Application struct {
	Config        *config.Config
	Logger        logger.Logger
	RedisClient   *redis.RedisClient
	PostGISClient *postgis.PostGISClient
	KafkaProducer *kafka.KafkaProducer
	Router        *router.Router
	HTTPHandler   transport.Transporter
	JWTService    *jwt.JWTService
}

// InitApp initialise l'application
func InitApp() (*Application, error) {
	// Initialisation du logger
	log := logger.InitLogger()

	// Charger la configuration
	cfg, err := initConfig("config/config.yaml", log)
	if err != nil {
		return nil, err
	}

	// Initialiser les services
	redisClient, err := initRedis(cfg, log)
	if err != nil {
		return nil, err
	}

	postgisClient, err := initPostGIS(cfg, log)
	if err != nil {
		return nil, err
	}

	kafkaProducer, err := initKafka(cfg, log)
	if err != nil {
		return nil, err
	}

	// Initialiser le système d'événements
	events.InitEventSystem(kafkaProducer, log)

	// Initialiser le service JWT
	jwtService := jwt.NewJWTService(log, kafkaProducer, postgisClient)

	// Initialiser le transport HTTP (Gin ou un autre HTTPHandler)
	httpHandler := transport.NewGinTransport(":"+cfg.Server.Port, log)

	// Initialiser le routeur avec le transport HTTP générique
	appRouter := router.NewRouter(httpHandler, log)

	// Enregistrement des routes
	router.RegisterRoutes(appRouter, jwtService)

	log.Info("Application initialisée avec succès")

	return &Application{
		Config:        cfg,
		Logger:        log,
		RedisClient:   redisClient,
		PostGISClient: postgisClient,
		KafkaProducer: kafkaProducer,
		Router:        appRouter,
		HTTPHandler:   httpHandler,
		JWTService:    jwtService,
	}, nil
}
