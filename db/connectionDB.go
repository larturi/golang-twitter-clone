package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectDB()
var clientOptions = options.Client().ApplyURI(getConnectionString())

/* ConectDB hace la conexion a mongodb */
func ConectDB() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Successful connection to database")
	return client
}

/* CheckDB hace un ping a mongodb */
func CheckDB() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}

func getConnectionString() string {
	godotenv.Load(".env")
	return os.Getenv("MONGO_CONNECTION")
}
