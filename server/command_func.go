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
					send_mp((*users)[id], ": You left the channel" + names[i])
					tmp := (*channels)[names[i]][len((*channels)[names[i]]) - 1]
					(*channels)[names[i]][j] = tmp
					(*channels)[names[i]] = (*channels)[names[i]][0:len((*channels)[names[i]]) - 1]
					break
			//for k:= range (*channels)[names[i]]{
				//send message
			//}
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

func	NAMES_cmd(msg string, id int, users *[]User, channels *map[string][]int)(int){
	fmt.Println("NAMES_cmd:")
	fmt.Println("Client: ", msg)

	var boolean bool
	names := strings.Split(msg, ",")
	if names[0] != ""{
		for i:= range names{
			fmt.Println("debug", names[i])
			boolean = true
			if _, ok := (*channels)[names[i]]; ok{
				for value, _:= range (*channels)[names[i]]{
					send_mp((*users)[id], "353 = " + names[i] + " :" + (*users)[value].nickname)
				}
				send_mp((*users)[id], "366 " + (*users)[id].nickname + " " + names[i] + " :End of /NAMES list")
			}else{
				send_mp((*users)[id], "401 " + names[i] + " :No such channel")
			}
		}
	}
	if !boolean{
		for i:= range *channels{
			for value, _:= range (*channels)[i]{
				send_mp((*users)[id], "353 = " + i + " :" + (*users)[value].nickname)
			}
			send_mp((*users)[id], "366 " + (*users)[id].nickname + " " + i + " :End of /NAMES list")
		}
		send_mp((*users)[id], "366 " + (*users)[id].nickname + " * :End of /NAMES list")
	}
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

func	PRIVMSG_cmd(str string, id int, users *[]User, channel map[string][]int)(int){
	fmt.Println("PRIVMSG_cmd:")
	fmt.Println(str)

	whole := strings.Split(str, ":")
	tab := strings.Split(whole[0], " ")
	names := strings.Split(tab[0], ",")
	if (len(whole) > 1){
		for i := range names{
			if names[i][0] == '#'{
				var chan_found bool
				if _, ok := channel[names[i]];ok{
					chan_found = true
					for id_of_chan,_ := range channel[names[i]]{
						send_mp((*users)[id_of_chan], ": " + (*users)[id].nickname + " PRIVMSG " + names[i] + " :" + whole[1])
					}
					if !chan_found{
						send_mp((*users)[id], "401: " + (*users)[id].nickname + " " + names[i] + " :No such channel")
					}
				}else{
					send_mp((*users)[id], "401: " + (*users)[id].nickname + " " + names[i] + " :No such channel")
				}
			}else{
				var someone_found bool
				for j := range *users{
					if (*users)[j].nickname == names[i]{
						someone_found = true
						send_mp((*users)[j], ": " + (*users)[id].nickname + " PRIVMSG " + names[i] + " :" + whole[1])
					}
				}
				if (!someone_found){
					send_mp((*users)[id], "401: " + (*users)[id].nickname + " " + names[i] + " :No such nick")
				}
			}
		}
	}else{
		send_mp((*users)[id], "412 " + (*users)[id].nickname + " :No text to send")
	}
	return (1)
}

func CAP_END_cmd(this_user *User) (){
	fmt.Println("CAP_END_cmd")
	send_mp(*this_user, "001 " + (*this_user).nickname + " :Welcome to the most modern place to chat, " + this_user.nickname)
	send_mp(*this_user,
	"                             ,--.\"\"")
	send_mp(*this_user,
	"                      __,----( o ))")
	send_mp(*this_user,
	"                    ,'--.      , (")
	send_mp(*this_user,
	"             -\"\",:-(    o ),-'/  ;")
	send_mp(*this_user,
	"               ( o) `o  _,'\\ / ;(")
	send_mp(*this_user,
	"                `-;_-<'\\_|-'/ '  )")
	send_mp(*this_user,
	"                    `.`-.__/ '   |")
	send_mp(*this_user,
	"       \\`.            `. .__,   ;")
	send_mp(*this_user,
	"        )_;--.         \\`       |")
	send_mp(*this_user,
	"       /'(__,-:         )      ;")
	send_mp(*this_user,
	"     ;'    (_,-:     _,::     .|")
	send_mp(*this_user,
	"    ;       ( , ) _,':::'    ,;")
	send_mp(*this_user,
	"   ;         )-,;'  `:'     .::")
	send_mp(*this_user,
	"   |         `'  ;         `:::\\")
	send_mp(*this_user,
	"   :       ,'    '            `:\\")
	send_mp(*this_user,
	"   ;:    '  _,-':         .'     `-.")
	send_mp(*this_user,
	"    ';::..,'  ' ,        `   ,__    `.")
	send_mp(*this_user,
	"      `;''   / ;           _;_,-'     `.")
	send_mp(*this_user,
	"            /            _;--.          \\")
	send_mp(*this_user,
	"          ,'            / ,'  `.         \\")
	send_mp(*this_user,
	"         /:            (_(   ,' \\         )")
	send_mp(*this_user,
	"        /:.               \\_(  /-. .:::,;/")
	send_mp(*this_user,
	"       (::..                 `-'\\ \"`\"\"'")
	send_mp(*this_user,
	"       ;::::.                    \\        __")
	send_mp(*this_user,
	"       ,::::::.            .:'    )    ,-'  )")
	send_mp(*this_user,
	"      /  `;:::::::'`__,:.:::'    /`---'   ,'")
	send_mp(*this_user,
	"     ;    `\"\"\"\"'   (  \\:::'     /     _,-'")
	send_mp(*this_user,
	"     ;              \\  \\:'    ,';:.,-'")
	send_mp(*this_user,
	"     (              :  )\\    (")
	send_mp(*this_user,
	"      `.             \\   \\    ;")
	send_mp(*this_user,
	"        `-.___       : ,\\ \\  (")
	send_mp(*this_user,
	"           ,','._::::| \\ \\ \\  \\")
	send_mp(*this_user,
	"          (,(,---;;;;;  \\ \\|;;;)")
	send_mp(*this_user,
	"                      `._\\_\\")
	return
}
