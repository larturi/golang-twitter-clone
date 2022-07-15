package main

import (
	"log"

	"github.com/larturi/golang-twitter-clone/db"
	"github.com/larturi/golang-twitter-clone/handlers"
)

func main() {
	if db.CheckDB() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Handlers()
}
