package models

type AppConfig struct {
	Host             string `json:"host"`
	Port             string `json:"port"`
	MongoDbHost      string `json:"mongohost"`
	DatabaseName     string `json:"Database"`
	DbUserName       string `json:"username"`
	DbPassword       string `json:"password"`
	ConnectionString string `json:"connectionstring"`
}
