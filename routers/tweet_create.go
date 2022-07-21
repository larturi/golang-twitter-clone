package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/larturi/golang-twitter-clone/db"
	"github.com/larturi/golang-twitter-clone/models"
)

func TweetCreate(w http.ResponseWriter, r *http.Request) {

	var message models.Tweet
	json.NewDecoder(r.Body).Decode(&message)

	newTweet := models.TweetSave{
		UserId:    IDUser,
		Message:   message.Message,
		CreatedAt: time.Now(),
	}

	_, status, err := db.TweetInsert(newTweet)
	if err != nil {
		http.Error(w, "Error al intentar crear el tweet"+err.Error(), http.StatusBadRequest)
	}

	if !status {
		http.Error(w, "Error al intentar crear el tweet", http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)

}
