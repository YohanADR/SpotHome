package adapters

import (
	"SpotHome/internal/infrastructure/config"
	"SpotHome/internal/infrastructure/postgis/ports"
	"SpotHome/internal/user/domain"

	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
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

func (p *PostGISAdapter) Query(query string, args ...interface{}) (interface{}, error) {
	rows, err := p.pool.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Process rows and return results
	var results []domain.User // Adjust based on your actual result type
	for rows.Next() {
		var user domain.User                                                 // Using the User struct
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil { // Adjust based on your result fields
			return nil, err
		}
		results = append(results, user)
	}
	return results, nil
}

// QueryRow exécute une requête SQL et retourne un pgx.Row
func (p *PostGISAdapter) QueryRow(query string, args ...interface{}) pgx.Row {
	return p.pool.QueryRow(context.Background(), query, args...)
}

// Close ferme la connexion à la base de données
func (p *PostGISAdapter) Close() {
	p.pool.Close()
}
