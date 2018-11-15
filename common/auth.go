package common

import (
	"log"
	"time"

	"../models"
	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(userName string) (string, error) {
	appClaims := models.AppClaims{
		userName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, appClaims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (bool, *models.AppClaims) {
	appClaims := models.AppClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &appClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err == nil && token.Valid {
		return true, &appClaims
	}
	log.Println(err)
	return false, nil
}
