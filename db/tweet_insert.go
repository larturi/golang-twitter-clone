package db

import (
	"context"
	"time"

	"github.com/larturi/golang-twitter-clone/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TweetInsert(t models.TweetSave) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database(GetDBName())
	col := db.Collection("tweet")

	newTweet := bson.M{
		"user_id":    t.UserId,
		"message":    t.Message,
		"created_at": t.CreatedAt,
	}

	result, err := col.InsertOne(ctx, newTweet)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}
