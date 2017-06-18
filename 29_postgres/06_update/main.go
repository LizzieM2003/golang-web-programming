package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

var tpl *template.Template
var db *sql.DB

type book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://bond:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You are connected to the bookstore database!")

	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/books", booksIndex)
	http.HandleFunc("/books/show", bookShow)
	http.HandleFunc("/books/create", createBook)
	http.HandleFunc("/books/create/process", createBookProcess)
	http.HandleFunc("/books/update", updateBook)
	http.HandleFunc("/books/update/process", updateBookProcess)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/books", http.StatusSeeOther)
	return
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query(`SELECT * FROM books;`)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	bks := make([]book, 0)
	for rows.Next() {
		bk := book{}
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		bks = append(bks, bk)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w, "books.gohtml", bks)
}

func bookShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	isbn := r.FormValue("isbn")

	if isbn == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := db.QueryRow(`SELECT * FROM books WHERE isbn = $1;`, isbn)

	bk := book{}
	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)

	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), 500)
		return

	}

	tpl.ExecuteTemplate(w, "show.gohtml", bk)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "create.gohtml", nil)
}

func createBookProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk := book{}

	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		http.Error(w, http.StatusText(406)+" Please hit back to enter a numeric price", http.StatusNotAcceptable)
		return
	}
	bk.Price = float32(f64)

	_, err = db.Exec(`INSERT INTO books (isbn, title, author, price) VALUES ($1, $2, $3, $4);`, bk.Isbn, bk.Title, bk.Author, bk.Price)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	tpl.ExecuteTemplate(w, "created.gohtml", bk)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	isbn := r.FormValue("isbn")

	if isbn == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	bk := book{}

	row := db.QueryRow(`SELECT * FROM books WHERE isbn = $1;`, isbn)

	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)

	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "update.gohtml", bk)
}

func updateBookProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk := book{}

	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		http.Error(w, http.StatusText(405), http.StatusBadRequest)
		return
	}

	f64, err := strconv.ParseFloat(p, 32)

	if err != nil {
		http.Error(w, http.StatusText(406)+"Press back to enter numeric price value", http.StatusNotAcceptable)
		return
	}

	bk.Price = float32(f64)

	_, err = db.Exec(`UPDATE books SET title=$2, author=$3, price=$4 WHERE isbn=$1;`, bk.Isbn, bk.Title, bk.Author, bk.Price)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	tpl.ExecuteTemplate(w, "updated.gohtml", bk)
}
