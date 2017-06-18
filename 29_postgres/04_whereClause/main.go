package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

type book struct {
	Isbn   string
	Title  string
	Author string
	Price  float64
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
	fmt.Println("Successfully connected to bookstore database!")
}

func main() {
	http.HandleFunc("/books", booksIndex)
	http.HandleFunc("/book", showBook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func booksIndex(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query(`SELECT * FROM books;`)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	defer rows.Close()

	bks := make([]book, 0)
	for rows.Next() {
		bk := book{}
		err = rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)

		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		bks = append(bks, bk)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	jsn, _ := json.Marshal(bks)
	w.Write(jsn)
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

	row := db.QueryRow(`SELECT * FROM books WHERE isbn = $1`, isbn)

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
	jsn, _ := json.Marshal(bk)
	w.Write(jsn)
}
