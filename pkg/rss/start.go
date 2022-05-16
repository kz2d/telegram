package rss

var client *Client

func Create() *Client {
	client = &Client{}
	return client
}
