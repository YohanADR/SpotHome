package handlers

import (
	postgisPort "SpotHome/internal/infrastructure/postgis/ports"
	redisPort "SpotHome/internal/infrastructure/redis/ports"
	"encoding/json"
	"net/http"
	"strconv"
)

// User représente la structure d'un utilisateur
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// GetUserHandler gère la récupération des données d'un utilisateur
func GetUserHandler(postgres postgisPort.PostGISPort, redis redisPort.RedisPort) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("id")
		id, err := strconv.Atoi(userID)
		if err != nil {
			http.Error(w, "ID invalide", http.StatusBadRequest)
			return
		}

		// Utilise QueryRow pour obtenir l'utilisateur
		row := postgres.QueryRow("SELECT id, name, email FROM \"users\" WHERE id=$1", id)

		var user User // Assurez-vous que User est défini avec les bons champs
		if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			http.Error(w, "Erreur lors de la lecture des données de l'utilisateur", http.StatusInternalServerError)
			return
		}

		// Sérialisation et réponse
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
