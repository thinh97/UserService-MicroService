package router

import (
	"../controller"
	"github.com/gorilla/mux"
)

func InitRouterUser(r *mux.Router) {
	r.HandleFunc("/user", controller.UserInfo).Methods("GET")
	r.HandleFunc("/user", controller.SignUp_SignIn_SignOut).Methods("POST")
	r.HandleFunc("/user", controller.UpdateUser).Methods("PUT")
	r.HandleFunc("/user", controller.DeleteUser).Methods("DELETE")
}
