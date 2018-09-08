package routes

import (
	"net/http"

	"github.com/golang_workspace/authentication/models"

	"github.com/julienschmidt/httprouter"
)

// GtIndex will to GET on / route
func GtIndex(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	alerts := models.UserAlerts{}
	_, err := Authenticate(req)
	if err != nil {
		alerts.LoggedIn = false
		tpl.ExecuteTemplate(w, "index.html", alerts)
		return
	}
	alerts.LoggedIn = true
	tpl.ExecuteTemplate(w, "index.html", alerts)
}
