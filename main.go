package main

import (
	"log"
	"net/http"

	"./db"
	"./router"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	r := mux.NewRouter()
	router.InitRouterUser(r)

	session, err := mgo.Dial("mongodb://admin:123456789x@ds157493.mlab.com:57493/microservice")
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	userdao.InitDb(session)

	if err := http.ListenAndServe(":3001", r); err != nil {
		log.Fatal(err)
	}
}
