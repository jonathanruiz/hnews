package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Story struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type model struct {
	stories []Story // story from stories
	err     error   // error from stories
	cursor  int     // which story our cursor is pointing at
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

// Init returns the initial command to run.
func (m model) Init() tea.Cmd {
	return nil
}

// Update handles messages.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}

	return m, nil
}

// View renders the model.
func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v", m.err)
	}

	if len(m.stories) == 0 {
		return "No stories available"
	}

	var lines []string
	for _, story := range m.stories {
		lines = append(lines, fmt.Sprintf("- %s (%s)", story.Title, story.URL))
	}

	return strings.Join(lines, "\n")
}

// main starts the program.
func main() {
	stories, err := getTopStories(10)

	p := tea.NewProgram(model{stories: stories, err: err, cursor: 1})

	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v", err)
	}

}
