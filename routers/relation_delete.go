package routers

import (
	"net/http"

	"github.com/larturi/golang-twitter-clone/db"
	"github.com/larturi/golang-twitter-clone/models"
)

func RelationDeleteRouter(w http.ResponseWriter, r *http.Request) {

	IDUserFollow := r.URL.Query().Get("id")
	if len(IDUserFollow) < 1 {
		http.Error(w, "El parametro id es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = IDUserFollow

	status, err := db.RelationDelete(t)
	if err != nil {
		http.Error(w, "Error al intentar eliminar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "Error al intentar eliminar la relacion", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}
