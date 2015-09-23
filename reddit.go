package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Feed struct {
	URL   string
	Title string
}

type Payload struct {
	Data struct {
		Children []struct {
			Data struct {
				URL   string `json:"url"`
				Title string `json:"title"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

func GetFeed(term string) ([]Feed, error) {

	feeds := make([]Feed, 0)

	url := fmt.Sprintf("http://reddit.com/r/%s.json", term)

	response, err := http.Get(url)

	var payload Payload

	if err != nil {
		return feeds, err
	}

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return feeds, err
	}

	err = json.Unmarshal(bytes, &payload)

	if err != nil {
		return feeds, err
	}
	for _, v := range payload.Data.Children {
		feeds = append(feeds, Feed{v.Data.URL, v.Data.Title})
	}

	return feeds, nil
}
