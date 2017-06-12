package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "test01_user:password@/test_01")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/read", read)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "at index")
	check(err)
}

func amigos(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT aName from amigos;`)
	check(err)
	defer rows.Close()

	var name, s string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(50))`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "CREATED TABLE customer", n)
}

func insert(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO customer VALUES ("Lizzie");`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	fmt.Fprintln(w, "Inserted Record:", n)
}

func update(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`UPDATE customer SET name = "Elizabeth";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)
	fmt.Fprintln(w, "UPDATED RECORD:", n)
}

func delete(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM customer WHERE name="Elizabeth";`)
	check(err)

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)
	fmt.Fprintln(w, "DELETED RECORD:", n)
}

func read(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT * FROM customer;`)
	check(err)
	defer rows.Close()

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Fprintln(w, "Retrieved record:", name)
	}
}

func drop(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE customer;`)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(w, "CUSTOMER TABLE DROPPED")
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
