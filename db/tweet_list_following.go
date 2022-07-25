package db

import (
	"context"
	"time"

	"github.com/larturi/golang-twitter-clone/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Listado de tweets de las personas que estoy siguiendo ordenados por fecha desc
func TweetsListFollowing(ID string, page int64) ([]models.TweetsFollowing, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database(GetDBName())
	col := db.Collection("relation")

	skip := ((page - 1) * 20)

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"user_id": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "user_relation_id",
			"foreignField": "user_id",
			"as":           "tweet",
		},
	})
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweet.created_at": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursor, _ := col.Aggregate(ctx, conditions)
	var results []models.TweetsFollowing
	err := cursor.All(ctx, &results)
	if err != nil {
		return results, false
	}
	return results, true

}
