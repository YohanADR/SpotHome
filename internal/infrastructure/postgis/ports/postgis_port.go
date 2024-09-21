// postgis_port.go (dans le package domain)
package ports

// PostGISPort est une interface qui définit les opérations possibles avec la base de données
type PostGISPort interface {
	Query(query string, args ...interface{}) (interface{}, error)
	Close()
}
