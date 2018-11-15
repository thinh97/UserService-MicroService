package config

import (
	"encoding/json"
	"os"

	"../models"
)

func GetConfig() (*models.AppConfig, error) {
	file, err := os.Open("config/config.json")
	defer file.Close()
	if err != nil {
		return nil, err
	}
	var appConfig *models.AppConfig
	if err := json.NewDecoder(file).Decode(&appConfig); err != nil {
		return nil, err
	}
	return appConfig, nil
}
