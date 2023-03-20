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

func getTopStories(numStories int) ([]Story, error) {
	// Via the Hacker News API, retrieve a list of the top story IDs.
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")

	// Return any error if the GET request fails.
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

	// For the top stories, retrieve the story and append it to the slice of stories.
	// Number of stories is based off of numStories.
	stories := make([]Story, 0, numStories)
	for _, id := range ids[:numStories] {

		// Via the Hacker News API, retrieve the story for the given ID.
		url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)

		// Retrieve the story through GET request
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		// Decode the JSON response into a Story.
		var story Story
		err = json.NewDecoder(resp.Body).Decode(&story)
		if err != nil {
			return nil, err
		}

		// Append the story to the slice.
		stories = append(stories, story)
	}

	// Return the slice of stories.
	return stories, nil
}

func main() {
	// Retrieve the top stories.
	stories, err := getTopStories(3)

	// Panic if an error is returned.
	if err != nil {
		panic(err)
	}

	// Print the title and URL for each story.
	for _, story := range stories {
		println(story.Title)
		println(story.URL)
	}

}
