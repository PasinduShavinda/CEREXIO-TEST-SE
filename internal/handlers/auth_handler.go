package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/PasinduShavinda/go-project-api/internal/models"
	"github.com/dgrijalva/jwt-go"
)

var users = map[string]string{
	"rootadmin": "demo",
}

var jwtKey = []byte("secret_key")

// @Summary User Login
// @Description User Login
// @Tags user
// @Accept json
// @Produce json
// @Param credentials body models.Credentials true "Credentials Data"
// @Success 201 {object} models.Credentials
// @Failure 400 {object} map[string]string
// @Router /login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[credentials.Username]
	if !ok || expectedPassword != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &models.Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	// Log the token for debugging purposes
	fmt.Println("Token received:", tokenString)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
