package routers

import (
	"encoding/json"
	"net/http"

	"github.com/larturi/golang-twitter-clone/db"
	"github.com/larturi/golang-twitter-clone/models"
)

func ViewUserProfileRouter(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	perfil, err := db.GetUserProfile(ID)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar buscar el registro "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}

func UpdateUserProfileRouter(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Datos incorrectos para actualizar el perfil "+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool
	status, err = db.UpdateUserProfile(t, IDUser)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar actualizar el perfil "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "Ourrió un error al intentar actualizar el perfil "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
