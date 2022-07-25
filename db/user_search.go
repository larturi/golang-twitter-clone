package db

import (
	"context"
	"fmt"
	"time"

	"github.com/larturi/golang-twitter-clone/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UsersSearch(ID string, page int64, search string, tipo string) ([]*models.User, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(GetDBName())
	col := db.Collection("users")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var encontrado, incluir bool

	for cur.Next(ctx) {
		var s models.User
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var relation models.Relation
		relation.UserID = ID
		relation.UserRelationID = s.ID.Hex()

		incluir = false

		encontrado, _ = RelationCheck(relation)

		// Si no esta siguiendo al usuario de la iteracion, setea incluir en true
		if tipo == "new" && !encontrado {
			incluir = true
		}

		// Si esta siguiendo al usuario de la iteracion, setea incluir en true
		if tipo == "follow" && encontrado {
			incluir = true
		}

		// Para no seguirme a mi mismo
		if relation.UserRelationID == ID {
			incluir = false
		}

		if incluir {
			s.Password = ""
			s.Biography = ""
			s.WebSite = ""
			s.City = ""
			s.Banner = ""

			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	cur.Close(ctx)

	return results, true

}
