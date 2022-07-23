package routers

import (
	"encoding/json"
	"net/http"

	"github.com/larturi/golang-twitter-clone/db"
	"github.com/larturi/golang-twitter-clone/models"
)

func RelationCheckRouter(w http.ResponseWriter, r *http.Request) {

	IDUserFollow := r.URL.Query().Get("id")
	if len(IDUserFollow) < 1 {
		http.Error(w, "El parametro id es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = IDUserFollow

	var resp models.ResponseRelation

	status, err := db.RelationCheck(t)
	if err != nil || !status {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}
