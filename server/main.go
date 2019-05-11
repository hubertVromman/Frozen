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
	cur_channel int
}

type Message struct {
	data string
	channel_id int
}

type Channel struct {
	name string
	users []User
	is_open int
}

func sendData(conn net.Conn, message *Message, users *[]User, user_id *int, channels *[]Channel) {
	for {
		if *user_id != -1 && message.channel_id != -1 {
			if message.channel_id == (*users)[*user_id].cur_channel {
				conn.Write([]byte(message.data))
			}
			<-time.After(time.Millisecond)
			message.channel_id = -1
		}
	}
}

func getData(conn net.Conn, message *Message, users *[]User, user_id *int, channels *[]Channel) {
	reader := bufio.NewReader(conn)
	for {
		// buf := make([]byte, 0, 4096)
		buf, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		if strings.HasPrefix(buf, "PASS NICK USER") {
			buf = strings.TrimSpace(strings.Replace(buf, "PASS NICK USER", "", 1))
			separated := strings.Split(buf, " ")
			*user_id = -1
			for n := range *users {
				if (*users)[n].username == separated[2] {
					if (*users)[n].password == separated[0] {
						*user_id = n
					} else {
						return //*message = "auth failed\n"
					}
				}
			}
			if *user_id == -1 {
				*users = append(*users, User{separated[0], separated[1], separated[2], -1})
				*user_id = len(*users) - 1;
			}
			fmt.Println((*users)[*user_id], *user_id)
		} else if *user_id != -1 && strings.HasPrefix(buf, "JOIN") {
			buf = strings.TrimSpace(strings.Replace(buf, "JOIN", "", 1))
			(*users)[*user_id].cur_channel = -1
			for n := range *channels {
				if (*channels)[n].name == buf {
					(*users)[*user_id].cur_channel = n
					(*channels)[n].users = append((*channels)[n].users, (*users)[*user_id]);
				}
			}
			if (*users)[*user_id].cur_channel == -1 {
				*channels = append(*channels, Channel{buf, []User{(*users)[*user_id]}, 1})
				(*users)[*user_id].cur_channel = len(*channels) - 1
			}
			fmt.Println((*channels)[(*users)[*user_id].cur_channel])
		} else if *user_id != -1 && (*users)[*user_id].cur_channel != -1 {
			*message = Message{buf, (*users)[*user_id].cur_channel}
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