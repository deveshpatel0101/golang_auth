package routes

import (
	"net/http"

	"github.com/golang_workspace/authentication/flash"
	"github.com/golang_workspace/authentication/models"
	"github.com/julienschmidt/httprouter"
)

// GtAdmin will listen to GET on admin route
func GtAdmin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	alerts := models.UserAlerts{}
	ui, err := Authenticate(req)
	if err != nil {
		http.Redirect(w, req, "/user/login", http.StatusSeeOther)
		return
	}

	val, _ := flash.GetFlash(w, req, "success")
	if string(val) != "" {
		alerts.SuccessMessage = string(val)
	}

	alerts.User = ui
	tpl.ExecuteTemplate(w, "admin.html", alerts)
}
