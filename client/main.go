package main

import (
	"net"
	"os"
	"fmt"
	"bufio"
	// "strings"
	// "io"
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
	conn, err := net.Dial("tcp", "192.168.1.14:6667")
	
	if err != nil {
		fmt.Println("connection failed")
		return
	}
	defer conn.Close()
	// fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	// status, err := bufio.NewReader(conn).ReadString('\n')
	// if err != nil {
	// 	// handle error
	// }
	// fmt.Println(status);
	reader := bufio.NewReader(os.Stdin)
	go getMessage(conn)
	for {
		buf, _ := reader.ReadString('\n')
		// for i := 0; i < len(buf); i++ {
			// fmt.Print("\b")
		// }
		conn.Write([]byte(buf))
	}
}