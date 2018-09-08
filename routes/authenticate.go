package routes

import (
	"net/http"

	"github.com/golang_workspace/authentication/controllers"
	"github.com/golang_workspace/authentication/models"
)

// Authenticate will check if user is logged in or not
func Authenticate(req *http.Request) (models.UserDB, error) {
	c, err := req.Cookie("sid")
	if err != nil {
		return models.UserDB{}, err
	}
	userID, err := controllers.GetSession(c.Value)
	if err != nil {
		return models.UserDB{}, err
	}
	ui, err := controllers.GetUserByID(userID.Hex())
	if err != nil {
		return models.UserDB{}, err
	}
	return ui, nil
}
