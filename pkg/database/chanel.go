package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func addChanel(channel string) int {
	isTelegram := true
	if strings.Contains(channel, ":") || strings.Contains(channel, "/") {
		isTelegram = false
	}

	var id int
	err := Db.QueryRow("insert into db.public.channels(channel_url, is_telegram,last_updated)values ($1, $2, 0) returning id", channel, isTelegram).Scan(&id)
	if err != nil {
		log.Fatalln("add chenel 53", err)
	}

	return id
}

func AddChannelToUser(userId int, channel string) bool {
	var resalt bool
	fmt.Println(userId, channel)
	channelId := GetChannelId(channel)

	fmt.Println(channelId)

	var id int
	err := Db.QueryRow("select id from db.public.sub where user_id=$1 and channel_id=$2", userId, channelId).Scan(&id)
	fmt.Println(err)

	if err != sql.ErrNoRows {
		fmt.Println("add channels error: ", err)
		return false
	}
	if channelId == -1 {
		resalt = true
		channelId = addChanel(channel)
	}
	fmt.Println(channelId)
	Db.Exec("insert into db.public.sub (user_id, channel_id)values ($1, $2)", userId, channelId)

	return resalt
}

func DeleteChannelToUser(userId int, channel string) {
	id := GetChannelId(channel)
	Db.Exec("delete from db.public.sub where user_id=$1 and channel_id=$2", userId, id)
}

func GetChanelName(id int) string {
	var name string
	Db.QueryRow("select channel_url from db.public.channels where id=$1", id).Scan(&name)
	return name
}

func GetChannelId(channel string) int {
	var id int

	err := Db.QueryRow("select id from db.public.channels where channel_url=$1", channel).Scan(&id)
	if err != nil {
		return -1
	}
	return id
}

func GetChannels(user int64) []string {
	r, err := Db.Query("select channel_url from db.public.channels")
	if user != 0 {
		r, err = Db.Query("select channel_url from db.public.channels inner join db.public.sub s on channels.id = s.channel_id where user_id=$1", user)
	}
	if err != nil {
		log.Fatalln("get chenel", err)
	}

	var s string
	arr := []string{}

	for r.Next() {
		r.Scan(&s)
		arr = append(arr, s)
	}
	return arr

}
