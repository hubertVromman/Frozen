package main

import (
	"net"
	"fmt"
	"bufio"
	// "html"
)

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		// buf := make([]byte, 0, 4096)
		buf, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		// fmt.Println(nbytes);
		fmt.Println("New message :")
		fmt.Print(buf)
	}
}

func main() {
	fmt.Println("Starting server...")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	fmt.Println("Server started")
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		fmt.Println("New connection !")
		go handleConnection(conn)
	}
}