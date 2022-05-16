package telegram

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type command struct {
	name        string
	description string
}

var commands = [...]command{{"list", "list all subs"}, {"start", "start bot"}, {"channels", "add chennel"}}

func SetMyCommands() error {
	//5298008548:AAEhOUFeM7WmX5lR6QduTdQpjbrTXQL-ME8
	//1253744008
	cl := http.Client{Timeout: 10 * time.Second}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/setMyCommands", client.token)

	var commandsWithDescription []string
	for _, v := range commands {
		commandsWithDescription = append(commandsWithDescription, fmt.Sprintf(`{
					"command": %q,
					"description": %q
				}`, v.name, v.description))
	}

	payload := fmt.Sprintf(`
	{
		"commands": ['%s']
	}
`, strings.Join(commandsWithDescription, "','"))

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
