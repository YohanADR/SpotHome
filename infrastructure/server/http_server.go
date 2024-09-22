package server

import (
	"net/http"

	"github.com/YohanADR/SpotHome/infrastructure/config"
	"github.com/YohanADR/SpotHome/infrastructure/logger"
	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	Engine *gin.Engine
	Addr   string
	Logger logger.Logger
	Config config.ServerConfig
}

// NewHTTPServer initialise un nouveau serveur HTTP avec le moteur Gin et la configuration
func NewHTTPServer(cfg config.ServerConfig, engine *gin.Engine, log logger.Logger) *HTTPServer {
	return &HTTPServer{
		Engine: engine,
		Addr:   ":" + cfg.Port,
		Logger: log,
		Config: cfg,
	}
}

// Start démarre le serveur HTTP avec la configuration provenant du fichier .yaml
func (s *HTTPServer) Start() error {
	s.Logger.Info("Démarrage du serveur HTTP", "addr", s.Addr)

	// Configuration du serveur HTTP
	srv := &http.Server{
		Addr:           s.Addr,
		Handler:        s.Engine,
		ReadTimeout:    s.Config.ReadTimeout,
		WriteTimeout:   s.Config.WriteTimeout,
		MaxHeaderBytes: s.Config.MaxHeaderBytes,
	}

	// Démarrer le serveur HTTP
	if err := srv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.Logger.Info("Le serveur HTTP a été arrêté proprement", "addr", s.Addr)
		} else {
			s.Logger.Error("Erreur critique lors du démarrage du serveur HTTP", "addr", s.Addr, "error", err)
		}
		return err
	}

	s.Logger.Info("Le serveur HTTP s'est arrêté de manière inattendue", "addr", s.Addr)
	return nil
}
