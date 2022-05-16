package rss

import (
	"encoding/xml"
	"fmt"
	"github.com/kz2d/telegram-bot/internal/tools"
	"github.com/kz2d/telegram-bot/pkg/database"
	"io"
	"log"
	"net/http"
	"time"
)

const rfc2822 = "Mon, 02 Jan 2006 15:04:05 -0700"

func Start(url string, callback func(Item, string)) {
	//url := fmt.Sprintf("https://www.linuxjournal.com/node/feed")
	cl := http.Client{Timeout: 10 * time.Second}

	var lastTime int64 = database.GetLastUpdate(url)
	fmt.Println("last", lastTime)
	for {
		res, _ := cl.Get(url)

		body, _ := io.ReadAll(res.Body)

		update := Rss{}
		err := xml.Unmarshal(body, &update)
		if err != nil {
			log.Fatalln(err, string(body))
			return
		}

		if update.Version == "2.0" {
			var viewedDate int64
			for _, u := range update.Channel.Item {
				t, err := time.Parse(rfc2822, u.PubDate)
				u.Date = t.Unix()
				if err != nil {
					log.Println("parse warning: ", err)
				}
				if lastTime < u.Date {
					viewedDate = tools.Max(viewedDate, u.Date)

					go callback(u, url)
				}
			}
			lastTime = tools.Max(lastTime, viewedDate)
		}

		database.UpdateLastUpdate(url)
		time.Sleep(1 * time.Minute)

	}

}
