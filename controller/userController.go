package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"../db"
	"../models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJson(w, code, map[string]string{"error": msg})
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func UserInfo(w http.ResponseWriter, r *http.Request) {
	user, err := userdao.FindByUserName("id")
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}
	RespondWithJson(w, http.StatusOK, user)
}

func SignUp_SignIn_SignOut(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if query["signin"] != nil {
		signIn(w, r)
	} else if query["signup"] != nil {
		signUp(w, r)
	} else if query["signout"] != nil {
		signOut(w, r)
	} else {
		RespondWithError(w, http.StatusBadRequest, "Invalid Query")
		return
	}
}

func signIn(w http.ResponseWriter, r *http.Request) {
	var user model.UserSignIn
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	isSuccess, err := userdao.FindByUserNameAndPassword(user)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}
	if isSuccess == true {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username":  user.UserName,
			"expiresAt": time.Now().Add(time.Minute * 10),
		})
		tokenString, err := token.SignedString([]byte("secret"))
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "Parse token fail")
			return
		}
		res := model.UserResponse{IsSuccess: true, Token: tokenString}
		RespondWithJson(w, http.StatusOK, res)
	} else {
		res := model.UserResponse{IsSuccess: false, Token: ""}
		RespondWithJson(w, http.StatusNotFound, res)
	}
}

func signUp(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user.ID = bson.NewObjectId()
	if err := userdao.Insert(user); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusCreated, user)
}

func signOut(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := userdao.FindByUserName(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}
	RespondWithJson(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := userdao.Update(user); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := userdao.Delete(user); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
