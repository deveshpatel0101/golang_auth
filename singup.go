package main

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/golang_workspace/authentication/controllers"
	"github.com/golang_workspace/authentication/flash"
	"github.com/golang_workspace/authentication/models"
	"github.com/julienschmidt/httprouter"
)

type singup struct {
	Fname     string
	Lname     string
	Email     string
	Password  string
	CPassword string
}

func gtSignup(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	_, err := authenticate(req)
	if err == nil {
		http.Redirect(w, req, "/admin", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.html", nil)
}

func pstSignup(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	userSingup := singup{
		Fname:     req.FormValue("fname"),
		Lname:     req.FormValue("lname"),
		Email:     req.FormValue("email"),
		Password:  req.FormValue("password01"),
		CPassword: req.FormValue("password02"),
	}
	str := checkSignup(userSingup)

	alerts := models.UserAlerts{}
	if str != "" {
		alerts.ErrorMessage = str
		tpl.ExecuteTemplate(w, "signup.html", alerts)
		return
	}
	err := controllers.CreateUser(convert(userSingup))
	if err != nil {
		alerts.ErrorMessage = err.Error()
		tpl.ExecuteTemplate(w, "signup.html", alerts)
		return
	}
	flash.SetFlash(w, "success", []byte("User created you can now login."))
	http.Redirect(w, req, "/login", http.StatusSeeOther)
}

func checkSignup(u singup) string {
	if len(u.Fname) < 3 {
		return "First name should be atleast three characters long."
	} else if len(u.Lname) < 3 {
		return "Last name should be atleast three characters long."
	} else if !govalidator.IsEmail(u.Email) {
		return "Email is invalid."
	} else if u.Password == "" || u.CPassword == "" {
		return "Both password fields are required."
	} else if len(u.Password) < 6 {
		return "Password should be atleast 6 characters long."
	} else if u.Password != u.CPassword {
		return "Both passwords should match."
	}
	return ""
	// Todo: else if !govalidator.Matches(u.Password, passRegExp) {
	// 	return "Password should be atleast 6 characters long and should include atleast one uppercase letter or numeric character."
	// }
}

func convert(s singup) models.UserDB {
	return models.UserDB{
		Fname:    s.Fname,
		Lname:    s.Lname,
		Email:    s.Email,
		Password: s.Password,
		UserType: "local",
	}
}
