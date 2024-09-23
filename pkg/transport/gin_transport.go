package transport

import (
	"github.com/YohanADR/SpotHome/infrastructure/logger"
	"github.com/gin-gonic/gin"
)

// GinTransport struct pour Gin
type GinTransport struct {
	Engine *gin.Engine
	Logger logger.Logger
	Addr   string
}

// RegisterRoutes est une fonction type qui prend une méthode, un chemin et un gestionnaire HTTP générique
type RegisterRoutes func(method string, path string, handler interface{})

// NewGinTransport initialise un nouveau GinTransport
func NewGinTransport(addr string, log logger.Logger) *GinTransport {
	engine := gin.Default()
	return &GinTransport{
		Engine: engine,
		Logger: log,
		Addr:   addr,
	}
}

// RegisterRoutes implémente Transporter en utilisant RegisterRoutes
func (g *GinTransport) RegisterRoutes(register func(RegisterRoutes)) {
	register(func(method string, path string, handler interface{}) {
		if h, ok := handler.(gin.HandlerFunc); ok {
			switch method {
			case "GET":
				g.Engine.GET(path, h)
			case "POST":
				g.Engine.POST(path, h)
			// Ajouter d'autres méthodes HTTP si nécessaire
			default:
				g.Logger.Error("Méthode HTTP non supportée", "method", method)
			}
		} else {
			g.Logger.Error("Handler invalide pour Gin", "path", path)
		}
	})
}

// Start démarre le serveur Gin
func (g *GinTransport) Start() error {
	g.Logger.Info("Démarrage du serveur Gin", "addr", g.Addr)
	return g.Engine.Run(g.Addr)
}
