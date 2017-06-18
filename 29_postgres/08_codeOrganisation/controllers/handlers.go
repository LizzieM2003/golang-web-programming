package controllers

import (
	"database/sql"
	"net/http"

	"github.com/LizzieM2003/golang-web-programming/29_postgres/08_codeOrganisation/config"
	"github.com/LizzieM2003/golang-web-programming/29_postgres/08_codeOrganisation/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bks, err := models.AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.Tpl.ExecuteTemplate(w, "books.gohtml", bks)
}

func Show(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := models.OneBook(r)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}

	config.Tpl.ExecuteTemplate(w, "show.gohtml", bk)
}

func Create(w http.ResponseWriter, r *http.Request) {
	config.Tpl.ExecuteTemplate(w, "create.gohtml", nil)
}

func CreateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := models.PutBook(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}
	config.Tpl.ExecuteTemplate(w, "created.gohtml", bk)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := models.OneBook(r)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.Tpl.ExecuteTemplate(w, "update.gohtml", bk)
}

func UpdateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	bk, err := models.UpdateBook(r)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	config.Tpl.ExecuteTemplate(w, "updated.gohtml", bk)
}

func DeleteProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	err := models.DeleteBook(r)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
