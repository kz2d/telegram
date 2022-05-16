package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var Db *sql.DB

func Connect(host string, port int, dbname, user, password string) error {
	connStr := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		host, port, dbname, user, password)
	fmt.Println(connStr)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	Db = conn

	Db.SetMaxOpenConns(10)

	err = Db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func AddUser(userId int, chatId int64) {
	Db.Exec("insert into db.public.users (user_id, chat_id, show, state)values ($1, $2, true,'')", userId, chatId)
}

func ChangeState(userId int, state string) {
	Db.Exec("update db.public.users set state=$1 where user_id=$2", state, userId)
}

func ChangeShow(userId int, show bool) {
	Db.Exec("update db.public.users set show=$1 where user_id=$2", show, userId)
}

func GetSubsToChanel(channel string) []int64 {
	var id int64
	id = int64(int(GetChannelId(channel)))
	r, err := Db.Query("select chat_id from db.public.sub inner join db.public.users c on c.user_id = sub.user_id where channel_id=$1 and show=true", id)
	if err != nil {
		log.Fatalln("get subs to chenel", err)
	}

	arr := []int64{}

	for r.Next() {
		r.Scan(&id)
		arr = append(arr, id)
	}
	return arr
}

func GetLastUpdate(channel string) int64 {
	var id int64
	Db.QueryRow("select last_updated from db.public.channels where channel_url=$1", channel).Scan(&id)
	return id
}

func UpdateLastUpdate(channel string) {
	Db.Exec("update db.public.channels set last_updated=$1 where channel_url=$2", time.Now().Unix(), channel)
}

func GetState(userId int) string {
	rows, err := Db.Query("SELECT state FROM db.public.users where user_id=$1", userId)
	if err != nil {
		fmt.Println(err)
	}

	var name string

	defer rows.Close()
	if rows.Next() {

		err = rows.Scan(&name)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(name)
	}
	return name
}

func SetLastActiviti(date int64) {
	Db.Exec("update db.public.lastdate set date=$1 where name='update'", date)
}

func GetLastActiviti() int64 {
	var id int64
	err := Db.QueryRow("select date from db.public.lastdate where name='update'").Scan(&id)
	if err != nil {
		return 0
	}
	fmt.Println(id)
	return id
}
