package main

import (
	"github.com/YohanADR/SpotHome/infrastructure/app"
)

func main() {
	// Initialisation de l'application
	application, err := app.InitApp()
	if err != nil {
		panic("Échec de l'initialisation de l'application")
	}
	defer application.CloseResources()

	// Démarrage du serveur HTTP
	if err := application.Router.Transporter.Start(); err != nil {
		application.Logger.Fatal("Erreur lors du démarrage du serveur HTTP", "error", err)
	}
}
