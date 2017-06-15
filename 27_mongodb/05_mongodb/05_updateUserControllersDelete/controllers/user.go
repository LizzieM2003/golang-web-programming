package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LizzieM2003/golang-web-programming/27_mongodb/05_mongodb/05_updateUserControllersDelete/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")
	fmt.Println("id", id)

	// Verify id is ObjectId hex representation
	if !bson.IsObjectIdHex(id) {
		fmt.Println("Not hexadecimal")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// ObjectIdHex returns an ObjectId from the provided hex representation
	oid := bson.ObjectIdHex(id)
	fmt.Println("oid", oid)

	// fetch user from mongodb
	u := models.User{}

	if err := uc.session.DB("go-web-dev-db").C("users").FindId(oid).One(&u); err != nil {
		fmt.Println(err)
		fmt.Println("Didn't find record in mongodb")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// marshal into JSON
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	// send Content-Type, status and payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// test via command:
	// curl -X POST -H "Content-Type: application/json" -d '{"Name":"James Bond","Gender":"male","Age":32}' http://localhost:8080/user
	u := models.User{}

	// decode the json in the request body
	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	// store the user in mongodb
	uc.session.DB("go-web-dev-db").C("users").Insert(u)

	// marshal the new user into JSON and send it in the response
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// test via command:
	// curl -X DELETE -H "Content-Type: application/json" http://localhost:8080/user/<user_id>
	id := p.ByName("id")

	// Check if id is in hexadecimal
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Convert hexadecimal id to Object id
	oid := bson.ObjectIdHex(id)

	// delete user from database
	if err := uc.session.DB("go-web-dev-db").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Deleted user", oid)
}
