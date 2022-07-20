package db

import (
	"context"
	"time"

	"github.com/larturi/golang-twitter-clone/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database(GetDBName())
	col := db.Collection("users")

	var perfil models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		return perfil, err
	}
	return perfil, nil
}

func UpdateUserProfile(user models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database(GetDBName())
	col := db.Collection("users")

	registro := make(map[string]interface{})

	if len(user.Name) > 0 {
		registro["name"] = user.Name
	}
	if len(user.LastName) > 0 {
		registro["last_name"] = user.LastName
	}
	if len(user.Avatar) > 0 {
		registro["avatar"] = user.Avatar
	}
	if len(user.Banner) > 0 {
		registro["banner"] = user.Banner
	}
	if len(user.Biography) > 0 {
		registro["biography"] = user.Biography
	}
	if len(user.City) > 0 {
		registro["city"] = user.City
	}
	if len(user.WebSite) > 0 {
		registro["web_site"] = user.WebSite
	}
	registro["birth_date"] = user.BirthDate

	updateString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updateString)

	if err != nil {
		return false, err
	}

	return true, nil
}
