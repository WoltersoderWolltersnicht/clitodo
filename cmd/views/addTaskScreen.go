package views

import (
	"fmt"

	"clitodo/cmd"
	"clitodo/pkg/domain"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type addTaskScreen struct {
	textInput textinput.Model
	KeyMap    cmd.KeyMap
}

func NewAddTaskScreen() addTaskScreen {
	ti := textinput.New()
	ti.Placeholder = "TaskName"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return addTaskScreen{
		textInput: ti,
		KeyMap:    cmd.DefaultKeyMap(),
	}
}

func (m addTaskScreen) Init() tea.Cmd {
	return textinput.Blink
}

func (m addTaskScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.KeyMap.AddTask) { //"enter"
			return m, enterTask(m)
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

func enterTask(m addTaskScreen) tea.Cmd {
	return func() tea.Msg {
		item := domain.NewItem(m.textInput.Value())
		return cmd.TaskAdded{IsSucces: true, Item: item}
	}
}
