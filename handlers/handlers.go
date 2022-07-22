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
	router.HandleFunc("/register", middlewares.CheckDB(routers.RegisterRouter)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDB(routers.LoginRouter)).Methods("POST")
	router.HandleFunc("/profile", middlewares.CheckDB(middlewares.ValidateJWT(routers.ViewUserProfileRouter))).Methods("GET")
	router.HandleFunc("/profile", middlewares.CheckDB(middlewares.ValidateJWT(routers.UpdateUserProfileRouter))).Methods("PUT")

	// Tweet Routes
	router.HandleFunc("/tweet", middlewares.CheckDB(middlewares.ValidateJWT(routers.TweetCreateRouter))).Methods("POST")
	router.HandleFunc("/tweets", middlewares.CheckDB(middlewares.ValidateJWT(routers.TweetsListRouter))).Methods("GET")
	router.HandleFunc("/tweet", middlewares.CheckDB(middlewares.ValidateJWT(routers.TweetDeleteRouter))).Methods("DELETE")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8083"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
