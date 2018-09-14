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
	mux.GET("/user/login", routes.GtLogin)
	mux.POST("/user/login", routes.PstLogin)
	mux.GET("/user/signup", routes.GtSignup)
	mux.POST("/user/signup", routes.PstSignup)
	mux.GET("/user/logout", routes.GtLogout)
	mux.GET("/user/admin", routes.GtAdmin)
	mux.GET("/google/login", routes.Google)
	mux.GET("/google/callback", routes.Callback)
	mux.GET("/user/forgot", routes.GtForgot)
	mux.POST("/user/forgot", routes.PstForgot)
	mux.GET("/user/reset", routes.GtReset)
	mux.POST("/user/reset", routes.PstReset)
	mux.GET("/user/settings", routes.GtChangePwd)
	mux.POST("/user/settings", routes.PstChangePwd)
	fmt.Println("Server started on port", port)
	http.ListenAndServe(":"+port, mux)
}
