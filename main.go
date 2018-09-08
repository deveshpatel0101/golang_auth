package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang_workspace/authentication/controllers"
	"github.com/golang_workspace/authentication/routes"
	"github.com/julienschmidt/httprouter"
)

const passRegExp string = "^(((?=.*[a-z])(?=.*[A-Z]))|((?=.*[a-z])(?=.*[0-9]))|((?=.*[A-Z])(?=.*[0-9])))(?=.{6,})"

func init() {
	controllers.Connect()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	mux := httprouter.New()
	mux.GET("/", routes.GtIndex)
	mux.GET("/login", routes.GtLogin)
	mux.POST("/login", routes.PstLogin)
	mux.GET("/signup", routes.GtSignup)
	mux.POST("/signup", routes.PstSignup)
	mux.GET("/logout", routes.GtLogout)
	mux.GET("/admin", routes.GtAdmin)
	mux.GET("/google/login", routes.Google)
	mux.GET("/google/callback", routes.Callback)
	fmt.Println("Server started on port", port)
	http.ListenAndServe(":"+port, mux)
}
