package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/larturi/golang-twitter-clone/models"
)

func JWTGenerate(t models.User) (string, error) {

	secret := []byte("asdfg12345$@!")

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Name,
		"apellidos":        t.LastName,
		"fecha_nacimiento": t.BirthDate,
		"biografia":        t.Biography,
		"ubicacion":        t.City,
		"sitioweb":         t.WebSite,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(secret)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil

}
