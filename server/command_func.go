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
func	NICK_cmd(conn net.Conn, msg string, user_id *int, users *[]User, this_user *User)(string){
	fmt.Println("NICK_cmd:")
	fmt.Println("Client: ", msg)
	var nickname string
	nickname = msg
	if (len(msg) == 0){
		fmt.Println("error 431")
		return("431 " + " :No nickname received")
	}
	if (len(msg) > MAXNICKLEN || strings.Contains(msg, " ")){
		fmt.Println("error 432")
		return ("432 client " + nickname + " :Erroneus nickname")
	}
	for i := range *users{
		if (strings.Compare((*users)[i].nickname, msg) == 0){
			fmt.Println("error 433")
			return ("433 client " + nickname + " :Nickname is already in use")
		}
	}
	if (*user_id == -1){
		*user_id = len(*users)
		(*this_user).nickname = nickname
		*users = append(*users, *this_user)
		fmt.Println("Added a new user with ID :" , *user_id)
		return (":" + nickname + " NICK " + msg)
	}
	nickname = (*users)[*user_id].nickname
	(*users)[*user_id].nickname = msg
	(*this_user).nickname = msg
	fmt.Println("Changed name properly")
	return (":" + nickname + " NICK")
}

func	USER_cmd(conn net.Conn, msg string, user_id *int, users *[]User, this_user *User)(int){
	fmt.Println("USER_cmd:")
	fmt.Println("Client: ", msg)
	send_mp(*this_user, "CAP *")
	return (1)
}

func	PASS_cmd(conn net.Conn, msg string, user_id *int, users *[]User, this_user *User)(int){
	fmt.Println("PASS_cmd:")
	fmt.Println("Client: ", msg)
	return (1)
}
func	JOIN_cmd(conn net.Conn, msg string, user_id *int, users *[]User, this_user *User)(int){
	fmt.Println("JOIN_cmd:")
	fmt.Println("Client: ", msg)
	return (1)
}
func	PART_cmd(conn net.Conn, msg string, user_id *int, users *[]User, this_user *User)(int){
	fmt.Println("PART_cmd:")
	fmt.Println("Client: ", msg)
	return (1)
}
func	NAMES_cmd(conn net.Conn, msg string, user_id *int, users *[]User, this_user *User)(int){
	fmt.Println("NAMES_cmd:")
	fmt.Println("Client: ", msg)
	return (1)
}
func	LIST_cmd(conn net.Conn, msg string, user_id *int, users *[]User, this_user *User)(int){
	fmt.Println("LIST_cmd:")
	fmt.Println("Client: ", msg)
	return (1)
}
func	PRIVMSG_cmd(conn net.Conn, msg string, user_id *int, users *[]User, this_user *User)(int){
	fmt.Println("PRIVMSG_cmd:")
	fmt.Println(msg)
	return (1)
}
func CAP_END_cmd(conn net.Conn, this_user *User) (){
	fmt.Println("CAP_END_cmd")
	send_mp(*this_user, "001: Welcome to the most modern place to chat, " + this_user.nickname)	
	send_mp(*this_user, "002: " +
"                             ,--.\"\"")
	send_mp(*this_user, "002: " +
"                      __,----( o ))")
	send_mp(*this_user, "002: " +
"                    ,'--.      , (")
	send_mp(*this_user, "002: " +
"             -\"\",:-(    o ),-'/  ;")
	send_mp(*this_user, "002: " +
"               ( o) `o  _,'\\ / ;(")
	send_mp(*this_user, "002: " +
"                `-;_-<'\\_|-'/ '  )")
	send_mp(*this_user, "002: " +
"                    `.`-.__/ '   |")
	send_mp(*this_user, "002: " +
"       \\`.            `. .__,   ;")
	send_mp(*this_user, "002: " +
"        )_;--.         \\`       |")
	send_mp(*this_user, "002: " +
"       /'(__,-:         )      ;")
	send_mp(*this_user, "002: " +
"     ;'    (_,-:     _,::     .|")
	send_mp(*this_user, "002: " +
"    ;       ( , ) _,':::'    ,;")
	send_mp(*this_user, "002: " +
"   ;         )-,;'  `:'     .::")
	send_mp(*this_user, "002: " +
"   |         `'  ;         `:::\\")
	send_mp(*this_user, "002: " +
"   :       ,'    '            `:\\")
	send_mp(*this_user, "002: " +
"   ;:    '  _,-':         .'     `-.")
	send_mp(*this_user, "002: " +
"    ';::..,'  ' ,        `   ,__    `.")
	send_mp(*this_user, "002: " +
"      `;''   / ;           _;_,-'     `.")
	send_mp(*this_user, "002: " +
"            /            _;--.          \\")
	send_mp(*this_user, "002: " +
"          ,'            / ,'  `.         \\")
	send_mp(*this_user, "002: " +
"         /:            (_(   ,' \\         )")
	send_mp(*this_user, "002: " +
"        /:.               \\_(  /-. .:::,;/")
	send_mp(*this_user, "002: " +
"       (::..                 `-'\\ \"`\"\"'")
	send_mp(*this_user, "002: " +
"       ;::::.                    \\        __")
	send_mp(*this_user, "002: " +
"       ,::::::.            .:'    )    ,-'  )")
	send_mp(*this_user, "002: " +
"      /  `;:::::::'`__,:.:::'    /`---'   ,'")
	send_mp(*this_user, "002: " +
"     ;    `\"\"\"\"'   (  \\:::'     /     _,-'")
	send_mp(*this_user, "002: " +
"     ;              \\  \\:'    ,';:.,-'")
	send_mp(*this_user, "002: " +
"     (              :  )\\    (")
	send_mp(*this_user, "002: " +
"      `.             \\   \\    ;")
	send_mp(*this_user, "002: " +
"        `-.___       : ,\\ \\  (")
	send_mp(*this_user, "002: " +
"           ,','._::::| \\ \\ \\  \\")
	send_mp(*this_user, "002: " +
"          (,(,---;;;;;  \\ \\|;;;)")
	send_mp(*this_user, "002: " +
"                      `._\\_\\")
}
