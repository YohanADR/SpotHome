package jwt

import (
	"context"
	"fmt"
	"time"

	"github.com/YohanADR/SpotHome/infrastructure/db/postgis"
	"github.com/YohanADR/SpotHome/infrastructure/logger" // Utilisation de la base de données pour stocker les refresh tokens
	"github.com/YohanADR/SpotHome/pkg/errors"
	"github.com/YohanADR/SpotHome/pkg/events"
	"github.com/YohanADR/SpotHome/pkg/messaging"
	"github.com/golang-jwt/jwt/v4"
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

// GenerateToken génère un access token et un refresh token
func (j *JWTService) GenerateToken(username string, userID int) (string, string, error) {
	// Créer les claims pour l'access token (expirant dans 15 minutes)
	accessTokenClaims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		j.Logger.Error("Erreur lors de la génération de l'access token JWT", "error", err)
		return "", "", errors.New(500, "token_generation_failed", "Erreur lors de la génération de l'access token JWT")
	}

	// Générer un refresh token
	refreshTokenString, err := j.generateRefreshToken(userID)
	if err != nil {
		j.Logger.Error("Erreur lors de la génération du refresh token", "error", err)
		return "", "", errors.New(500, "refresh_token_generation_failed", "Erreur lors de la génération du refresh token")
	}

	j.Logger.Info("Access et refresh tokens générés avec succès", "username", username)
	return accessTokenString, refreshTokenString, nil
}

// generateRefreshToken génère un refresh token, l'enregistre en base de données, et retourne le token
func (j *JWTService) generateRefreshToken(userID int) (string, error) {
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenString, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	// Enregistrer le refresh token dans la base de données
	expiresAt := time.Now().Add(time.Hour * 24 * 7) // Expire dans 7 jours
	query := "INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES ($1, $2, $3)"
	_, err = j.DB.Pool.Exec(context.Background(), query, userID, refreshTokenString, expiresAt)
	if err != nil {
		return "", err
	}

	return refreshTokenString, nil
}

// ValidateToken valide un token JWT et retourne les claims si le token est valide
func (j *JWTService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		j.Logger.Error("Erreur lors de la validation du token JWT", "error", err)
		event := events.Event{
			Name:    "JWTValidationFailed",
			Payload: map[string]interface{}{"token": tokenString, "error": err},
		}
		j.emitEvent(event)
		return nil, errors.New(500, "token_validation_failed", "Le token JWT est invalide ou expiré")
	}

	j.Logger.Info("Token JWT validé avec succès", "token", tokenString)
	event := events.Event{
		Name:    "JWTValidated",
		Payload: map[string]interface{}{"token": tokenString},
	}
	j.emitEvent(event)

	return token, nil
}

// ValidateRefreshToken vérifie si le refresh token est valide en base de données
func (j *JWTService) ValidateRefreshToken(refreshToken string) (bool, error) {
	query := "SELECT COUNT(*) FROM refresh_tokens WHERE token = $1 AND expires_at > NOW()"
	var count int
	err := j.DB.Pool.QueryRow(context.Background(), query, refreshToken).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// emitEvent émet un événement via Kafka
func (j *JWTService) emitEvent(event events.Event) {
	eventMessage := fmt.Sprintf("Event: %s, Payload: %v", event.Name, event.Payload)
	err := j.KafkaProducer.Produce(context.Background(), "events-topic", eventMessage)
	if err != nil {
		j.Logger.Error("Erreur lors de l'envoi de l'événement à Kafka", "event", event.Name, "error", err)
	} else {
		j.Logger.Info("Événement envoyé à Kafka avec succès", "event", event.Name)
	}
}
