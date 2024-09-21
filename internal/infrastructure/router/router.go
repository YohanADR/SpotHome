package router

import (
	postgisPort "SpotHome/internal/infrastructure/postgis/ports"
	redisPort "SpotHome/internal/infrastructure/redis/ports"
	userHandler "SpotHome/internal/user/adapters/http"
	"net/http"

	"github.com/gorilla/mux"
)

// Router est une structure qui encapsule le routeur
type Router struct {
	mux *mux.Router
}

// NewRouter crée et retourne un nouveau Router
func NewRouter() *Router {
	return &Router{
		mux: mux.NewRouter(),
	}
}

// RegisterRoutes enregistre les routes de l'application
func (r *Router) RegisterRoutes(postgres postgisPort.PostGISPort, redis redisPort.RedisPort) {
	r.mux.HandleFunc("/api/example", ExampleHandler).Methods("GET")
	r.mux.HandleFunc("/api/user", userHandler.GetUserHandler(postgres, redis)).Methods("GET")
	// Ajoute d'autres routes ici
}

// GetMux retourne le routeur mux
func (r *Router) GetMux() *mux.Router {
	return r.mux
}

// Exemple de gestionnaire de route
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}
