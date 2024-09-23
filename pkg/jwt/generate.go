package jwt

import (
	"fmt"
	"strings"
	"time"

	"github.com/YohanADR/SpotHome/pkg/errors"
	"github.com/YohanADR/SpotHome/pkg/events"
	"github.com/golang-jwt/jwt/v4"
)

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

// ValidateToken valide un token JWT et retourne les claims si le token est valide
func (j *JWTService) ValidateToken(tokenString string) (*jwt.Token, error) {
	// Supprimer les éventuels préfixes "Bearer " dans le token string
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Parser le token JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Vérifier que la méthode de signature du token est bien la méthode attendue
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			j.Logger.Error("Méthode de signature inattendue", "token", tokenString)
			return nil, fmt.Errorf("méthode de signature JWT inattendue")
		}
		return jwtSecret, nil
	})

	// Gestion des erreurs de parsing ou si le token est invalide
	if err != nil || !token.Valid {
		j.Logger.Error("Erreur lors de la validation du token JWT", "error", err)
		event := events.Event{
			Name:    "JWTValidationFailed",
			Payload: map[string]interface{}{"token": tokenString, "error": err},
		}
		events.EmitEvent(event)
		return nil, errors.New(401, "token_validation_failed", "Le token JWT est invalide ou expiré")
	}

	// Si le token est valide, loguer l'événement et retourner le token
	j.Logger.Info("Token JWT validé avec succès", "token", tokenString)
	event := events.Event{
		Name:    "JWTValidated",
		Payload: map[string]interface{}{"token": tokenString},
	}
	events.EmitEvent(event)

	return token, nil
}
