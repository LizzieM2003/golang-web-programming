package main

import (
	"net/http"

	"github.com/LizzieM2003/golang-web-programming/27_mongodb/05_mongodb/03_updateUserControllersPost/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	// connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")
	// check if connection err
	if err != nil {
		panic(err)
	}
	return s
}
