package db

import (
	"github.com/larturi/golang-twitter-clone/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email, password string) (models.User, bool) {
	user, userExists, _ := CheckUserExists(email)

	if !userExists {
		return user, false
	}

	passwordDB := []byte(user.Password)
	passwordBytes := []byte(password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return user, false
	}

	return user, true
}
