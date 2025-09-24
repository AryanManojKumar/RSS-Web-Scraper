package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type Rssfeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Items       []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func Urltofeed(url string) (Rssfeed, error) {
	httpclient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpclient.Get(url)
	if err != nil {
		return Rssfeed{}, err
	}
	defer resp.Body.Close()

	bye, err := io.ReadAll(resp.Body)
	if err != nil {
		return Rssfeed{}, err
	}

	rssfeed := Rssfeed{}
	err = xml.Unmarshal(bye, &rssfeed)
	if err != nil {
		return Rssfeed{}, err
	}

	return rssfeed, nil

}
