package middlewares

import (
	"net/http"

	"github.com/larturi/golang-twitter-clone/db"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckDB() == 0 {
			http.Error(w, "Sin conexion a la base de datos", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}
