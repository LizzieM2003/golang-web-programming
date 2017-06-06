package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	var i int
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}

		if ln == "" {
			//headers are done
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	// request line
	method := strings.Fields(ln)[0]
	url := strings.Fields(ln)[1]
	fmt.Println("***METHOD", method)
	fmt.Println("***URL", url)

	// multiplexer
	if method == "GET" && url == "/" {
		index(conn)
	}

	if method == "GET" && url == "/about" {
		about(conn)
	}

	if method == "GET" && url == "/contact" {
		contact(conn)
	}

	if method == "GET" && url == "/apply" {
		apply(conn)
	}

	if method == "POST" && url == "/apply" {
		applyProcess(conn)
	}
}

func index(conn net.Conn) {
	body := `<DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Index</title>
	</head>
	<body>
	<h1>Index Page</h1>
	<a href="/about">About</a><br>
	<a href="/contact">Contact</a><br>
	<a href="/apply">Apply</a><br>
	</body>
	</html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html \r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
	body := `<DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>About</title>
	</head>
	<body>
	<h1>About Page</h1>
	<a href="/">Index</a><br>
	<a href="/contact">Contact</a><br>
	<a href="/apply">Apply</a><br>
	</body>
	</html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html \r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func contact(conn net.Conn) {
	body := `<DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Contact</title>
	</head>
	<body>
	<h1>Contact Page</h1>
	<a href="/">Index</a><br>
	<a href="/about">About</a><br>
	<a href="/apply">Apply</a><br>
	</body>
	</html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html \r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func apply(conn net.Conn) {
	body := `<DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Apply</title>
	</head>
	<body>
	<h1>Apply Page</h1>
	<a href="/">Index</a><br>
	<a href="/about">About</a><br>
	<a href="/contact">Contact</a><br>
	<a href="/apply">Apply</a><br>
	<form method="POST" action="/apply">
	<input type="submit" value="apply">
	</form>
	</body>
	</html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html \r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func applyProcess(conn net.Conn) {
	body := `<DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Apply Process</title>
	</head>
	<body>
	<h1>Apply Process</h1>
	<a href="/">Index</a><br>
	<a href="/about">About</a><br>
	<a href="/contact">Contact</a><br>
	<a href="/apply">Apply</a><br>
	</body>
	</html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html \r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
