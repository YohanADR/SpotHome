package main

import (
	"SpotHome/internal/infrastructure/config"
	"SpotHome/internal/infrastructure/http_server"
	"SpotHome/internal/infrastructure/logger"
	postgis "SpotHome/internal/infrastructure/postgis/adapters"
	redis "SpotHome/internal/infrastructure/redis/adapters"
	"SpotHome/internal/infrastructure/router"
)

func main() {

	log := logger.NewLogger() // Retourne l'interface logger
	cfg := config.Load(log)   // Charge la configuration
	r := router.NewRouter()   // Crée une nouvelle instance de router

	pg, err := postgis.NewPostGISAdapter(cfg) // Crée une nouvelle instance de PostGIS
	if err != nil {
		log.Error("Échec de la connexion à la base de données PostGIS: %v", logger.FieldError(err))
	}
	defer pg.Close()

	// Initialisation de Redis
	redisAdapter, err := redis.NewRedisAdapter(cfg)
	if err != nil {
		log.Error("Échec de la connexion à Redis: %v", logger.FieldError(err))
	}
	defer redisAdapter.Close()

	defer log.Sync()   // assure d'écrire les logs avant de quitter
	r.RegisterRoutes() // Enregistre les routes

	log.Info("Démarrage de l'application en mode", logger.FieldString("Environment", cfg.Environment))
	http_server.StartServer(cfg, r, log) // Démarrage du serveur HTTP
}
