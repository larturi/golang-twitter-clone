package db

import (
	"context"
	"time"

	"github.com/larturi/golang-twitter-clone/models"
)

func RelationDelete(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database(GetDBName())
	col := db.Collection("relation")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
