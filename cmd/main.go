package main

import (
	"SpotHome/internal/infrastructure/config"
	"SpotHome/internal/infrastructure/http_server"
	"SpotHome/internal/infrastructure/logger"
	"SpotHome/internal/infrastructure/postgis"
	"SpotHome/internal/infrastructure/router"
	"fmt"
)

func main() {

	log := logger.NewLogger() // Retourne l'interface logger
	cfg := config.Load(log)   // Charge la configuration
	r := router.NewRouter()   // Crée une nouvelle instance de router

	connString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_SSLMODE)
	pg, err := postgis.NewPostGIS(connString) // Crée une nouvelle instance de PostGIS

	if err != nil {
		log.Error("Échec de la connexion à la base de données PostGIS: %v", logger.FieldError(err))
	}
	defer pg.Close()

	defer log.Sync()   // assure d'écrire les logs avant de quitter
	r.RegisterRoutes() // Enregistre les routes

	log.Info("Démarrage de l'application en mode", logger.FieldString("Environment", cfg.Environment))
	http_server.StartServer(cfg, r, log) // Démarrage du serveur HTTP
}
