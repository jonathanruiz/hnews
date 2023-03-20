package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Story struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func getTopStories() ([]Story, error) {
	// Via the Hacker News API, retrieve a list of the top story IDs.
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode the JSON response into a slice of integers.
	var ids []int
	err = json.NewDecoder(resp.Body).Decode(&ids)
	if err != nil {
		return nil, err
	}

	// For the top 5 stories, retrieve the story and append it to the slice of stories.
	stories := make([]Story, 0, 5)
	for _, id := range ids[:5] {
		url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var story Story
		err = json.NewDecoder(resp.Body).Decode(&story)
		if err != nil {
			return nil, err
		}

		stories = append(stories, story)
	}

	return stories, nil
}

func main() {
	stories, err := getTopStories()

	if err != nil {
		panic(err)
	}

	for _, story := range stories {
		println(story.Title)
		println(story.URL)
	}

}
