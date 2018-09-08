package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/golang_workspace/authentication/controllers"
	"github.com/julienschmidt/httprouter"
)

const passRegExp string = "^(((?=.*[a-z])(?=.*[A-Z]))|((?=.*[a-z])(?=.*[0-9]))|((?=.*[A-Z])(?=.*[0-9])))(?=.{6,})"

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
	controllers.Connect()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
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
	fmt.Println("Server started on port", port)
	http.ListenAndServe(":"+port, mux)
}
