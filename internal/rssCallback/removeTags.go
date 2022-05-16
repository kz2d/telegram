package rssCallback

import (
	"regexp"
	"strings"
)

func remove(text string) string {
	if len(text) > 9 && text[:9] == "<![CDATA[" {
		text = text[9:]
	}

	if len(text) > 3 && text[len(text)-3:] == "]]>" {
		text = text[:len(text)-3]
	}
	return text
}

func Preprocess(text string) (string, string) {
	usedTags := [...]string{"b", "i", "strong", "em", "u", "ins", "s", "strike", "del", "a", "code", "pre"}
	image := ""

	text = remove(text)

	text = regexp.MustCompile("<h.+?>").ReplaceAllString(text, "<b>")
	text = regexp.MustCompile("<br.*?/>").ReplaceAllString(text, "\n")
	text = regexp.MustCompile("</h.>").ReplaceAllString(text, "</b>")

	pos := regexp.MustCompile(`(?Um)<img.*src\s?=\s?"(\S+)".*/>`).FindAllStringSubmatch(text, -1)

	if len(pos) > 0 {
		image = pos[0][1]
	}

	tags := regexp.MustCompile(`</?(\w+).*?/?>`).FindAllStringSubmatch(text, -1)

out:
	for _, v := range tags {
		for _, w := range usedTags {
			if v[1] == w {
				continue out
			}
		}
		text = strings.Replace(text, v[0], "", 1)
	}

	return strings.Trim(text, " \t\n"), image
}
