package transport

// HandlerFunc est un type abstrait de fonction pour gérer des requêtes (peut être implémenté par Gin, Echo, etc.)
type HandlerFunc func(method string, path string, handlerFunc interface{})

// Transporter interface abstrait la couche de transport (Gin, HTTP, etc.)
type Transporter interface {
	Start() error
	RegisterRoutes(registerFunc func(HandlerFunc)) // Méthode abstraite pour enregistrer des routes
}
