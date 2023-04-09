package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

var DEFAULT_NUM_STORIES = 10

// main starts the program.
func main() {
	stories, err := getTopStories(DEFAULT_NUM_STORIES)

	// Loads the model and starts the program in full-screen mode.
	p := tea.NewProgram(model{stories: stories, err: err, cursor: 1}, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
	}

}
