package main

import (
	"net"
	"fmt"
	"bufio"
	// "time"
//	"strings"
)

type User struct {
	password string
	nickname string
	username string
	cur_channel []int
	online bool
	ip string
	conn net.Conn
}

type Message struct {
	data string
	sender_id int
	dest string
}

type Channel struct {
	name string
	users_id []int
}

func send_mp(dest User, message string, error int) bool {
	if dest.online {
		dest.conn.Write([]byte(message))
		return true
	} else {
		return false
	}
}

func sendData(message Message, users []User, channels []Channel) {
	if message.dest[0] == '#' { //channel message
		message.dest = message.dest[1:]
		for _, channel := range channels {
			if message.dest == channel.name { //channel found
				for _, user_id := range channel.users_id { //all users of channel
					send_mp(users[user_id], message.data)
				}
				return
			}
		}
		send_mp(users[message.sender_id], "Channel not found")
	} else { //private message
		for _, user := range users {
			if message.dest == user.nickname { //user found
				if !send_mp(user, message.data) {
					send_mp(users[message.sender_id], "User not connected")
				}
				return
			}
		}
		send_mp(users[message.sender_id], "User not found")
	}
}

func getData(conn net.Conn, message *Message, users *[]User, user_id *int, channels *[]Channel) {
	reader := bufio.NewReader(conn)
	for {
		buf, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(conn == nil)
			return
		}
		fmt.Println(buf)
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
	}
}
