package postgis

import (
	"context"
	"errors"
	"time"

	"github.com/YohanADR/SpotHome/infrastructure/config"
	"github.com/YohanADR/SpotHome/infrastructure/logger"
	"github.com/YohanADR/SpotHome/pkg/database"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostGISClient struct {
	Pool   *pgxpool.Pool
	Logger logger.Logger
}

// Définition d'une erreur personnalisée pour un pool fermé
var ErrClosedPool = errors.New("le pool de connexions est fermé")

// Vérification que PostGISClient implémente bien l'interface Database
var _ database.Database = (*PostGISClient)(nil)

// NewPostGISClient initialise une connexion à PostGIS en utilisant les configurations
func NewPostGISClient(cfg config.DatabaseConfig, log logger.Logger) (*PostGISClient, error) {
	// Créer l'URL de connexion à PostGIS
	dsn := "postgres://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "/" + cfg.Name + "?sslmode=" + cfg.SSLMode

	// Configuration du pool de connexions
	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal("Erreur lors du parsing de la configuration PostGIS", "error", err)
		return nil, err
	}

	// Ajuster les paramètres de pool ici (max connexions, etc.)
	poolConfig.MaxConns = 10
	poolConfig.MinConns = 1
	poolConfig.HealthCheckPeriod = 2 * time.Minute

	// Initialisation du pool de connexions
	pool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
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

// Connect est une méthode requise par l'interface Database mais ici elle n'est pas utilisée directement
func (p *PostGISClient) Connect() error {
	if p.Pool == nil {
		return ErrClosedPool
	}
	return nil
}

// Query exécute une requête SQL
func (p *PostGISClient) Query(query string, args ...interface{}) (interface{}, error) {
	rows, err := p.Pool.Query(context.Background(), query, args...)
	if err != nil {
		p.Logger.Error("Erreur lors de l'exécution de la requête", "query", query, "error", err)
		return nil, err
	}
	defer rows.Close()

	// Ici, nous retournons simplement les lignes récupérées. La logique pourrait être améliorée selon les besoins.
	return rows, nil
}

// Close ferme la connexion au pool PostGIS
func (p *PostGISClient) Close() error {
	if p.Pool != nil {
		p.Pool.Close()
		p.Logger.Info("Connexion PostGIS fermée avec succès")
		return nil
	}
	return ErrClosedPool
}
