package main

import (
	"log"
	"net/http"

	"github.com/LizzieM2003/golang-web-programming/30_mongodb/01_goMongo/controllers"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/books", controllers.Index)
	http.HandleFunc("/books/show", controllers.Show)
	http.HandleFunc("/books/create", controllers.Create)
	http.HandleFunc("/books/create/process", controllers.CreateProcess)
	http.HandleFunc("/books/update", controllers.Update)
	http.HandleFunc("/books/update/process", controllers.UpdateProcess)
	http.HandleFunc("/books/delete", controllers.DeleteProcess)
	http.HandleFunc("/books/json", controllers.ShowJson)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
