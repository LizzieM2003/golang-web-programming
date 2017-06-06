// In addition to this program, also need to run 02_read that reads from a connection to the server
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	fmt.Fprintln(conn, "I dialled you!")
}
