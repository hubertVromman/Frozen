package main

import (
	"strings"
	"net"
	"fmt"
)

const (
//	NICK
	ERR_NONICKNAMEGIVEN = 431
	ERR_ERRONEUSNICKNAME = 432
	ERR_NICKNAMEINUSE = 433
	ERR_NICKCOLLISION = 436
	MAXNICKLEN = 15
//	USER
	R_NEEDMOREPARAMS = 461
	ERR_ALREADYREGISTRED = 462
)
func	NICK_cmd(conn net.Conn, msg string, user_id *int, users *[]User)(string){
	fmt.Println("NICK_cmd:")
	fmt.Println(msg)
	var nickname string
	nickname = msg
	if (len(msg) == 0){
		return("431 " + " :No nickname received")
	}
	if (len(msg) > MAXNICKLEN || strings.Contains(msg, " ")){
		return ("432 " + nickname + " :Erroneus nickname")
	}
	for i := range *users{
		if (strings.Compare((*users)[i].nickname, msg) == 0){
			return ("433 " + nickname + " :Nickname is already in use")
		}
	}
	if (*user_id == -1){
		return (":" + nickname + " NICK " + msg)
	}
	nickname = (*users)[*user_id].nickname
	(*users)[*user_id].nickname = msg
	return (":" + nickname + " NICK " + msg)
}

func	USER_cmd(conn net.Conn, msg string, user_id *int, users *[]User)(int){
	fmt.Println("USER_cmd:")
	fmt.Println(msg)
	return (1)
}

func	PASS_cmd(conn net.Conn, msg string, user_id *int, users *[]User)(int){
	fmt.Println("PASS_cmd:")
	fmt.Println(msg)
	return (1)
}
func	JOIN_cmd(conn net.Conn, msg string, user_id *int, users *[]User)(int){
	fmt.Println("JOIN_cmd:")
	fmt.Println(msg)
	return (1)
}
func	PART_cmd(conn net.Conn, msg string, user_id *int, users *[]User)(int){
	fmt.Println("PART_cmd:")
	fmt.Println(msg)
	return (1)
}
func	NAMES_cmd(conn net.Conn, msg string, user_id *int, users *[]User)(int){
	fmt.Println("NAMES_cmd:")
	fmt.Println(msg)
	return (1)
}
func	LIST_cmd(conn net.Conn, msg string, user_id *int, users *[]User)(int){
	fmt.Println("LIST_cmd:")
	fmt.Println(msg)
	return (1)
}
func	PRIVMSG_cmd(conn net.Conn, msg string, user_id *int, users *[]User)(int){
	fmt.Println("PRIVMSG_cmd:")
	fmt.Println(msg)
	return (1)
}
