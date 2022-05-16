package telegram

var client *Client

func Create(token string) *Client {
	client = &Client{token}
	return client
}
