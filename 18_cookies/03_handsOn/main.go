package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var counter int

func main() {
	http.HandleFunc("/", set)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func set(w http.ResponseWriter, req *http.Request) {
	counter++
	c, err := req.Cookie("visit-count")
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "visit-count",
			Value: strconv.Itoa(counter),
		}
		http.SetCookie(w, c)
	} else if err != nil {
		log.Println(err)
	} else {
		c.Value = strconv.Itoa(counter)
		http.SetCookie(w, c)
	}
	fmt.Fprintln(w, "You have visited this site", c.Value, "time(s)")
}
