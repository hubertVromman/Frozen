package main

import (
	"net"
	"os"
	"fmt"
	"bufio"
)

func getMessage(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		buf, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		fmt.Print(buf)
	}
}

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
	go getMessage(conn)
	for {
		buf, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		conn.Write([]byte(buf))
	}
}