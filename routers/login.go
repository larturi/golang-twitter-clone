package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/larturi/golang-twitter-clone/db"
	"github.com/larturi/golang-twitter-clone/jwt"
	"github.com/larturi/golang-twitter-clone/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario y/o password incorrectos: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido", http.StatusBadRequest)
		return
	}

	document, exists := db.Login(t.Email, t.Password)
	if !exists {
		http.Error(w, "Usuario y/o password incorrectos", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.JWTGenerate(document)
	if err != nil {
		http.Error(w, "Error al crear JWT: "+err.Error(), http.StatusBadRequest)
		return
	}

	response := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	expirationTime := time.Now().Add(24 * time.Hour)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
