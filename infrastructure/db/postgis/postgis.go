package postgis

import (
	"context"
	"time"

	"github.com/YohanADR/SpotHome/infrastructure/config"
	"github.com/YohanADR/SpotHome/pkg/logger"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostGISClient struct {
	Pool   *pgxpool.Pool
	Logger logger.Logger
}

// NewPostGISClient initialise une connexion à PostGIS en utilisant les configurations
func NewPostGISClient(cfg config.DatabaseConfig, log logger.Logger) (*PostGISClient, error) {
	// Créer l'URL de connexion à PostGIS
	dsn := "postgres://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "/" + cfg.Name + "?sslmode=" + cfg.SSLMode

	// Configuration de la pool de connexions
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal("Erreur lors du parsing de la configuration PostGIS", "error", err)
		return nil, err
	}

	// Optionnel : tu peux ajuster des paramètres de pool ici (max connexions, etc.)
	config.MaxConns = 10
	config.MinConns = 1
	config.HealthCheckPeriod = 2 * time.Minute

	// Initialisation du pool de connexions
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Impossible de se connecter à PostGIS", "error", err)
		return nil, err
	}

	log.Info("Connexion à PostGIS établie", "host", cfg.Host, "port", cfg.Port)

	return &PostGISClient{
		Pool:   pool,
		Logger: log,
	}, nil
}

// Close ferme la connexion au pool PostGIS
func (p *PostGISClient) Close() {
	if p.Pool != nil {
		p.Pool.Close()
		p.Logger.Info("Connexion PostGIS fermée avec succès")
	}
}
