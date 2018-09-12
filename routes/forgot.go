package routes

import (
	"fmt"
	"net/http"

	"github.com/golang_workspace/authentication/controllers"

	"github.com/julienschmidt/httprouter"
)

// GtForgot for GET on /forgot
func GtForgot(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(w, "forgot.html", nil)
}

// PstForgot for POST on /forgot
func PstForgot(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	email := req.FormValue("email")
	ui, err := controllers.GetUserByEmail(email)
	fmt.Println(ui, err)
	tpl.ExecuteTemplate(w, "forgot.html", nil)
}
