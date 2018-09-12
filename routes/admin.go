package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GtAdmin will listen to GET on admin route
func GtAdmin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ui, err := Authenticate(req)
	if err != nil {
		http.Redirect(w, req, "/user/login", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "admin.html", ui)
}
