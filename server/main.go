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
	cur_channel map[string]bool
	online bool
	conn net.Conn
}

type Message struct {
	data string
	sender_id int
	dest string
}

func send_mp(dest User, message string) bool {
	message += "\r\n"
	if dest.online {
		dest.conn.Write([]byte(message))
		fmt.Println("msg sent: ", message)
		return true
	} else {
		return false
	}
}

// func send_to_conn(conn net.Conn, message string) bool {
// 	conn.Write([]byte(message))
// }

// func sendData(message Message, users []User, channels map[string]int) {
// 	if message.dest[0] == '#' { //channel message
// 		message.dest = message.dest[1:]
// 		for _, channel := range channels {
// 			if message.dest == channel.name { //channel found
// 				for _, user_id := range channel.users_id { //all users of channel
// 					send_mp(users[user_id], message.data)
// 				}
// 				return
// 			}
// 		}
// 		send_mp(users[message.sender_id], "Channel not found")
// 	} else { //private message
// 		for _, user := range users {
// 			if message.dest == user.nickname { //user found
// 				if !send_mp(user, message.data) {
// 					send_mp(users[message.sender_id], "User not connected")
// 				}
// 				return
// 			}
// 		}
// 		send_mp(users[message.sender_id], string(401) + " " + message.dest +":No suck nick/channel")
// 	}
// }

func getData(users *[]User, channels *map[string][]int, id int) {
	reader := bufio.NewReader((*users)[id].conn)
	fmt.Println("user: ", id, " in getData loop")
	defer (*users)[id].conn.Close()
	for {
		fmt.Println("user: ", id, " in getData loop")
		buf, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connexion lost with", (*users)[id], (*users)[id].conn.RemoteAddr().String())
			(*users)[id].conn.Close()
			return
		}
		fmt.Println("New message from client: " , buf)
		if (len(buf) > 1){
			buf = buf[0:len(buf) - 2]
		}
		if (strings.HasPrefix(buf, "NICK ")){
			str := NICK_cmd(strings.Trim(strings.TrimPrefix(buf, "NICK "), " "), id, users)
			if (strings.Compare(str, "") == 0){
				continue
			}
			if (!(send_mp((*users)[id], str))){
				(*users)[id].online = false
				return
			}
		}
		if (strings.HasPrefix(buf, "USER ")) {
			send_mp((*users)[id], "462 :You may not reregister")
		}
		if (strings.HasPrefix(buf, "PASS ")) {
			send_mp((*users)[id], "462 :You may not reregister")
		}
		if (strings.HasPrefix(buf, "JOIN ")) {
			JOIN_cmd(strings.Trim(strings.Split(strings.TrimPrefix(buf, "JOIN "), " ")[0], " "), id, users, channels)
		}
		if (strings.HasPrefix(buf, "PART ")) {
			PART_cmd(strings.Split(strings.Trim(strings.TrimPrefix(buf, "PART "), " "), " "), id, users, channels)
		}
		// if (strings.HasPrefix(buf, "WHO ")){
		// 	buf = strings.Replace(buf, "PART ", "LIST ", 1)

		// }
		if (strings.HasPrefix(buf, "NAMES")) {
			NAMES_cmd(strings.Split(strings.Trim(strings.TrimPrefix(buf, "NAMES"), " "), " ")[0], id, users, channels)
		}
		if (strings.HasPrefix(buf, "LIST")) {
			LIST_cmd(strings.Split(strings.Trim(strings.TrimPrefix(buf, "LIST"), " "), " "), id, users, channels)
		}
		if (strings.HasPrefix(buf, "PRIVMSG ")) {
			PRIVMSG_cmd(strings.Trim(strings.TrimPrefix(buf, "PRIVMSG "), " "), id, users)
		}
		if (strings.HasPrefix(buf, "QUIT")){
			(*users)[id].online = false
			return
		}
	}
}

func	identification(users *[]User, this_user *User) (mod, id int) {
	for i := range *users{
		if (strings.Compare((*this_user).username, (*users)[i].username) == 0){
			if (strings.Compare((*this_user).password, (*users)[i].password) == 0){
				mod = 1
				id = i
				(*users)[id].online = true
				if (strings.Compare((*this_user).nickname, (*users)[i].nickname) == 0){
					send_mp((*users)[id], "CAP *")
					return
				}else{
					mod = 2
					send_mp(*this_user, "433 * " + this_user.nickname + " :Nickname is already in use")
					return
				}
			}else{
				send_mp(*this_user, "464 :Password incorrect")
				return
			}
		}
	}
	for i:= range *users{
		if (strings.Compare((*this_user).nickname, (*users)[i].nickname) == 0){
			mod = 2
			send_mp(*this_user, "433 * " + this_user.nickname + " :Nickname is already in use")
			return
		}
	}
	id = len(*users)
	*users = append(*users, *this_user)
	fmt.Println("Added a new user with ID :" , id)
	send_mp((*users)[id], "CAP *")
	mod = 1
	return
}

func tmp_getData(conn net.Conn, users *[]User, channels *map[string][]int) {
	reader := bufio.NewReader(conn)
	var this_user User
	this_user.online = true
	this_user.conn = conn
	this_user.username = ""
	this_user.cur_channel = make(map[string]bool)
	var id int
	var yolo bool
	for {
		buf, err:= reader.ReadString('\n')
		if (len(buf) > 1){
			buf = buf[0:len(buf) - 2]
		}
		fmt.Println(buf)
		if err != nil {
			fmt.Println("Connexion lost with", "unkown user", conn.RemoteAddr().String())
			return
		}
		if (strings.HasPrefix(buf, "NICK ")){
			this_user.nickname = strings.Split(strings.Trim(strings.TrimPrefix(buf, "NICK "), " "), " ")[0]
			if (len(this_user.nickname) > MAXNICKLEN){
				send_mp(this_user, "432 " + this_user.nickname + " :Erroneus nickname")
			}
			if (len(this_user.nickname) == 0){
				send_mp(this_user, "431 " + " :No nickname received")
			}
			fmt.Println("New co nick: " , this_user.nickname)
			if (strings.Compare(this_user.username, "") != 0){
				var mod int
				mod, id = identification(users, &this_user)
				fmt.Println("done with identif from nick", mod)
				if mod == 0{
					return
				}
				if mod == 1{
					yolo = true
				}
			}
		}
		if (strings.HasPrefix(buf, "PASS ")){
			this_user.password = strings.Trim(strings.TrimPrefix(buf, "PASS "), " ")
			fmt.Println("New co pass: " , this_user.password)
		}
		if (strings.HasPrefix(buf, "USER ")){
			this_user.username = strings.Split(strings.Trim(strings.TrimPrefix(buf, "USER "), " "), " ")[0]
			if (strings.Compare(this_user.username, "") == 0){
				send_mp(this_user, "461 USER :Not enough parameters")
			}
			var mod int
			mod, id = identification(users, &this_user)
			fmt.Println("New co user: " , this_user.username, "mod:", mod)
			if mod == 0{
				return
			}
			if mod == 1{
				yolo = true
			}
		}
		if (strings.HasPrefix(buf, "CAP END") && yolo){
			CAP_END_cmd(&this_user)
			break
		}
	}
	getData(users, channels, id)
}

func main() {
	fmt.Println("Starting server...")
	ln, err := net.Listen("tcp", ":6667")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(-1)
	}
	var users []User
	channels := make(map[string][]int)
	// go getAllMessages(&str)
	fmt.Println("Server started")
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		fmt.Println("New connection !")
		go tmp_getData(conn, &users, &channels)
	}
}
