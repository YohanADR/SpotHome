// postgis/postgis.go
package postgis

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

// PostGIS est une structure pour gérer la connexion à la base de données PostGIS
type PostGIS struct {
	pool *pgxpool.Pool
}

// NewPostGIS crée une nouvelle instance de PostGIS avec une connexion à la base de données
func NewPostGIS(connString string) (*PostGIS, error) {
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	return &PostGIS{pool: pool}, nil
}

// Close ferme la connexion à la base de données
func (p *PostGIS) Close() {
	p.pool.Close()
}
