package routes

import (
	"net/http"

	"github.com/golang_workspace/authentication/flash"

	"github.com/golang_workspace/authentication/controllers"
	"github.com/golang_workspace/authentication/models"

	"github.com/julienschmidt/httprouter"
)

// GtReset will listen for GET on /reset
func GtReset(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	id := req.FormValue("id")
	alerts := models.UserAlerts{}
	_, err := controllers.GetReset(id)
	if err != nil {
		alerts.ErrorMessage = "Either token was already used or the reset password link is broken."
		tpl.ExecuteTemplate(w, "reset.html", alerts)
		return
	}

	tpl.ExecuteTemplate(w, "reset.html", nil)
}

// PstReset will list for POST on /reset
func PstReset(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	id := req.FormValue("id")
	alerts := models.UserAlerts{}
	ur, err := controllers.GetReset(id)
	if err != nil {
		alerts.ErrorMessage = "Either token was already used or the reset password link is broken."
		tpl.ExecuteTemplate(w, "reset.html", alerts)
		return
	}

	str := checkPassword(req)
	if str != "" {
		alerts.ErrorMessage = str
		tpl.ExecuteTemplate(w, "reset.html", alerts)
		return
	}

	err = controllers.UpdatePassword(ur.Email, req.FormValue("password01"))
	if err != nil {
		if err.Error() == "not found" {
			alerts.ErrorMessage = "User does not exists."
		} else {
			alerts.ErrorMessage = err.Error()
		}
		tpl.ExecuteTemplate(w, "reset.html", alerts)
		return
	}
	flash.SetFlash(w, "success", []byte("Your password was reset successfully you can now login."))
	http.Redirect(w, req, "/user/login", http.StatusSeeOther)
}

func checkPassword(req *http.Request) string {
	if req.FormValue("password01") == "" || req.FormValue("password02") == "" {
		return "Both passwords are required."
	} else if len(req.FormValue("password01")) < 6 {
		return "Password should be atleast 6 characters long."
	} else if !(req.FormValue("password01") == req.FormValue("password02")) {
		return "Both passwords should match."
	}
	return ""
}
