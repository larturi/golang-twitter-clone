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

	// Auth & User Routes
	router.HandleFunc("/register", middlewares.CheckDB(routers.RegisterRouter)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDB(routers.LoginRouter)).Methods("POST")
	router.HandleFunc("/profile", middlewares.CheckDB(middlewares.ValidateJWT(routers.ViewUserProfileRouter))).Methods("GET")
	router.HandleFunc("/profile", middlewares.CheckDB(middlewares.ValidateJWT(routers.UpdateUserProfileRouter))).Methods("PUT")
	router.HandleFunc("/userSearch", middlewares.CheckDB(middlewares.ValidateJWT(routers.UserSearchRouter))).Methods("GET")

	// Avatar & Banner Routes
	router.HandleFunc("/uploadAvatar", middlewares.CheckDB(middlewares.ValidateJWT(routers.UserSaveAvatarRouter))).Methods("POST")
	router.HandleFunc("/getAvatar", middlewares.CheckDB(routers.UserGetAvatarRouter)).Methods("GET")
	router.HandleFunc("/uploadBanner", middlewares.CheckDB(middlewares.ValidateJWT(routers.UserSaveBannerRouter))).Methods("POST")
	router.HandleFunc("/getBanner", middlewares.CheckDB(routers.UserGetBannerRouter)).Methods("GET")

	// Tweet Routes
	router.HandleFunc("/tweet", middlewares.CheckDB(middlewares.ValidateJWT(routers.TweetCreateRouter))).Methods("POST")
	router.HandleFunc("/tweets", middlewares.CheckDB(middlewares.ValidateJWT(routers.TweetsListRouter))).Methods("GET")
	router.HandleFunc("/tweetsFollowing", middlewares.CheckDB(middlewares.ValidateJWT(routers.TweetsListFollowingRouter))).Methods("GET")
	router.HandleFunc("/tweet", middlewares.CheckDB(middlewares.ValidateJWT(routers.TweetDeleteRouter))).Methods("DELETE")

	// Relations
	router.HandleFunc("/relationAdd", middlewares.CheckDB(middlewares.ValidateJWT(routers.RelationAddRouter))).Methods("POST")
	router.HandleFunc("/relationRemove", middlewares.CheckDB(middlewares.ValidateJWT(routers.RelationDeleteRouter))).Methods("DELETE")
	router.HandleFunc("/relationCkeck", middlewares.CheckDB(middlewares.ValidateJWT(routers.RelationCheckRouter))).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8083"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
