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

var db *sql.DB
var tpl *template.Template

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://bond:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You are now connected to the bookstore database!")

	tpl = template.Must(template.ParseGlob("templates/*"))
}

type book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/books", booksIndex)
	http.HandleFunc("/books/show", showBook)
	http.HandleFunc("/books/create", createBook)
	http.HandleFunc("/books/create/process", createBookProcess)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	rows, err := db.Query(`SELECT * FROM books;`)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	bks := make([]book, 0)

	for rows.Next() {
		bk := book{}
		if err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price); err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		}
		bks = append(bks, bk)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "books.gohtml", bks)
}

func showBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	isbn := r.FormValue("isbn")
	if isbn == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := db.QueryRow(`SELECT * from books WHERE isbn = $1;`, isbn)
	bk := book{}

	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
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
		http.Error(w, http.StatusText(406)+"Please hit back and enter a number for the price", http.StatusNotAcceptable)
		return
	}
	bk.Price = float32(f64)

	// Insert values into db

	_, err = db.Exec(`INSERT INTO BOOKS(isbn, title, author, price) VALUES ($1, $2, $3, $4);`, bk.Isbn, bk.Title, bk.Author, bk.Price)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "created.gohtml", bk)
}
