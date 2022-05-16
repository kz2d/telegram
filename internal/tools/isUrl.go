package tools

import (
	"regexp"
	"strings"
)

func IsUrl(channel string) bool {
	ok := regexp.MustCompile(`^https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)$`).MatchString(channel)
	if !ok {
		return channel[0] == '@' && !strings.Contains(channel, ":") && !strings.Contains(channel, "/")
	}
	return true
}
