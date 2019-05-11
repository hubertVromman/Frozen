package main

import (
	"net"
	"fmt"
	"bufio"
	"time"
	"strings"
)

type User struct {
	username string
	password string
	nickname string
	cur_channel int
}

func sendData(conn net.Conn, message *string , users []User, user_id *int) {
	for {
		if *message != "" {
			fmt.Println("New message :")
			fmt.Print(*message)
			conn.Write([]byte(*message))
			<-time.After(time.Millisecond)
			*message = ""
		}
	}
}

func getData(conn net.Conn, message *string, users *[]User, user_id *int) {
	reader := bufio.NewReader(conn)
	for {
		// buf := make([]byte, 0, 4096)
		buf, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		if strings.HasPrefix(buf, "PASS NICK USER") {
			fmt.Printf("full %v\n", buf)
			buf = strings.TrimSpace(strings.Replace(buf, "PASS NICK USER", "", 1))	
			fmt.Printf("after replace %v\n", buf)
			separated := strings.Split(buf, " ")
			fmt.Println(separated, len(separated))
			*user_id = -1
			for n := range *users {
				if (*users)[n].username == separated[0] {
					if (*users)[n].password == separated[1] {
						*user_id = n
					} else {
						*message = "auth failed\n"
					}
				}
			}
			if *user_id == -1 {
				*users = append(*users,  User{separated[0], separated[1], separated[0], -1})
				*user_id = len(*users) - 1;
			}
			fmt.Println((*users)[*user_id], *user_id)
		} else {
			newbuf := buf
			*message = newbuf
		}
	}
}

func main() {
	fmt.Println("Starting server...")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	var users []User
	var users_id []int
	var message string
	// go getAllMessages(&str)
	fmt.Println("Server started")
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		fmt.Println("New connection !")
		users_id = append(users_id, -1)
		go getData(conn, &message, &users, &users_id[len(users_id) - 1])
		go sendData(conn, &message, users, &users_id[len(users_id) - 1])
	}
}