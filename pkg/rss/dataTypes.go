package rss

import (
	"encoding/xml"
)

type Client struct {
}

type Item struct {
	Text     string `xml:",chardata"`
	Title    string `xml:"title"`
	Link     string `xml:"link"`
	Comments string `xml:"comments"`
	Creator  string `xml:"creator"`
	PubDate  string `xml:"pubDate"`
	Date     int64
	Category []string `xml:"category"`
	Guid     struct {
		Text        string `xml:",chardata"`
		IsPermaLink string `xml:"isPermaLink,attr"`
	} `xml:"guid"`
	Description string `xml:"description"`
	CommentRss  string `xml:"commentRss"`
}

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Content string   `xml:"content,attr"`
	Wfw     string   `xml:"wfw,attr"`
	Dc      string   `xml:"dc,attr"`
	Atom    string   `xml:"atom,attr"`
	Sy      string   `xml:"sy,attr"`
	Slash   string   `xml:"slash,attr"`
	Channel struct {
		Text  string `xml:",chardata"`
		Title string `xml:"title"`
		Link  struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
		Description     string `xml:"description"`
		LastBuildDate   string `xml:"lastBuildDate"`
		Language        string `xml:"language"`
		UpdatePeriod    string `xml:"updatePeriod"`
		UpdateFrequency string `xml:"updateFrequency"`
		Generator       string `xml:"generator"`
		Image           struct {
			Text   string `xml:",chardata"`
			URL    string `xml:"url"`
			Title  string `xml:"title"`
			Link   string `xml:"link"`
			Width  string `xml:"width"`
			Height string `xml:"height"`
		} `xml:"image"`
		Item []Item `xml:"item"`
	} `xml:"channel"`
}
