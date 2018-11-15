package userRepository

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

func FindByUserNameAndPassword(username string, password string) (*models.User, error) {
	var user *models.User
	err := db.C(COLLECTION).Find(bson.M{"username": username, "password": password}).One(&user)
	if err != nil && err != mgo.ErrNotFound {
		return nil, err
	}
	return user, nil
}

func FindByUserName(username string) (*models.UserInfo, error) {
	var user *models.UserInfo
	err := db.C(COLLECTION).Find(bson.M{"username": username}).One(&user)
	if err != nil && err != mgo.ErrNotFound {
		return nil, err
	}
	return user, nil
}

func Insert(user *models.User) error {
	user.ID = bson.NewObjectId()
	err := db.C(COLLECTION).Insert(&user)
	return err
}

func Delete(user models.User) error {
	err := db.C(COLLECTION).Remove(&user)
	return err
}

func Update(user models.User) error {
	err := db.C(COLLECTION).UpdateId(user.ID, &user)
	return err
}
