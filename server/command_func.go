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
//	USER
	R_NEEDMOREPARAMS = 461
	ERR_ALREADYREGISTRED = 462
)
func	NICK_cmd(conn net.Conn, msg string, user_id *int, users *[]User)(int){
	fmt.Println("NICK_cmd:")
	fmt.Println(msg)
	for i := range *users{
		if (strings.Compare((*users)[i].nickname, msg) == 0){
			return (ERR_NICKNAMEINUSE)
		}
	}
	return (1)
}

func	USER_cmd(conn net.Conn, msg string, user_id *int, users *[]User)(int){
	fmt.Println("USER_cmd:")
	fmt.Println(msg)
	return (1)
}
