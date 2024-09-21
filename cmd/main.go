package main

import (
	"SpotHome/internal/infrastructure/config"
	"SpotHome/internal/infrastructure/http_server"
	"SpotHome/internal/infrastructure/logger"
	"SpotHome/internal/infrastructure/router"
)

func main() {

	cfg := config.Load()      // Charge la configuration
	log := logger.NewLogger() // Retourne l'interface logger
	r := router.NewRouter()   // Crée une nouvelle instance de router

	defer log.Sync()   // assure d'écrire les logs avant de quitter
	r.RegisterRoutes() // Enregistre les routes

	log.Info("Démarrage de l'application en mode", logger.FieldString("Environment", cfg.Environment))
	http_server.StartServer(cfg, r, log)
}
