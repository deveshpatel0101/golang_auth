package routes

import (
	"net/http"

	"github.com/golang_workspace/authentication/flash"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang_workspace/authentication/controllers"

	"github.com/golang_workspace/authentication/models"

	"github.com/julienschmidt/httprouter"
)

// GtChangePwd will listen for GET on /settings
func GtChangePwd(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	_, err := Authenticate(req)
	if err != nil {
		http.Redirect(w, req, "/user/login", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "changepwd.html", nil)
}

// PstChangePwd will listen for POST on /settings
func PstChangePwd(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ui, err := Authenticate(req)
	if err != nil {
		http.Redirect(w, req, "/user/login", http.StatusSeeOther)
		return
	}

	alerts := models.UserAlerts{}
	str := changePassword(req)
	if str != "" {
		alerts.ErrorMessage = str
		tpl.ExecuteTemplate(w, "changepwd.html", alerts)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(ui.Password), []byte(req.FormValue("currPwd")))
	if err != nil {
		alerts.ErrorMessage = "Your current password is wrong."
		tpl.ExecuteTemplate(w, "changepwd.html", alerts)
		return
	}

	if req.FormValue("currPwd") == req.FormValue("password01") {
		alerts.ErrorMessage = "Your new password should not match old password."
		tpl.ExecuteTemplate(w, "changepwd.html", alerts)
		return
	}

	err = controllers.UpdatePassword(ui.Email, req.FormValue("password01"))
	if err != nil {
		if err.Error() == "not found" {
			alerts.ErrorMessage = "User does not exists."
		} else {
			alerts.ErrorMessage = err.Error()
		}
		tpl.ExecuteTemplate(w, "changepwd.html", alerts)
		return
	}

	flash.SetFlash(w, "success", []byte("Your password was updated successfully."))
	http.Redirect(w, req, "/user/admin", http.StatusSeeOther)
}

func changePassword(req *http.Request) string {
	if req.FormValue("currPwd") == "" || req.FormValue("password01") == "" || req.FormValue("password02") == "" {
		return "All fields are required."
	} else if len(req.FormValue("password01")) < 6 {
		return "New password should be at least 6 character long."
	} else if req.FormValue("password01") != req.FormValue("password02") {
		return "Confirm password does not match new password."
	}
	return ""
}
