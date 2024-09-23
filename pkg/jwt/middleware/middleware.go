package jwt

import (
	"net/http"
	"strings"

	"github.com/YohanADR/SpotHome/infrastructure/logger"
	"github.com/YohanADR/SpotHome/pkg/jwt"
	"github.com/YohanADR/SpotHome/pkg/messaging"
	"github.com/gin-gonic/gin"
)

// JWTMiddleware est un middleware qui vérifie la validité du token JWT
func JWTMiddleware(log logger.Logger, kafkaProducer messaging.MessageProducer) gin.HandlerFunc {
	jwtService := jwt.NewJWTService(log, kafkaProducer) // Créer une instance du service JWT avec Kafka

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwtService.ValidateToken(tokenString) // Appel à la méthode ValidateToken
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
