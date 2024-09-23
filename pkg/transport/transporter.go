package transport

// RegisterRoutes est une fonction type pour enregistrer les routes

// Transporter est une interface générique pour tout type de serveur HTTP
type Transporter interface {
	Start() error                        // Méthode pour démarrer le serveur HTTP
	RegisterRoutes(func(RegisterRoutes)) // Enregistrement des routes via un RegisterRoutes abstrait
}
