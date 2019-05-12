package main

import (
	"net"
	"fmt"
	"bufio"
	"time"
	"os"
	"strings"
)

type User struct {
	password string
	nickname string
	username string
	cur_channel []int
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

func getData(conn net.Conn, users *[]User, user_id *int, channels *[]Channel) {
		reader := bufio.NewReader(conn)
	for {
		buf, err := reader.ReadString('\n')
		if err != nil {
			if (*user_id == -1){
				fmt.Println("Connexion lost with", "unkown user", conn.RemoteAddr().String())
			}else {
				fmt.Println("Connexion lost with", (*users)[*user_id], conn.RemoteAddr().String())
			}
			return
		}
		fmt.Println(buf)
		if (strings.HasPrefix(buf, "NICK ")){
			NICK_cmd(conn, strings.Trim(strings.TrimPrefix(buf, "NICK "), " "), user_id, users)
		}
		if (strings.HasPrefix(buf, "USER ")){
			USER_cmd(conn, strings.Trim(strings.TrimPrefix(buf, "USER "), " "), user_id, users)
		}
	}
}

func main() {
	fmt.Println("Starting server...")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(-1)
	}
	var users []User
	var channels []Channel
	var users_id []int
	// go getAllMessages(&str)
	fmt.Println("Server started")
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		fmt.Println("New connection !")
		users_id = append(users_id, -1)
		go getData(conn, &users, &users_id[len(users_id) - 1], &channels)
	}
}
