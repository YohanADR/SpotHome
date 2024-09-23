package transport

import (
	"net/http"

	"github.com/YohanADR/SpotHome/infrastructure/logger"
	"github.com/gin-gonic/gin"
)

// GinTransport structure qui utilise Gin comme moteur HTTP et implémente HTTPHandler
type GinTransport struct {
	Engine *gin.Engine
	Addr   string
	Logger logger.Logger
}

// NewGinTransport crée un nouveau transport HTTP utilisant Gin
func NewGinTransport(addr string, log logger.Logger) *GinTransport {
	r := gin.Default()

	log.Info("Initialisation du transport avec Gin")
	return &GinTransport{
		Engine: r,
		Addr:   addr,
		Logger: log,
	}
}

// Start démarre le serveur Gin
func (g *GinTransport) Start() error {
	g.Logger.Info("Démarrage du serveur Gin", "addr", g.Addr)
	return g.Engine.Run(g.Addr)
}

// RegisterRoutes enregistre les routes avec Gin
func (g *GinTransport) RegisterRoutes(registerFunc func(RegisterRoutes)) {
	// Enregistre les routes en utilisant le handler Gin
	registerFunc(func(method, path string, handlerFunc http.HandlerFunc) {
		switch method {
		case "GET":
			g.Engine.GET(path, gin.WrapF(handlerFunc))
		case "POST":
			g.Engine.POST(path, gin.WrapF(handlerFunc))
		// Ajouter les autres méthodes HTTP ici
		default:
			g.Logger.Error("Méthode HTTP non supportée", "method", method)
		}
	})
}

// ServeHTTP permet à Gin de respecter l'interface HTTPHandler
func (g *GinTransport) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.Engine.ServeHTTP(w, r)
}
