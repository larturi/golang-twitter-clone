package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/larturi/golang-twitter-clone/middlewares"
	"github.com/larturi/golang-twitter-clone/routers"
)

func Handlers() {
	router := mux.NewRouter()

	// Auth Routes
	router.HandleFunc("/register", middlewares.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middlewares.CheckDB(middlewares.ValidateJWT(routers.ViewUserProfile))).Methods("GET")
	router.HandleFunc("/profile", middlewares.CheckDB(middlewares.ValidateJWT(routers.UpdateUserProfile))).Methods("PUT")

	// Tweet Routes
	router.HandleFunc("/tweet", middlewares.CheckDB(middlewares.ValidateJWT(routers.TweetCreate))).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8083"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
