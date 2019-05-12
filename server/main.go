package main

import (
	"net"
	"fmt"
	"bufio"
	"time"
	"strings"
)

type User struct {
	password string
	nickname string
	username string
	cur_channel []int
}

type Message struct {
	data string
	sender_id int
	dest string
}

type Channel struct {
	name string
	users []User
	is_open int
}

func sendData(conn net.Conn, message *Message, users *[]User, user_id *int, channels *[]Channel) {
	for {
		if *user_id != -1 {
			if message.channel_id == (*users)[*user_id].cur_channel {
				conn.Write([]byte(message.nick + ": " + message.data))
			}
			<-time.After(time.Millisecond)
			message.channel_id = -1
		}
	}
}

func getData(conn net.Conn, message *Message, users *[]User, user_id *int, channels *[]Channel) {
	reader := bufio.NewReader(conn)
	buf, err := reader.ReadString('\n')
	if err != nil {
		return
	}
}

func main() {
	fmt.Println("Starting server...")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	var users []User
	var channels []Channel
	var users_id []int
	var message Message
	// go getAllMessages(&str)
	fmt.Println("Server started")
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		fmt.Println("New connection !")
		users_id = append(users_id, -1)
		go getData(conn, &message, &users, &users_id[len(users_id) - 1], &channels)
		go sendData(conn, &message, &users, &users_id[len(users_id) - 1], &channels)
	}
}