package server

import (
	"net/http"
	"time"

	"github.com/YohanADR/SpotHome/infrastructure/config"
	"github.com/YohanADR/SpotHome/pkg/logger"
)

type HTTPServer struct {
	Engine       *http.ServeMux
	Addr         string
	Logger       logger.Logger
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// NewHTTPServer initialise un nouveau serveur HTTP avec les configurations depuis config.yaml
func NewHTTPServer(cfg config.ServerConfig, log logger.Logger) *HTTPServer {
	return &HTTPServer{
		Engine:       http.NewServeMux(),
		Addr:         ":" + cfg.Port,
		Logger:       log,
		ReadTimeout:  cfg.ReadTimeout * time.Second,  // Conversion des secondes en time.Duration
		WriteTimeout: cfg.WriteTimeout * time.Second, // Conversion des secondes en time.Duration
	}
}

// Start démarre le serveur HTTP et retourne une erreur si nécessaire
func (s *HTTPServer) Start() error {
	srv := &http.Server{
		Addr:           s.Addr,
		Handler:        s.Engine,
		ReadTimeout:    s.ReadTimeout,
		WriteTimeout:   s.WriteTimeout,
		MaxHeaderBytes: 1 << 20, // Limite de 1 Mo pour la taille des en-têtes
	}

	if s.Logger != nil {
		s.Logger.Info("Starting HTTP server", "addr", s.Addr)
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		if s.Logger != nil {
			s.Logger.Error("Failed to start HTTP server", "error", err)
		}
		return err
	}

	return nil
}
