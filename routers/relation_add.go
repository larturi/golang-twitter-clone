package routers

import (
	"net/http"

	"github.com/larturi/golang-twitter-clone/db"
	"github.com/larturi/golang-twitter-clone/models"
)

func RelationAddRouter(w http.ResponseWriter, r *http.Request) {

	IDUserFollow := r.URL.Query().Get("id")
	if len(IDUserFollow) < 1 {
		http.Error(w, "El parametro id es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = IDUserFollow

	status, err := db.RelationAdd(t)
	if err != nil {
		http.Error(w, "Error al intentar crear la relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "Error al intentar crear la relacion", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
