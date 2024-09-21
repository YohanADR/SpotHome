package adapters

import (
	"SpotHome/internal/infrastructure/config"
	"SpotHome/internal/infrastructure/redis/ports"

	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// redisClient est une structure qui implémente l'adaptateur Redis
type redisClient struct {
	client *redis.Client
}

// NewRedisAdapter crée une nouvelle instance de redisClient
func NewRedisAdapter(cfg *config.Config) (ports.RedisPort, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.REDIS_HOST, cfg.REDIS_PORT),
		DB:   cfg.REDIS_DB,
	})

	// Teste la connexion
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("unable to connect to Redis: %v", err)
	}

	return &redisClient{client: rdb}, nil
}

// Close ferme la connexion à Redis
func (r *redisClient) Close() error {
	return r.client.Close()
}

// Set ajoute une clé-value à Redis
func (r *redisClient) Set(key string, value interface{}) error {
	_, err := r.client.Set(context.Background(), key, value, 0).Result()
	return err
}

// Get récupère la valeur d'une clé dans Redis
func (r *redisClient) Get(key string) (string, error) { // Change le type de retour en string
	val, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err // Retourne une chaîne vide en cas d'erreur
	}
	return val, nil
}
