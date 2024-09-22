package main

import (
	"github.com/YohanADR/SpotHome/infrastructure/server"
	"github.com/YohanADR/SpotHome/pkg/app"
	"github.com/YohanADR/SpotHome/pkg/logger"
)

func main() {
	// Initialisation de l'application (config + logger)
	application, err := app.InitApp()
	if err != nil {
		panic("Échec de l'initialisation de l'application")
	}
	defer logger.ShutdownLogger(application.Logger)

	// Démarrage du serveur HTTP
	httpServer := server.NewHTTPServer(application.Config.Server, application.Logger)
	if err := httpServer.Start(); err != nil {
		application.Logger.Fatal("Erreur lors du démarrage du serveur HTTP", "error", err)
	}
}
