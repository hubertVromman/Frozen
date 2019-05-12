package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
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

func send_mp(dest User, message string) bool {
	message += "\r\n"
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

func getData(conn net.Conn, users *[]User, channels *[]Channel) {
	reader := bufio.NewReader(conn)
	var this_user User
	this_user.online = true
	this_user.conn = conn
	var user_id int = -1
	for {
		buf, err := reader.ReadString('\n')
		if err != nil {
			if (user_id == -1){
				fmt.Println("Connexion lost with", "unkown user", conn.RemoteAddr().String())
			}else {
				fmt.Println("Connexion lost with", (*users)[user_id], conn.RemoteAddr().String())
			}
			return
		}
		fmt.Println(buf)
		buf = buf[0:len(buf) - 2]
		if (strings.HasPrefix(buf, "NICK ")){
			str := NICK_cmd(conn, strings.Trim(strings.TrimPrefix(buf, "NICK "), " "), &user_id, users, &this_user)
			if (!(send_mp(this_user, str))){
				if (user_id != -1){
					(*users)[user_id].online = false
				}
				return
			}
		}
		if (strings.HasPrefix(buf, "USER ")){
			USER_cmd(conn, strings.Trim(strings.Split(strings.TrimPrefix(buf, "USER "), " ")[0], " "), &user_id, users, &this_user)
		}
		if (strings.HasPrefix(buf, "PASS ")){
			PASS_cmd(conn, strings.Trim(strings.TrimPrefix(buf, "USER "), " "), &user_id, users, &this_user)
		}
		if (strings.HasPrefix(buf, "JOIN ")){
			JOIN_cmd(conn, strings.Trim(strings.TrimPrefix(buf, "USER "), " "), &user_id, users, &this_user)
		}
		if (strings.HasPrefix(buf, "PART ")){
			PART_cmd(conn, strings.Trim(strings.TrimPrefix(buf, "USER "), " "), &user_id, users, &this_user)
		}
		if (strings.HasPrefix(buf, "NAMES ")){
			NAMES_cmd(conn, strings.Trim(strings.TrimPrefix(buf, "USER "), " "), &user_id, users, &this_user)
		}
		if (strings.HasPrefix(buf, "LIST ")){
			LIST_cmd(conn, strings.Trim(strings.TrimPrefix(buf, "USER "), " "), &user_id, users, &this_user)
		}
		if (strings.HasPrefix(buf, "PRIVMSG ")){
			PRIVMSG_cmd(conn, strings.Trim(strings.TrimPrefix(buf, "USER "), " "), &user_id, users, &this_user)
		}
		if (strings.HasPrefix(buf, "CAP END")){
			CAP_END_cmd(conn, &this_user)
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
	// go getAllMessages(&str)
	fmt.Println("Server started")
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		fmt.Println("New connection !")
		go getData(conn, &users, &channels)
	}
}
