package app

import (
	"github.com/YohanADR/SpotHome/infrastructure/config"
	"github.com/YohanADR/SpotHome/infrastructure/db/postgis"
	"github.com/YohanADR/SpotHome/infrastructure/db/redis"
	"github.com/YohanADR/SpotHome/infrastructure/logger"
	"github.com/YohanADR/SpotHome/infrastructure/messaging/kafka"
)

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
