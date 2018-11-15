package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"../common"
	"../models"
	"../repository"
	"github.com/gorilla/mux"
)

func UserInfo(w http.ResponseWriter, r *http.Request) {
	userName := r.Context().Value("UserName")
	if userName != nil {
		user, err := userRepository.FindByUserName(userName.(string))
		if err != nil {
			common.RespondWithError(w, http.StatusBadRequest, "Invalid User ID")
		} else {
			common.RespondWithJson(w, http.StatusOK, user)
		}
		return
	}
	common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
}

func Login(w http.ResponseWriter, r *http.Request) {
	var userSign models.UserSignIn
	if err := json.NewDecoder(r.Body).Decode(&userSign); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user, err := userRepository.FindByUserNameAndPassword(userSign.UserName, userSign.Password)
	if err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}
	if user != nil {
		tokenString, err := common.GenerateToken(user.UserName)
		if err != nil {
			common.RespondWithError(w, http.StatusBadRequest, "Parse token fail")
			return
		}
		res := models.UserResponse{IsSuccess: true, Token: tokenString}
		common.RespondWithJson(w, http.StatusOK, res)
	} else {
		res := models.UserResponse{IsSuccess: false, Token: ""}
		common.RespondWithJson(w, http.StatusNotFound, res)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user, err := userRepository.FindByUserName(newUser.UserName)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if user == nil {
		if err := userRepository.Insert(&newUser); err != nil {
			common.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		common.RespondWithJson(w, http.StatusCreated, newUser)
	} else {
		err = errors.New("UserName already exists")
		common.RespondWithError(w, http.StatusConflict, err.Error())
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := userRepository.FindByUserName(params["id"])
	if err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}
	common.RespondWithJson(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	userName := r.Context().Value("UserName")
	user, err := userRepository.FindByUserName(userName.(string))
	if err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
	} else {
		newUser.ID = user.ID
		newUser.UserName = user.UserName
	}
	if err := userRepository.Update(newUser); err != nil {
		common.RespondWithJson(w, http.StatusInternalServerError, map[string]string{"success": "false"})
		return
	}
	common.RespondWithJson(w, http.StatusOK, map[string]string{"success": "true"})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := userRepository.Delete(user); err != nil {
		common.RespondWithJson(w, http.StatusInternalServerError, map[string]string{"success": "false"})
		return
	}
	common.RespondWithJson(w, http.StatusOK, map[string]string{"success": "true"})
}
