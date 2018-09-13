package routes

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/golang_workspace/authentication/controllers"
	"github.com/golang_workspace/authentication/flash"
	"github.com/golang_workspace/authentication/models"
	"github.com/julienschmidt/httprouter"
)

// GtSignup will listen to GET on signup route
func GtSignup(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	_, err := Authenticate(req)
	if err == nil {
		http.Redirect(w, req, "/user/admin", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.html", nil)
}

// PstSignup will listen to POST on signup route
func PstSignup(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	userSingup := models.UserDB{
		Fname:    req.FormValue("fname"),
		Lname:    req.FormValue("lname"),
		Email:    req.FormValue("email"),
		Password: req.FormValue("password01"),
		UserType: "local",
	}
	str := checkSignup(req)

	alerts := models.UserAlerts{}
	if str != "" {
		alerts.ErrorMessage = str
		tpl.ExecuteTemplate(w, "signup.html", alerts)
		return
	}
	err := controllers.CreateUser(userSingup)
	if err != nil {
		if err.Error() == "user already exists" {
			alerts.ErrorMessage = "User already exists."
		} else {
			alerts.ErrorMessage = err.Error()
		}
		tpl.ExecuteTemplate(w, "signup.html", alerts)
		return
	}
	flash.SetFlash(w, "success", []byte("User created you can now login."))
	http.Redirect(w, req, "/user/login", http.StatusSeeOther)
}

func checkSignup(req *http.Request) string {
	if len(req.FormValue("fname")) < 3 {
		return "First name should be atleast three characters long."
	} else if len(req.FormValue("lname")) < 3 {
		return "Last name should be atleast three characters long."
	} else if !govalidator.IsEmail(req.FormValue("email")) {
		return "Email is invalid."
	} else if req.FormValue("password01") == "" || req.FormValue("password02") == "" {
		return "Both password fields are required."
	} else if len(req.FormValue("password01")) < 6 {
		return "Password should be atleast 6 characters long."
	} else if req.FormValue("password01") != req.FormValue("password02") {
		return "Both passwords should match."
	}
	return ""
	// Todo: else if !govalidator.Matches(u.Password, passRegExp) {
	// 	return "Password should be atleast 6 characters long and should include atleast one uppercase letter or numeric character."
	// }
}
