// In addition to running this file, you also need to run main.go in 01_write
// in order to have a tcp server writing output which this program can read

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}
