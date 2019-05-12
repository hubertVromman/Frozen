package main

import (
	"strings"
	"strconv"
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

func	JOIN_cmd(msg string, user_id int, users *[]User, channels *map[string][]int) (int) {
	fmt.Println("JOIN_cmd:")
	fmt.Println("Client: ", msg)
	splitted := strings.Split(msg, ",")
	for _, chan_name := range splitted { //loop over param
		if strings.ContainsAny(chan_name, ", \007") || (chan_name[0] != '#' && chan_name[0] != '&') {
			send_mp((*users)[user_id], "479 :" + chan_name + " Illegal channel name")
			return (-1)
		}
		if _, ok := (*users)[user_id].cur_channel[chan_name]; ok {
			continue
		}
		if _, ok := (*channels)[chan_name]; ok { //search channel
			(*channels)[chan_name] = append((*channels)[chan_name], user_id) //join channel
		} else {
			(*channels)[chan_name] = []int{user_id} //create channel
		}
		(*users)[user_id].cur_channel[chan_name] = true
		send_mp((*users)[user_id], "JOIN " + chan_name)
		send_mp((*users)[user_id], "MODE " + chan_name + " +nt")
		send_mp((*users)[user_id], "353 = " + chan_name + " :@hvromman")
		send_mp((*users)[user_id], "366 " + chan_name + " :End of /NAMES list")
	}
	return (0)
}

func	PART_cmd(msg []string, id int, users *[]User, channels *map[string][]int)(int){
	fmt.Println("PART_cmd:")
	fmt.Println("Client: ", msg)


	fmt.Println("msg received: ", msg)

	names := strings.Split(msg[0], ",")
	for i := range names{
		fmt.Println("name: ", names)
		if _, ok := (*channels)[names[i]]; ok{
			var boolean bool = false
			for j:= range (*channels)[names[i]]{
				if ((*channels)[names[i]][j] == id){
					boolean = true
			//for k:= range (*channels)[names[i]]{
				//send message
			//}
					tmp := (*channels)[names[i]][len((*channels)[names[i]]) - 1]
					(*channels)[names[i]][j] = tmp
					(*channels)[names[i]] = (*channels)[names[i]][0:len((*channels)[names[i]]) - 1]
				}
			}
			if _, ok2 := (*users)[id].cur_channel[names[i]]; ok2{
				delete((*users)[id].cur_channel, names[i])
			}
			if !boolean{
			send_mp((*users)[id], "442 " + (*users)[id].nickname + " " + names[i] + " :You're not on that channel")
		}
		}else{
			send_mp((*users)[id], "403 " + (*users)[id].nickname + " " + names[i] + " :No such channel")
		}
	}
	return (1)
}

func	NAMES_cmd(msg string, id int, users *[]User)(int){
	fmt.Println("NAMES_cmd:")
	fmt.Println("Client: ", msg)
	return (1)
}

func	LIST_cmd(msg []string, user_id int, users *[]User, channels *map[string][]int)(int) {
	fmt.Println("LIST_cmd:")
	fmt.Println("Client: ", msg)
	send_mp((*users)[user_id], "321 :Channel Users Name")
	if msg[0] != "" {
		splitted := strings.Split(msg[0], ",")
		for _, chan_name := range splitted {
			if users_id, ok := (*channels)[chan_name]; ok {
				send_mp((*users)[user_id], "322 :" + chan_name + " " + strconv.Itoa(len(users_id)) + " :[+nt]")
			}
		}
	} else {
		for chan_name, users_id := range *channels {
			send_mp((*users)[user_id], "322 :" + chan_name + " " + strconv.Itoa(len(users_id)) + " :[+nt]")
		}
	}
	send_mp((*users)[user_id], "323 " + (*users)[user_id].nickname + " :End of /LIST")
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
