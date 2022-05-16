package main

import (
	"flag"
	"github.com/kz2d/telegram-bot/config"
	"github.com/kz2d/telegram-bot/internal/grpc"
	"github.com/kz2d/telegram-bot/internal/rssCallback"
	"github.com/kz2d/telegram-bot/internal/telegramCallback"
	"github.com/kz2d/telegram-bot/pkg/database"
	"github.com/kz2d/telegram-bot/pkg/rss"
	"github.com/kz2d/telegram-bot/pkg/telegram"
	"log"
)

var configPath = flag.String("config", "config/build.yaml", "config path")

func main() {
	flag.Parse()
	parsedFile, err := config.Load(*configPath)
	if err != nil {
		log.Fatalln(err)
	}
	err = database.Connect(parsedFile.Database.Host, parsedFile.Database.Port, parsedFile.Database.Dbname, parsedFile.Database.User, parsedFile.Database.Password)
	if err != nil {
		log.Fatalln("database connection error: ", err)
	}
	defer database.Db.Close()

	telegram.Create(parsedFile.Telegram.Token)

	go grpc.Serve("8080")
	rss.Create()
	for _, v := range database.GetChannels(0) {
		go rss.Start(v, rssCallback.Callback)
	}

	telegram.Start(telegramCallback.StateMachine)
}
