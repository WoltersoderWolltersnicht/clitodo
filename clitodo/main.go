package main

import (
	"clitodo/pkg"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var MainListScreen *pkg.ListScreen

func main() {
	MainListScreen = pkg.NewListScreen()
	p := tea.NewProgram(MainListScreen, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
