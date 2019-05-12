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
	online	bool
}

type Message struct {
	data string
	sender_id int
	dest string
}

type Channel struct {
	name string
	user_id []int
}

func sendData(conn net.Conn, message *Message, users *[]User, user_id *int, channels *[]Channel) {
	for {
		if *user_id != -1 {
//			if message.channel_id == (*users)[*user_id].cur_channel {
//				conn.Write([]byte(message.nick + ": " + message.data))
//			}
			<-time.After(time.Millisecond)
//			message.channel_id = -1
		}
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
	ln, err := net.Listen("tcp", ":6667")
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
