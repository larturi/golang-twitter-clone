package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TweetDelete(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database(GetDBName())
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	rules := bson.M{"_id": objID, "user_id": UserID}

	_, err := col.DeleteOne(ctx, rules)
	return err
}
