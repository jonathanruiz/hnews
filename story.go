package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Story struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func (s Story) Init() tea.Cmd {
	return nil
}

func (s Story) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return s, tea.Quit
		}
	}

	return s, nil
}

func (s Story) View() string {
	lines := "Here is the story you selected:\n\n"
	lines += fmt.Sprintf("%s (%s)\n", s.Title, s.URL)
	lines += "\nPress 'q' to quit. Press 'r' to refresh."

	return lines
}

func storyView(story Story) Story {
	return story
}
