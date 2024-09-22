package transport

import (
	"github.com/YohanADR/SpotHome/infrastructure/logger"
	"github.com/gin-gonic/gin"
)

type GinTransport struct {
	Engine *gin.Engine
	Addr   string
	Logger logger.Logger
}

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

// RegisterRoutes enregistre les routes dans Gin
func (g *GinTransport) RegisterRoutes(registerFunc func(HandlerFunc)) {
	// Définir un HandlerFunc pour le registre qui sera compatible avec Gin
	registerFunc(func(method string, path string, handlerFunc interface{}) {
		switch method {
		case "GET":
			g.Engine.GET(path, handlerFunc.(gin.HandlerFunc))
		case "POST":
			g.Engine.POST(path, handlerFunc.(gin.HandlerFunc))
		// Ajouter les autres méthodes HTTP ici
		default:
			g.Logger.Info("Méthode HTTP non supportée", "method", method)
		}
	})
}
