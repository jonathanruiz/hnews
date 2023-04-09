package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Comment struct {
	Id       int    `json:"id"`
	By       string `json:"by"`
	Parent   int    `json:"parent"`
	Kids     []int  `json:"kids"`
	Text     string `json:"text"`
	Time     int    `json:"time"`
	Type     string `json:"type"`
	Dead     bool   `json:"dead"`
	Deleted  bool   `json:"deleted"`
	DeadKids []int  `json:"deadKids"`
}

func getComments(id int) Comment {
	// Via the Hacker News API, retrieve the comment for the given ID.
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)

	// Retrieve the comment through GET request
	resp, err := http.Get(url)

	// Return any error if the GET request fails.
	if err != nil {
		return Comment{}
	}
	defer resp.Body.Close()

	// Decode the JSON response into a Comment.
	var comment Comment
	err = json.NewDecoder(resp.Body).Decode(&comment)
	if err != nil {
		return Comment{}
	}

	// Return the comment.
	return comment
}
