package router

import (
	"github.com/YohanADR/SpotHome/infrastructure/logger"
	"github.com/YohanADR/SpotHome/pkg/transport"
)

// Router struct qui contiendra le transport
type Router struct {
	Transporter transport.Transporter
	Logger      logger.Logger
}

// NewRouter initialise le routeur avec un transporteur
func NewRouter(transporter transport.Transporter, log logger.Logger) *Router {
	log.Info("Router initialisé avec un transporteur")

	return &Router{
		Transporter: transporter,
		Logger:      log,
	}
}

// RegisterRoutes permet à chaque contexte métier d'enregistrer ses routes via un HandlerFunc abstrait
func (r *Router) RegisterRoutes(registerFunc func(transport.RouteRegistrar)) {
	// Utilise le transporteur abstrait pour enregistrer des routes avec HandlerFunc
	r.Transporter.RegisterRoutes(registerFunc)
}
