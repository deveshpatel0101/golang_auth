package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func gtIndex(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}
