package main

import (
	"strings"
	// "net"
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
func	NICK_cmd(msg string, id int, users *[]User)(string){
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
		return ("432 " + nickname + " :Erroneus nickname")
	}
	for i := range *users{
		if (i == id && strings.Compare((*users)[i].nickname, msg) == 0){
			return ("")
		}
		if (i != id && strings.Compare((*users)[i].nickname, msg) == 0){
			fmt.Println("error 433")
			return ("433 " + (*users)[id].nickname + " " +  nickname  + " :Nickname is already in use")
		}
	}
	nickname = (*users)[id].nickname
	(*users)[id].nickname = msg
	(*users)[id].nickname = msg
	fmt.Println("Changed name properly")
	return (":" + nickname + " NICK")
}

func	JOIN_cmd(msg string, user_id int, users *[]User, channels *[]Channel) (int) {
	fmt.Println("JOIN_cmd:")
	fmt.Println("Client: ", msg)
	splitted := strings.Split(msg, ",")
	for _, chan_name := range splitted { //loop over param
		if strings.ContainsAny(chan_name, ", \007") || (chan_name[0] != '#' && chan_name[0] != '&') {
			send_mp((*users)[user_id], "479 :" + chan_name + " Illegal channel name")
			return (-1)
		}
		for channel_id, channel := range *channels { //search channel
			if channel.name == chan_name {
				for _, chan_id := range (*users)[user_id].cur_channel { //search if already in channel
					if chan_id == channel_id {
						return (0)
					}
				}
				(*users)[user_id].cur_channel = append((*users)[user_id].cur_channel, channel_id) //join channel
				channel.users_id = append(channel.users_id, user_id)
				send_mp((*users)[user_id], "JOIN " + chan_name)
				send_mp((*users)[user_id], "MODE " + chan_name + " +nt")
				send_mp((*users)[user_id], "353 = " + chan_name + " :@hvromman")
				send_mp((*users)[user_id], "366 " + chan_name + " :End of /NAMES list")
				return (0)
			}
		}
		(*users)[user_id].cur_channel = append((*users)[user_id].cur_channel, len(*channels))
		*channels = append(*channels, Channel{chan_name, []int{user_id}}) //create channel
		send_mp((*users)[user_id], "JOIN " + chan_name)
		send_mp((*users)[user_id], "MODE " + chan_name + " +nt")
		send_mp((*users)[user_id], "353 = " + chan_name + " :@hvromman")
		send_mp((*users)[user_id], "366 " + chan_name + " :End of /NAMES list")
	}
	return (0)
}
func	PART_cmd(msg string, id int, users *[]User)(int){
	fmt.Println("PART_cmd:")
	fmt.Println("Client: ", msg)
	
	return (1)
}
func	NAMES_cmd(msg string, id int, users *[]User)(int){
	fmt.Println("NAMES_cmd:")
	fmt.Println("Client: ", msg)
	return (1)
}
func	LIST_cmd(msg string, user_id int, users *[]User)(int) {
	fmt.Println("LIST_cmd:")
	fmt.Println("Client: ", msg)
	return (0)
}
func	PRIVMSG_cmd(msg string, id int, users *[]User)(int){
	fmt.Println("PRIVMSG_cmd:")
	fmt.Println(msg)
	return (1)
}
func CAP_END_cmd(this_user *User) (){
	fmt.Println("CAP_END_cmd")
	send_mp(*this_user, "001 :Welcome to the most modern place to chat, " + this_user.nickname)
	send_mp(*this_user, "002 :" +
"                             ,--.\"\"")
	send_mp(*this_user, "002 :" +
"                      __,----( o ))")
	send_mp(*this_user, "002 :" +
"                    ,'--.      , (")
	send_mp(*this_user, "002 :" +
"             -\"\",:-(    o ),-'/  ;")
	send_mp(*this_user, "002 :" +
"               ( o) `o  _,'\\ / ;(")
	send_mp(*this_user, "002 :" +
"                `-;_-<'\\_|-'/ '  )")
	send_mp(*this_user, "002 :" +
"                    `.`-.__/ '   |")
	send_mp(*this_user, "002 :" +
"       \\`.            `. .__,   ;")
	send_mp(*this_user, "002 :" +
"        )_;--.         \\`       |")
	send_mp(*this_user, "002 :" +
"       /'(__,-:         )      ;")
	send_mp(*this_user, "002 :" +
"     ;'    (_,-:     _,::     .|")
	send_mp(*this_user, "002 :" +
"    ;       ( , ) _,':::'    ,;")
	send_mp(*this_user, "002 :" +
"   ;         )-,;'  `:'     .::")
	send_mp(*this_user, "002 :" +
"   |         `'  ;         `:::\\")
	send_mp(*this_user, "002 :" +
"   :       ,'    '            `:\\")
	send_mp(*this_user, "002 :" +
"   ;:    '  _,-':         .'     `-.")
	send_mp(*this_user, "002 :" +
"    ';::..,'  ' ,        `   ,__    `.")
	send_mp(*this_user, "002 :" +
"      `;''   / ;           _;_,-'     `.")
	send_mp(*this_user, "002 :" +
"            /            _;--.          \\")
	send_mp(*this_user, "002 :" +
"          ,'            / ,'  `.         \\")
	send_mp(*this_user, "002 :" +
"         /:            (_(   ,' \\         )")
	send_mp(*this_user, "002 :" +
"        /:.               \\_(  /-. .:::,;/")
	send_mp(*this_user, "002 :" +
"       (::..                 `-'\\ \"`\"\"'")
	send_mp(*this_user, "002 :" +
"       ;::::.                    \\        __")
	send_mp(*this_user, "002 :" +
"       ,::::::.            .:'    )    ,-'  )")
	send_mp(*this_user, "002 :" +
"      /  `;:::::::'`__,:.:::'    /`---'   ,'")
	send_mp(*this_user, "002 :" +
"     ;    `\"\"\"\"'   (  \\:::'     /     _,-'")
	send_mp(*this_user, "002 :" +
"     ;              \\  \\:'    ,';:.,-'")
	send_mp(*this_user, "002 :" +
"     (              :  )\\    (")
	send_mp(*this_user, "002 :" +
"      `.             \\   \\    ;")
	send_mp(*this_user, "002 :" +
"        `-.___       : ,\\ \\  (")
	send_mp(*this_user, "002 :" +
"           ,','._::::| \\ \\ \\  \\")
	send_mp(*this_user, "002 :" +
"          (,(,---;;;;;  \\ \\|;;;)")
	send_mp(*this_user, "002 :" +
"                      `._\\_\\")
	return
}
