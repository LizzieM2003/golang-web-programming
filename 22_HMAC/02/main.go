package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/authenticate", auth)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		c = &http.Cookie{
			Name:  "session",
			Value: "",
			Path:  "/",
		}
		http.SetCookie(w, c)
	}

	if req.Method == http.MethodPost {
		e := req.FormValue("email")
		c.Value = e + "|" + getCode(e)
		http.SetCookie(w, c)
	}

	io.WriteString(w, `
	<!DOCTYPE html>
	<html>
	<body>
		<form method="post">
			<input type="email" name="email">
			<input type="submit">
		</form>
		<a href="/authenticate">Validate This `+c.Value+`</a>
	</body>
	</html>
	`)
}

func auth(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if c.Value == "" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	xs := strings.Split(c.Value, "|")
	email := xs[0]
	codeRcvd := xs[1]
	codeCheck := getCode(email)

	fmt.Println(codeRcvd)
	fmt.Println(codeCheck)

	if codeRcvd != codeCheck {
		fmt.Println("HMAC codes didn't match!")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	io.WriteString(w, `
	<!DOCTYPE html>
	<html>
	<body>
		<h1>`+codeRcvd+` - RECEIVED </h1>
		<h1>`+codeCheck+` - RECALCULATED </h1>
	</body>
	</html>
	`)
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))

}
