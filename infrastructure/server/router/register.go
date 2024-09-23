package router

import (
	"net/http"

	"github.com/YohanADR/SpotHome/pkg/jwt"
	"github.com/YohanADR/SpotHome/pkg/jwt/middleware"
	"github.com/YohanADR/SpotHome/pkg/transport"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(appRouter *Router, jwtService *jwt.JWTService) {
	appRouter.RegisterRoutes(func(register transport.RegisterRoutes) {
		// Route publique
		register("GET", "/health", gin.HandlerFunc(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "OK"})
		}))

		// Route protégée avec le middleware JWT
		register("GET", "/protected", gin.HandlerFunc(func(c *gin.Context) {
			middleware.JWTMiddleware(jwtService)(c) // Appliquer le middleware
			if c.IsAborted() {
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Protected route access granted"})
		}))

		// Route pour générer un token
		register("POST", "/generate-token", gin.HandlerFunc(func(c *gin.Context) {
			// Récupérer le nom d'utilisateur à partir de la requête JSON
			var requestBody struct {
				Username string `json:"username"`
			}

			if err := c.ShouldBindJSON(&requestBody); err != nil || requestBody.Username == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request. Username is required."})
				return
			}

			// Générer le token et le refresh token
			token, refreshToken, err := jwtService.GenerateToken(requestBody.Username, 1)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
				return
			}

			// Retourner le token et le refresh token générés
			c.JSON(http.StatusOK, gin.H{
				"token":        token,
				"refreshToken": refreshToken,
			})
		}))
	})
}
