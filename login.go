package main

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/golang_workspace/authentication/controllers"
	"github.com/golang_workspace/authentication/flash"
	"github.com/golang_workspace/authentication/models"

	"github.com/julienschmidt/httprouter"
)

type login struct {
	Email    string
	Password string
}

func gtLogin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	val, _ := flash.GetFlash(w, req, "success")
	if string(val) != "" {
		alerts := models.UserAlerts{
			SuccessMessage: string(val),
		}
		tpl.ExecuteTemplate(w, "login.html", alerts)
		return
	}
	_, err := authenticate(req)
	if err == nil {
		http.Redirect(w, req, "/admin", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func pstLogin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	userLogin := login{
		Email:    req.FormValue("email"),
		Password: req.FormValue("password"),
	}
	str := checkLogin(userLogin)
	alerts := models.UserAlerts{}
	if str != "" {
		alerts.ErrorMessage = str
		tpl.ExecuteTemplate(w, "login.html", alerts)
		return
	}
	user, err := controllers.ValidateUser(models.UserDB{Email: userLogin.Email, Password: userLogin.Password})
	if err != nil {
		if err.Error() == "not found" {
			alerts.ErrorMessage = "Email or Password is wrong."
			tpl.ExecuteTemplate(w, "login.html", alerts)
			return
		} else if err.Error() == "wrong password" {
			alerts.ErrorMessage = "Email or Password is wrong."
			tpl.ExecuteTemplate(w, "login.html", alerts)
			return
		}
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
	http.Redirect(w, req, "/admin", http.StatusSeeOther)
}

func checkLogin(l login) string {
	if !govalidator.IsEmail(l.Email) {
		return "Invalid email."
	} else if l.Password == "" {
		return "Password is required."
	}
	return ""
}
