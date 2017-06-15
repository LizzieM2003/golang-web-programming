package main

import (
	"net/http"

	"github.com/LizzieM2003/golang-web-programming/27_mongodb/06_handsOn/03/controllers"
	"github.com/LizzieM2003/golang-web-programming/27_mongodb/06_handsOn/03/sessions"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(sessions.GetSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}
