package routers

import (
	"../controllers"
	"../middleware"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func SetUserRouter(r *mux.Router) *mux.Router {
	userRouter := mux.NewRouter()
	userRouter.HandleFunc("/user", controllers.UserInfo).Methods("GET")
	userRouter.HandleFunc("/user", controllers.UpdateUser).Methods("PUT")
	userRouter.HandleFunc("/user", controllers.DeleteUser).Methods("DELETE")
	r.Handle("/user", negroni.New(
		negroni.HandlerFunc(middleware.Authorize),
		negroni.Wrap(userRouter),
	))

	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	return r
}
