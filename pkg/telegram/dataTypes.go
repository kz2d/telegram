package telegram

type Client struct {
	token string
}

type User struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type Chat struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
	Title     string `json:"title"`
}

type MessageEntity struct {
	Type   string `json:"type"`
	Offset int64  `json:"offset"`
	Length int64  `json:"length"`
	Url    string `json:"url"`
	User   string `json:"user"`
}

type Message struct {
	MessageID int64           `json:"message_id"`
	From      User            `json:"from"`
	Chat      Chat            `json:"chat"`
	Date      int             `json:"date"`
	Text      string          `json:"text"`
	Entities  []MessageEntity `json:"entities"`
}

type AutoGenerated struct {
	Ok     bool `json:"ok"`
	Result []struct {
		UpdateID     int64   `json:"update_id"`
		Message      Message `json:"message,omitempty"`
		MyChatMember struct {
			Chat          Chat `json:"chat"`
			From          User `json:"from"`
			Date          int  `json:"date"`
			OldChatMember struct {
				User   User   `json:"user"`
				Status string `json:"status"`
			} `json:"old_chat_member"`
			NewChatMember struct {
				User   User   `json:"user"`
				Status string `json:"status"`
			} `json:"new_chat_member"`
		} `json:"my_chat_member,omitempty"`
	} `json:"result"`
}
