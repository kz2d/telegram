package rssCallback

import (
	"github.com/kz2d/telegram-bot/pkg/database"
	"github.com/kz2d/telegram-bot/pkg/rss"
	"github.com/kz2d/telegram-bot/pkg/telegram"
	"log"
)

func Callback(item rss.Item, url string) {
	text, img := Preprocess(item.Description)
	title, _ := Preprocess(item.Title)

	text = "<b>" + title + "</b>\n" + text

	//fmt.Println(database.GetSubsToChanel(url))

	for _, v := range database.GetSubsToChanel(url) {
		if img != "" && len(text) < 1024 {
			err := telegram.SendPhoto(v, text, img)
			if err != nil {
				log.Fatalln("send photo rss", err)
			}
		} else {
			err := telegram.SendMessage(v, text)
			if err != nil {
				log.Fatalln("send message rss", err)
			}
		}
	}
}
