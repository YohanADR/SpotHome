package middleware

import (
	"net/http"
	"strings"

	"github.com/YohanADR/SpotHome/pkg/jwt"
	"github.com/gin-gonic/gin"
)

// JWTMiddleware est un middleware qui vérifie la validité du token JWT
func JWTMiddleware(jwtService *jwt.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		// Vérification si l'en-tête Authorization est manquant
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Suppression du préfixe "Bearer " (attention aux majuscules)
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validation du token
		token, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
			c.Abort()
			return
		}

		// Si tout est OK, on continue la requête
		c.Set("token", token)
		c.Next()
	}
}
