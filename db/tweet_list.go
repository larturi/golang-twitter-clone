package db

import (
	"context"
	"log"
	"time"

	"github.com/larturi/golang-twitter-clone/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TweetsList(ID string, page int64) ([]*models.Tweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database(GetDBName())
	col := db.Collection("tweet")

	var results []*models.Tweets

	filter := bson.M{"user_id": ID}

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "created_at", Value: -1}})
	opciones.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, filter, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var registro models.Tweets
		err := cursor.Decode(&registro)

		if err != nil {
			return results, false
		}

		results = append(results, &registro)
	}

	return results, true
}
