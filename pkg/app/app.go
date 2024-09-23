package app

import (
	"net/http"

	"github.com/YohanADR/SpotHome/infrastructure/config"
	"github.com/YohanADR/SpotHome/infrastructure/db/postgis"
	"github.com/YohanADR/SpotHome/infrastructure/db/redis"
	"github.com/YohanADR/SpotHome/infrastructure/logger"
	"github.com/YohanADR/SpotHome/infrastructure/messaging/kafka"
	"github.com/YohanADR/SpotHome/infrastructure/server/router"
	"github.com/YohanADR/SpotHome/pkg/jwt"
	"github.com/YohanADR/SpotHome/pkg/jwt/middleware"
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

	// Initialiser le service JWT
	jwtService := jwt.NewJWTService(log, kafkaProducer, postgisClient)

	// Initialiser le transport HTTP (Gin ou un autre HTTPHandler)
	httpHandler := transport.NewGinTransport(":"+cfg.Server.Port, log)

	// Initialiser le routeur avec le transport HTTP générique
	appRouter := router.NewRouter(httpHandler, log)

	// Enregistrement des routes
	registerRoutes(appRouter, jwtService)

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

// Centralise la gestion des erreurs fatales
func handleFatalError(log logger.Logger, message string, err error) error {
	log.Fatal(message, "error", err)
	return err
}

// Charger la configuration
func initConfig(path string, log logger.Logger) (*config.Config, error) {
	cfg, err := config.LoadConfig(path, log)
	if err != nil {
		return nil, handleFatalError(log, "Erreur lors du chargement de la configuration", err)
	}
	return cfg, nil
}

// Initialiser Redis
func initRedis(cfg *config.Config, log logger.Logger) (*redis.RedisClient, error) {
	redisClient, err := redis.NewRedisClient(cfg.Redis, log)
	if redisClient == nil {
		return nil, handleFatalError(log, "Erreur lors de l'initialisation de Redis", err)
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

// Register les routes avec possibilité d'utiliser le middleware JWT
func registerRoutes(appRouter *router.Router, jwtService *jwt.JWTService) {
	appRouter.RegisterRoutes(func(register transport.RegisterRoutes) {
		// Route publique
		register("GET", "/health", gin.HandlerFunc(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "OK"})
		}))

		// Route protégée avec le middleware JWT
		register("GET", "/protected", gin.HandlerFunc(func(c *gin.Context) {
			middleware.JWTMiddleware(jwtService)(c) // Appliquer le middleware
			if c.IsAborted() {
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Protected route access granted"})
		}))

		// Route pour générer un token
		// Route pour générer un token
		register("POST", "/generate-token", gin.HandlerFunc(func(c *gin.Context) {
			// Récupérer le nom d'utilisateur à partir de la requête JSON
			var requestBody struct {
				Username string `json:"username"`
			}

			if err := c.ShouldBindJSON(&requestBody); err != nil || requestBody.Username == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request. Username is required."})
				return
			}

			// Générer le token et le refresh token
			token, refreshToken, err := jwtService.GenerateToken(requestBody.Username, 1)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
				return
			}

			// Retourner le token et le refresh token générés
			c.JSON(http.StatusOK, gin.H{
				"token":        token,
				"refreshToken": refreshToken,
			})
		}))
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
