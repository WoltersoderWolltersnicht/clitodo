package views

import (
	"clitodo/cmd"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type ViewID int

const (
	View1Const ViewID = iota
	View2Const
)

type MainView struct {
	currentView ViewID
	view1       tea.Model
	view2       tea.Model
	KeyMap      cmd.KeyMap
}

func NewMainView() tea.Model {
	return MainView{
		0,
		NewListScreen(),
		nil,
		cmd.DefaultKeyMap(),
	}
}

func (m MainView) Init() tea.Cmd {
	return nil
}

func (m MainView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.KeyMap.ForceQuit) {
			return m, tea.Quit
		}
	case cmd.AddTaskTrigger:
		m.view2 = NewAddTaskScreen()
		m.currentView = View2Const
	case cmd.TaskAdded:
		m.currentView = View1Const
	}

	var cmd tea.Cmd

	switch m.currentView {
	case View1Const:
		m.view1, cmd = m.view1.Update(msg)
	case View2Const:
		m.view2, cmd = m.view2.Update(msg)
	}

	return m, cmd
}

// The main view, which just calls the appropriate sub-view
func (m MainView) View() string {
	switch m.currentView {
	case View1Const:
		return m.view1.View()
	case View2Const:
		return m.view2.View()
	default:
		return "Unknown view"
	}
}
