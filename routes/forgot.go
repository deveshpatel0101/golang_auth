package routes

import (
	"net/http"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/golang_workspace/authentication/models"

	"github.com/golang_workspace/authentication/controllers"
	"github.com/golang_workspace/authentication/mail"
	"github.com/julienschmidt/httprouter"
)

// GtForgot for GET on /forgot
func GtForgot(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	_, err := Authenticate(req)
	if err != nil {
		tpl.ExecuteTemplate(w, "forgot.html", nil)
		return
	}
	http.Redirect(w, req, "/user/admin", http.StatusSeeOther)
}

// PstForgot for POST on /forgot
func PstForgot(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	email := req.FormValue("email")
	alerts := models.UserAlerts{}
	if !(govalidator.IsEmail(email)) {
		alerts.ErrorMessage = "Email does not match required email pattern."
		tpl.ExecuteTemplate(w, "forgot.html", alerts)
		return
	}

	ui, err := controllers.GetUserByEmail(email)
	if err != nil {
		// Sleep to give the feeling to user that we are processing your mail send request
		// so that it does not differ with mail exists or not.
		time.Sleep(2 * time.Second)
		alerts.SuccessMessage = "If your mail exists then we have sent you a password reset link. Please check your mail."
		tpl.ExecuteTemplate(w, "forgot.html", alerts)
		return
	}

	ur, err := controllers.CreateReset(email)
	if err != nil {
		alerts.ErrorMessage = err.Error()
		tpl.ExecuteTemplate(w, "forgot.html", alerts)
		return
	}

	htmlContent, plainContent := getEmailContent(ur.ID.Hex(), strings.Split(config.RedirectURI, "/google/callback")[0])

	err = mailme.SendMail(ui.Fname, ui.Email, "Reset Password", plainContent, htmlContent)
	if err != nil {
		alerts.ErrorMessage = err.Error()
		tpl.ExecuteTemplate(w, "forgot.html", alerts)
		return
	}

	alerts.SuccessMessage = "If your mail exists then we have sent you a password reset link. Please check your mail."
	tpl.ExecuteTemplate(w, "forgot.html", alerts)
	return
}

func getEmailContent(id, url string) (htmlContent, plainContent string) {
	html := "Your password reset link is: <a href=\"" + url + "/user/reset?id=" + id + "\">" + url + "/user/reset?id=" + id + "</a><br><strong>Note: </strong>Link will be valid for only 3 hours"
	plain := "Your password reset link is: " + id + ". Note: This link will be valid for only 3 hours."
	return html, plain
}
