package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/wordwrap"
)

type Story struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	URL      string `json:"url"`
	Score    int    `json:"score"`
	Comments []int  `json:"kids"`
}

func (s Story) Init() tea.Cmd {
	return nil
}

func (s Story) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			stories, err := getTopStories(DEFAULT_NUM_STORIES)
			return model{stories: stories, err: err, cursor: 1}, nil
		}
	}

	return s, nil
}

func (s Story) View() string {
	lines := fmt.Sprintf("Here is the story you selected: ID# %d (https://news.ycombinator.com/item?id=%d)\n\n", s.Id, s.Id)
	lines += fmt.Sprintf("%s (%s)\n", s.Title, s.URL)
	lines += fmt.Sprintf("Score: %d\n", s.Score)
	lines += fmt.Sprintf("Comments: %d\n", len(s.Comments))

	for _, comment := range s.Comments {
		commentText := getComments(comment).Text
		wrappedText := wordwrap.String(commentText, 80)
		lines += fmt.Sprintf("  %s\n\n", wrappedText)
	}

	lines += "\nPress 'q' to quit. Press 'r' to refresh."

	return lines
}

func storyView(story Story) Story {
	return story
}
