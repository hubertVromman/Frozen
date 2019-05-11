package main

import (
	"net"
	"os"
	// "fmt"
	"bufio"
	// "io/ioutil"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	defer conn.Close()
	if err != nil {
		// handle error
	}
	// fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	// status, err := bufio.NewReader(conn).ReadString('\n')
	// if err != nil {
	// 	// handle error
	// }
	// fmt.Println(status);
	for {
		buf, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		conn.Write([]byte(buf))
	}
}