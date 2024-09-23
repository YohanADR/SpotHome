package redis

import (
	"context"

	"github.com/YohanADR/SpotHome/infrastructure/config"
	"github.com/YohanADR/SpotHome/infrastructure/logger"
	"github.com/YohanADR/SpotHome/pkg/cache"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// RedisClient struct qui contient le client Redis
type RedisClient struct {
	Client *redis.Client
	Logger logger.Logger
}

// Vérification que RedisClient implémente bien l'interface Cache
var _ cache.Cache = (*RedisClient)(nil)

// NewRedisClient initialise une connexion Redis en utilisant les configurations
func NewRedisClient(cfg config.RedisConfig, log logger.Logger) (*RedisClient, error) {
	// Initialisation du client Redis
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: "", // Par défaut, aucun mot de passe
		DB:       cfg.DB,
	})

	// Test de connexion à Redis
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Impossible de se connecter à Redis", "error", err)
		return nil, err
	}

	log.Info("Connexion à Redis établie", "host", cfg.Host, "port", cfg.Port)

	return &RedisClient{
		Client: client,
		Logger: log,
	}, nil
}

// Get récupère une valeur dans Redis (implémente Cache)
func (r *RedisClient) Get(key string) (interface{}, error) {
	result, err := r.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		r.Logger.Info("Clé non trouvée", "key", key)
		return "", nil
	} else if err != nil {
		r.Logger.Error("Erreur lors de la récupération de la clé", "error", err, "key", key)
		return "", err
	}
	r.Logger.Info("Clé récupérée avec succès", "key", key, "value", result)
	return result, nil
}

// Set stocke une valeur dans Redis (implémente Cache)
func (r *RedisClient) Set(key string, value interface{}) error {
	err := r.Client.Set(ctx, key, value, 0).Err() // expiration par défaut à 0
	if err != nil {
		r.Logger.Error("Erreur lors de l'enregistrement de la clé", "error", err, "key", key)
		return err
	}
	r.Logger.Info("Clé enregistrée avec succès", "key", key, "value", value)
	return nil
}

// Delete supprime une clé dans Redis (implémente Cache)
func (r *RedisClient) Delete(key string) error {
	err := r.Client.Del(ctx, key).Err()
	if err != nil {
		r.Logger.Error("Erreur lors de la suppression de la clé", "error", err, "key", key)
		return err
	}
	r.Logger.Info("Clé supprimée avec succès", "key", key)
	return nil
}

// Close ferme la connexion Redis (implémente Cache)
func (r *RedisClient) Close() error {
	if err := r.Client.Close(); err != nil {
		r.Logger.Error("Erreur lors de la fermeture de la connexion Redis", "error", err)
		return err
	}
	r.Logger.Info("Connexion Redis fermée avec succès")
	return nil
}
