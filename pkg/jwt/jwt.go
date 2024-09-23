package jwt

import (
	"github.com/YohanADR/SpotHome/infrastructure/db/postgis"
	"github.com/YohanADR/SpotHome/infrastructure/logger" // Utilisation de la base de données pour stocker les refresh tokens
	"github.com/YohanADR/SpotHome/pkg/messaging"
)

var jwtSecret = []byte("YourSecretKey")

// JWTService struct
type JWTService struct {
	Logger        logger.Logger
	KafkaProducer messaging.MessageProducer
	DB            *postgis.PostGISClient // Ajoute la connexion à la base de données
}

// NewJWTService initialise le service JWT avec le logger et la base de données
func NewJWTService(log logger.Logger, kafkaProducer messaging.MessageProducer, db *postgis.PostGISClient) *JWTService {
	return &JWTService{
		Logger:        log,
		KafkaProducer: kafkaProducer,
		DB:            db,
	}
}
