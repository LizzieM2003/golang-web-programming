package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LizzieM2003/golang-web-programming/27_mongodb/03_postAndeDelete/models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	r.POST("/user", createUser)
	r.DELETE("/user/:id", deleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="utf-8">
			<title></title>
		</head>
			<a href="/user/12345678">GO TO localhost:8080/user/12345678</a>
		</html>
	`
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// use the following command to simulate creating a user:
	// curl -X POST -H "Content-Type: application/json" -d '{"Name":"James Bond","Gender":"male","Age":32,"Id":"777"}' http://localhost:8080/user
	u := models.User{}

	// use decode to grab JSON in request
	json.NewDecoder(r.Body).Decode(&u)

	// change id
	u.Id = "007"

	// assign json to a variable via marshal
	uj, _ := json.Marshal(u)

	// write Content-Type, status and payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func deleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: write code to delete user
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Write code to delete user")
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James",
		Gender: "male",
		Age:    32,
		Id:     p.ByName("id"),
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}
