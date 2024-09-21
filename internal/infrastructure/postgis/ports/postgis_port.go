// postgis_port.go (dans le package domain)
package ports

import "github.com/jackc/pgx/v4"

// PostGISPort est une interface qui définit les opérations possibles avec la base de données
type PostGISPort interface {
	Query(query string, args ...interface{}) (interface{}, error)
	QueryRow(query string, args ...interface{}) pgx.Row
	Close()
}
