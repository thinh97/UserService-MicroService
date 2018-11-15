package models

import jwt "github.com/dgrijalva/jwt-go"

type AppClaims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}
