package app

import (
	"github.com/YohanADR/SpotHome/infrastructure/config"
	"github.com/YohanADR/SpotHome/infrastructure/db/postgis"
	"github.com/YohanADR/SpotHome/infrastructure/db/redis"
	"github.com/YohanADR/SpotHome/infrastructure/logger"
	"github.com/YohanADR/SpotHome/infrastructure/messaging/kafka"
	"github.com/YohanADR/SpotHome/infrastructure/server/router"
	"github.com/YohanADR/SpotHome/pkg/transport"
	"github.com/gin-gonic/gin"
)

// Application struct contient les services principaux
type Application struct {
	Config        *config.Config
	Logger        logger.Logger
	RedisClient   *redis.RedisClient
	PostGISClient *postgis.PostGISClient
	KafkaProducer *kafka.KafkaProducer
	Router        *router.Router
}

// InitApp initialise l'application
func InitApp() (*Application, error) {
	// Initialisation du logger
	log := logger.InitLogger()

	// Charger la configuration
	cfg, err := config.LoadConfig("config/config.yaml", log)
	if err != nil {
		return nil, handleFatalError(log, "Erreur lors du chargement de la configuration", err)
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

	// Initialiser le transport HTTP avec Gin
	transporter := transport.NewGinTransport(":"+cfg.Server.Port, log)

	// Initialiser le routeur
	appRouter := router.NewRouter(transporter, log)

	// Enregistrement des routes
	registerRoutes(appRouter)

	log.Info("Application initialisée avec succès")

	return &Application{
		Config:        cfg,
		Logger:        log,
		RedisClient:   redisClient,
		PostGISClient: postgisClient,
		KafkaProducer: kafkaProducer,
		Router:        appRouter,
	}, nil
}

// Centralise la gestion des erreurs fatales
func handleFatalError(log logger.Logger, message string, err error) error {
	log.Fatal(message, "error", err)
	return err
}

// Initialiser Redis
func initRedis(cfg *config.Config, log logger.Logger) (*redis.RedisClient, error) {
	redisClient := redis.NewRedisClient(cfg.Redis, log)
	if redisClient == nil {
		return nil, handleFatalError(log, "Erreur lors de l'initialisation de Redis", nil)
	}
	return redisClient, nil
}

// Initialiser PostGIS
func initPostGIS(cfg *config.Config, log logger.Logger) (*postgis.PostGISClient, error) {
	postgisClient, err := postgis.NewPostGISClient(cfg.Database, log)
	if err != nil {
		return nil, handleFatalError(log, "Erreur lors de l'initialisation de PostGIS", err)
	}
	return postgisClient, nil
}

// Initialiser Kafka
func initKafka(cfg *config.Config, log logger.Logger) (*kafka.KafkaProducer, error) {
	kafkaProducer, err := kafka.NewKafkaProducer(cfg.Kafka.Brokers, cfg.Kafka.Topic, log)
	if err != nil {
		return nil, handleFatalError(log, "Erreur lors de l'initialisation de Kafka", err)
	}
	return kafkaProducer, nil
}

// Register les routes
func registerRoutes(appRouter *router.Router) {
	appRouter.RegisterRoutes(func(register transport.HandlerFunc) {
		register("GET", "/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "OK"})
		})
	})
}

// CloseResources ferme toutes les connexions ouvertes et les ressources
func (app *Application) CloseResources() {
	app.Logger.Info("Fermeture des ressources...")

	// Fermer Redis
	if app.RedisClient != nil {
		app.RedisClient.Close()
	}

	// Fermer PostGIS
	if app.PostGISClient != nil {
		app.PostGISClient.Close()
	}

	// Fermer Kafka
	if app.KafkaProducer != nil {
		app.KafkaProducer.Close()
	}

	app.Logger.Info("Toutes les ressources ont été fermées.")
}
