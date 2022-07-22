package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/larturi/golang-twitter-clone/db"
	"github.com/larturi/golang-twitter-clone/models"
)

func UserAvatarRouter(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")
	var extension string = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/avatars/" + IDUser + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "No se ha logrado subir la imagen"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen en el servidor"+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Avatar = IDUser + "." + extension
	status, err = db.UpdateUserProfile(user, IDUser)

	if err != nil || !status {
		http.Error(w, "Error al cgrabar el avatar en la BD"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
