package routes

import (
	"net/http"

	"github.com/golang_workspace/authentication/controllers"
	"github.com/golang_workspace/authentication/flash"
	"github.com/golang_workspace/authentication/models"

	"github.com/julienschmidt/httprouter"
)

// GtLogin will listen to GET on login route
func GtLogin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	val, _ := flash.GetFlash(w, req, "success")
	if string(val) != "" {
		alerts := models.UserAlerts{
			SuccessMessage: string(val),
		}
		tpl.ExecuteTemplate(w, "login.html", alerts)
		return
	}
	_, err := Authenticate(req)
	if err == nil {
		http.Redirect(w, req, "/user/admin", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.html", nil)
}

// PstLogin will listen to POST on login route
func PstLogin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	str := checkLogin(req)
	uc := models.UserDB{
		Email:    req.FormValue("email"),
		Password: req.FormValue("password"),
	}
	alerts := models.UserAlerts{}
	if str != "" {
		alerts.ErrorMessage = str
		tpl.ExecuteTemplate(w, "login.html", alerts)
		return
	}
	user, err := controllers.ValidateUser(uc)
	if err != nil {
		alerts.ErrorMessage = "Email or Password is wrong."
		tpl.ExecuteTemplate(w, "login.html", alerts)
		return
	}
	us, err := controllers.CreateSession(user)
	if err != nil {
		alerts.ErrorMessage = err.Error()
		tpl.ExecuteTemplate(w, "login.html", alerts)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "sid",
		Value:    flash.Encode([]byte(us.UUID)),
		HttpOnly: true,
	})
	http.Redirect(w, req, "/user/admin", http.StatusSeeOther)
}

func checkLogin(req *http.Request) string {
	if req.FormValue("email") == "" {
		return "Email is required."
	} else if req.FormValue("password") == "" {
		return "Password is required."
	}
	return ""
}
