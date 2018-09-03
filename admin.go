package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func gtAdmin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ui, err := authenticate(req)
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "admin.html", ui)
}
