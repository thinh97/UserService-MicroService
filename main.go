package main

import (
	"log"
	"net/http"

	"./config"
	"./repository"
	"./routers"
	"github.com/urfave/negroni"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	appConfig, err := config.GetConfig()
	if err != nil {
		log.Fatal("Load config error: ", err)
	}
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	// mongoDBDialInfo := &mgo.DialInfo{
	// 	Addrs:    []string{appConfig.MongoDbHost},
	// 	Timeout:  60 * time.Second,
	// 	Database: appConfig.DatabaseName,
	// 	Username: appConfig.DbUserName,
	// 	Password: appConfig.DbPassword,
	// }
	dialInfo, err0 := mgo.ParseURL(appConfig.ConnectionString)
	if err0 != nil {
		log.Fatal("Parse connection string error: ", err)
	}

	session, err := mgo.DialWithInfo(dialInfo)
	defer session.Close()
	if err != nil {
		log.Fatal("Connect DB error: ", err)
	}
	userRepository.InitDb(session)

	server := &http.Server{
		Addr:    appConfig.Host + ":" + appConfig.Port,
		Handler: n,
	}
	log.Println("Server started...")
	server.ListenAndServe()
}
