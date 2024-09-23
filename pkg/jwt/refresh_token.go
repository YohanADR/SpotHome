package jwt

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

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
