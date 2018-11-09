package userdao

import (
	"../models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database

const COLLECTION = "Users"

func InitDb(session *mgo.Session) {
	db = session.DB(COLLECTION)
}

func FindByUserNameAndPassword(signInData model.UserSignIn) (bool, error) {
	var user model.User
	err := db.C(COLLECTION).Find(bson.M{"username": signInData.UserName, "password": signInData.Password}).One(&user)
	if err == nil {
		return true, nil
	}
	return false, err
}

func FindByUserName(username string) (bool, error) {
	var user model.User
	err := db.C(COLLECTION).Find(bson.M{"username": username}).One(&user)
	if err != nil {
		return true, err
	}
	return false, err
}

func Insert(user model.User) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}

func Delete(user model.User) error {
	err := db.C(COLLECTION).Remove(&user)
	return err
}

func Update(user model.User) error {
	err := db.C(COLLECTION).UpdateId(user.ID, &user)
	return err
}
