package routers

import (
	"net/http"

	"github.com/larturi/golang-twitter-clone/db"
)

func TweetDeleteRouter(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id del tweet", http.StatusBadRequest)
		return
	}

	err := db.TweetDelete(ID, IDUser)
	if err != nil {
		http.Error(w, "Error al ententar eliminar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
