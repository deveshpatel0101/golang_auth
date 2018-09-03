package main

import (
	"net/http"

	"./controllers"
	"./models"
)

func authenticate(req *http.Request) (models.UserDB, error) {
	c, err := req.Cookie("sid")
	if err != nil {
		return models.UserDB{}, err
	}
	userID, err := controllers.GetSession(c.Value)
	if err != nil {
		return models.UserDB{}, err
	}
	ui, err := controllers.GetUserInfo(userID.Hex())
	if err != nil {
		return models.UserDB{}, err
	}
	return ui, nil
}
