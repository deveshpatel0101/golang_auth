package main

import (
	"html/template"
	"net/http"

	"github.com/golang_workspace/authentication/controllers"
	"github.com/julienschmidt/httprouter"
)

const passRegExp string = "^(((?=.*[a-z])(?=.*[A-Z]))|((?=.*[a-z])(?=.*[0-9]))|((?=.*[A-Z])(?=.*[0-9])))(?=.{6,})"

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
	controllers.Connect()
	go controllers.RemoveSessions()
}

func main() {
	mux := httprouter.New()
	mux.GET("/", gtIndex)
	mux.GET("/login", gtLogin)
	mux.POST("/login", pstLogin)
	mux.GET("/signup", gtSignup)
	mux.POST("/signup", pstSignup)
	mux.GET("/logout", gtLogout)
	mux.GET("/admin", gtAdmin)
	mux.GET("/google/login", google)
	mux.GET("/google/callback", callback)
	http.ListenAndServe(":8000", mux)
}
