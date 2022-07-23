package db

import (
	"context"
	"fmt"
	"time"

	"github.com/larturi/golang-twitter-clone/models"
	"go.mongodb.org/mongo-driver/bson"
)

func RelationCheck(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database(GetDBName())
	col := db.Collection("relation")

	condition := bson.M{
		"user_id":          t.UserID,
		"user_relation_id": t.UserRelationID,
	}

	var resultado models.Relation

	fmt.Println(resultado)

	err := col.FindOne(ctx, condition).Decode(&resultado)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
