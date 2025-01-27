package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var MainListScreen *ListScreen

func main() {
	MainListScreen = NewListScreen()
	p := tea.NewProgram(MainListScreen, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
