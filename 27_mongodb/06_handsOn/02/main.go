package main

import (
	"net/http"

	"github.com/LizzieM2003/golang-web-programming/27_mongodb/06_handsOn/02/controllers"
	"github.com/LizzieM2003/golang-web-programming/27_mongodb/06_handsOn/02/models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() map[string]models.User {
	// get data from file and store in map
	m := models.LoadUsers()
	return m
}
