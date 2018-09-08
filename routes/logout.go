package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GtLogout will listen to GET on logout route
func GtLogout(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ck := http.Cookie{
		Name:     "sid",
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, &ck)
	http.Redirect(w, req, "/", http.StatusSeeOther)
	return
}
