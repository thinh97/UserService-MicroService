package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	UserName string        `bson:"username" json:"username"`
	Password string        `bson:"password" json:"password"`
	DOB      string        `bson:"dob" json:"dob"`
	FullName string        `bson:"fullname" json:"fullname"`
}

type UserSignIn struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	IsSuccess bool   `json:"success"`
	Token     string `json:"token"`
}

type UserInfo struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	UserName string        `bson:"username" json:"username"`
	DOB      string        `bson:"dob" json:"dob"`
	FullName string        `bson:"fullname" json:"fullname"`
}
