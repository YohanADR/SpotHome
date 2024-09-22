package transport

import "net/http"

// HTTPHandler est une interface pour gérer les requêtes HTTP
type HTTPHandler interface {
	http.Handler
	Start() error
}
