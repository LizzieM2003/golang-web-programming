package config

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

var DB *mgo.Database

var Books *mgo.Collection

func init() {
	s, err := mgo.Dial("mongodb://bond:moneypenny007@localhost/bookstore")
	// s, err := mgo.Dial("mongodb://localhost/bookstore")

	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	DB = s.DB("bookstore")
	Books = DB.C("books")

	fmt.Println("You are now connected to the bookstore database!")
}
