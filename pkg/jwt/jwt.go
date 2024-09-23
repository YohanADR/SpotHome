package jwt

import (
	"context"
	"fmt"
	"time"

	"github.com/YohanADR/SpotHome/infrastructure/logger"
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
}

// NewJWTService initialise le service JWT avec le logger et Kafka
func NewJWTService(log logger.Logger, kafkaProducer messaging.MessageProducer) *JWTService {
	return &JWTService{
		Logger:        log,
		KafkaProducer: kafkaProducer,
	}
}

// GenerateToken génère un token JWT avec des claims personnalisés et loggue les événements
func (j *JWTService) GenerateToken(username string) (string, error) {
	// Créer les claims pour le token
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	// Générer le token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		// Utiliser le logger pour enregistrer l'erreur
		j.Logger.Error("Erreur lors de la génération du token JWT", "error", err)

		// Émettre un événement pour signaler l'échec de la génération du token
		event := events.Event{
			Name:    "JWTGenerationFailed",
			Payload: map[string]interface{}{"username": username, "error": err},
		}
		j.emitEvent(event)

		// Retourner une erreur gérée par le système d'erreur
		return "", errors.New(500, "token_generation_failed", "Erreur lors de la génération du token JWT")
	}

	// Loguer l'événement de succès
	j.Logger.Info("Token JWT généré avec succès", "username", username)
	event := events.Event{
		Name:    "JWTGenerated",
		Payload: map[string]interface{}{"username": username},
	}
	j.emitEvent(event)

	return tokenString, nil
}

// ValidateToken valide un token JWT et retourne les claims si le token est valide
func (j *JWTService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		// Loguer l'erreur
		j.Logger.Error("Erreur lors de la validation du token JWT", "error", err)

		// Émettre un événement pour signaler l'échec de la validation
		event := events.Event{
			Name:    "JWTValidationFailed",
			Payload: map[string]interface{}{"token": tokenString, "error": err},
		}
		j.emitEvent(event)

		// Retourner une erreur gérée par le système d'erreur
		return nil, errors.New(500, "token_validation_failed", "Le token JWT est invalide ou expiré")
	}

	// Loguer l'événement de validation réussie
	j.Logger.Info("Token JWT validé avec succès", "token", tokenString)
	event := events.Event{
		Name:    "JWTValidated",
		Payload: map[string]interface{}{"token": tokenString},
	}
	j.emitEvent(event)

	return token, nil
}

// emitEvent émet un événement via Kafka
func (j *JWTService) emitEvent(event events.Event) {
	// Convertir l'événement en string ou en format sérialisé
	eventMessage := fmt.Sprintf("Event: %s, Payload: %v", event.Name, event.Payload)

	// Envoyer l'événement à Kafka via le producteur
	err := j.KafkaProducer.Produce(context.Background(), "events-topic", eventMessage)
	if err != nil {
		j.Logger.Error("Erreur lors de l'envoi de l'événement à Kafka", "event", event.Name, "error", err)
	} else {
		j.Logger.Info("Événement envoyé à Kafka avec succès", "event", event.Name)
	}
}
