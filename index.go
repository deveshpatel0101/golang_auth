package main

import (
	"net/http"

	"github.com/golang_workspace/authentication/models"

	"github.com/julienschmidt/httprouter"
)

func gtIndex(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	alerts := models.UserAlerts{}
	_, err := authenticate(req)
	if err != nil {
		alerts.LoggedIn = false
		tpl.ExecuteTemplate(w, "index.html", alerts)
		return
	}
	alerts.LoggedIn = true
	tpl.ExecuteTemplate(w, "index.html", alerts)
}
