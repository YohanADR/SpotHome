package http_server

import (
	"net/http"

	"SpotHome/internal/infrastructure/config"
	"SpotHome/internal/infrastructure/logger"
	"SpotHome/internal/infrastructure/router"
)

// StartServer démarre le serveur HTTP
func StartServer(cfg *config.Config, r *router.Router, log logger.Logger) {

	log.Info("Démarrage du serveur", logger.FieldString("port", cfg.Port))
	if err := http.ListenAndServe(":"+cfg.Port, r.GetMux()); err != nil {
		log.Error("Erreur lors du démarrage du serveur", logger.FieldError(err))
	}
}
