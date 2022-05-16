package telegram

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

func SendMessage(chatId int64, text string) error {
	//5298008548:AAEhOUFeM7WmX5lR6QduTdQpjbrTXQL-ME8
	//1253744008
	cl := http.Client{Timeout: 10 * time.Second}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", client.token)
	payload := fmt.Sprintf(`
	{
		"text": %q,
		"chat_id": %d,
		"parse_mode": %q
	}
`, text, chatId, "HTML")

	res, err := cl.Post(url, "application/json", bytes.NewBuffer([]byte(payload)))

	if err != nil {
		return errors.New("bad request")
	}

	if res.StatusCode/100 == 4 {
		body, _ := io.ReadAll(res.Body)
		fmt.Println(payload)
		fmt.Println("status code = %d, body = %s", res.StatusCode, string(body))
	}

	return nil
}

func SendPhoto(chatId int64, text, photoUrl string) error {
	//5298008548:AAEhOUFeM7WmX5lR6QduTdQpjbrTXQL-ME8
	//1253744008
	cl := http.Client{Timeout: 10 * time.Second}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", client.token)
	payload := fmt.Sprintf(`
	{
		"caption": %q,
		"photo": %q,
		"chat_id": %d,
		"parse_mode": %q
	}
`, text, photoUrl, chatId, "HTML")

	res, err := cl.Post(url, "application/json", bytes.NewBuffer([]byte(payload)))

	if err != nil {
		return errors.New("bad request")
	}

	if res.StatusCode/100 == 4 {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("status code = %d, body = %s", res.StatusCode, string(body))
	}

	return nil

}
