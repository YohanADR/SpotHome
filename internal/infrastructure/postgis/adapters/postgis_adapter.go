package adapters

import (
	"SpotHome/internal/infrastructure/config"
	"SpotHome/internal/infrastructure/postgis/ports"

	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

// PostGISAdapter est une structure qui implémente le port PostGIS
type PostGISAdapter struct {
	pool *pgxpool.Pool
}

// NewPostGISAdapter crée une nouvelle instance de PostGISAdapter
func NewPostGISAdapter(cfg *config.Config) (ports.PostGISPort, error) {
	connString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_SSLMODE)

	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return &PostGISAdapter{pool: pool}, nil
}

// Query exécute une requête sur la base de données
func (p *PostGISAdapter) Query(query string, args ...interface{}) (interface{}, error) {
	row := p.pool.QueryRow(context.Background(), query, args...)
	var result interface{}
	if err := row.Scan(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// Close ferme la connexion à la base de données
func (p *PostGISAdapter) Close() {
	p.pool.Close()
}
