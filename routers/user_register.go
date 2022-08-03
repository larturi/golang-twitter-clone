package routers

import (
	"encoding/json"
	"net/http"

	"github.com/larturi/golang-twitter-clone/db"
	"github.com/larturi/golang-twitter-clone/models"
)

func RegisterRouter(w http.ResponseWriter, r *http.Request) {

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validaciones

	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "El password debe tener al menos 6 caracteres", http.StatusBadRequest)
		return
	}

	_, userExist, _ := db.CheckUserExists(t.Email)
	if userExist {
		http.Error(w, "Ya existe un usuario registrado con ese email", http.StatusBadRequest)
		return
	}

	_, status, err := db.UserInsert(t)
	if err != nil {
		http.Error(w, "Error al intentar crear el usuario", http.StatusBadRequest)
	}

	if !status {
		http.Error(w, "No se ha logrado dar de alta el usuario", http.StatusBadRequest)

	}

	// Se creo el usuario
	w.WriteHeader(http.StatusCreated)

}
