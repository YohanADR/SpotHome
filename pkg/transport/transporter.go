package transport

import "net/http"

// RouteRegistrar est un type pour l'enregistrement de routes abstrait
type RouteRegistrar func(method, path string, handler http.HandlerFunc)

// Transporter est une interface générique pour tout type de serveur HTTP
type Transporter interface {
	http.Handler                         // Respecte l'interface http.Handler
	Start() error                        // Méthode pour démarrer le serveur HTTP
	RegisterRoutes(func(RouteRegistrar)) // Enregistrement des routes via un RouteRegistrar abstrait
}
