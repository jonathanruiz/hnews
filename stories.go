package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
)

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
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "r":
			stories, err := getTopStories(5)
			m = model{
				stories: stories,
				err:     err,
				cursor:  1,
			}
			return m, nil
		case "up", "k":
			if m.cursor > 1 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.stories) {
				m.cursor++
			}
		case "enter":
			if m.cursor > 0 && m.cursor <= len(m.stories) {
				return storyView(m.stories[m.cursor-1]), nil
			}
		}
	}

	return m, nil
}

// View renders the model.
func (m model) View() string {
	return allStoriesView(m)
}

func allStoriesView(m model) string {
	lines := "Here are the top stories from Hacker News:\n\n"
	switch {
	case m.err != nil:
		return fmt.Sprintf("Error: %v", m.err)
	case len(m.stories) == 0:
		return "No stories available"
	default:
		for index, story := range m.stories {
			// Add an asterisk to the currently selected story
			if index+1 == m.cursor {
				lines += fmt.Sprintf("> %d. %s (%s)\n", index+1, story.Title, story.URL)
			} else {
				lines += fmt.Sprintf("  %d. %s (%s)\n", index+1, story.Title, story.URL)
			}
		}

		lines += "\nPress 'q' to quit. Press 'r' to refresh.\n"
		lines += fmt.Sprintf("Currently selected: %d\n", m.cursor)

		return lines
	}
}
