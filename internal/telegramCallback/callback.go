package telegramCallback

import (
	"fmt"
	"github.com/kz2d/telegram-bot/internal/rssCallback"
	"github.com/kz2d/telegram-bot/internal/tools"
	"github.com/kz2d/telegram-bot/pkg/database"
	"github.com/kz2d/telegram-bot/pkg/rss"
	"github.com/kz2d/telegram-bot/pkg/telegram"
	"strings"
)

func StateMachine(message telegram.Message) {
	fmt.Println("message:", message.Text)
	v := database.GetState(message.From.ID)
	switch {
	case v == "add link":
		{
			if tools.IsUrl(message.Text) {
				telegram.SendMessage(message.Chat.ID, "Сылка добавлена")
				if database.AddChannelToUser(message.From.ID, message.Text) {
					go rss.Start(message.Text, rssCallback.Callback)
				}
				fmt.Println(message.Text)
			} else {
				telegram.SendMessage(message.Chat.ID, "wrong url")
			}
			database.ChangeState(message.From.ID, "")
		}

	case message.Text == "/chanels":
		{
			telegram.SendMessage(message.Chat.ID, "введите ссылку канала который нужно добавить")
			database.ChangeState(message.From.ID, "add link")
		}

	case message.Text == "/list":
		{
			s := database.GetChannels(int64(message.From.ID))
			telegram.SendMessage(message.Chat.ID, strings.Join(s, "\n"))
		}

	case message.Text == "/stop":
		{
			telegram.SendMessage(message.Chat.ID, "больше присылаться подбока не будет")
			database.ChangeShow(message.From.ID, false)
		}

	case message.Text == "/start":
		{
			telegram.SendMessage(message.Chat.ID, "Приветсвую в мое rss боте. Чтобы добавить каналы введи /chanels или /help или /list")
			database.AddUser(message.From.ID, message.Chat.ID)
			database.ChangeShow(message.From.ID, true)
		}

	case message.Text == "/help":
		telegram.SendMessage(message.Chat.ID, "Приветсвую в мое rss боте. Чтобы добавить каналы введи /chanels или /help, /stop")
	}

}
