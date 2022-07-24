package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/larturi/golang-twitter-clone/db"
)

func UserSearchRouter(w http.ResponseWriter, r *http.Request) {

	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parámetro page como entero mayor a cero", http.StatusBadRequest)
		return
	}

	pag := int64(pageTemp)

	results, status := db.UsersSearch(IDUser, pag, search, typeUser)
	if !status {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)

}
