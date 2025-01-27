package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type addTaskScreen struct {
	textInput textinput.Model
}

func AddTaskScreen() addTaskScreen {
	ti := textinput.New()
	ti.Placeholder = "TaskName"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return addTaskScreen{
		textInput: ti,
	}
}

func (m addTaskScreen) Init() tea.Cmd {
	return textinput.Blink
}

func (m addTaskScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			return MainListScreen.Update(m)
		}
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m addTaskScreen) View() string {
	return fmt.Sprintf(
		"Task Title\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
