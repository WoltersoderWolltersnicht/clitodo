package views

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"

	"clitodo/cmd"
	"clitodo/pkg/domain"
)

// DefaultItemStyles defines styling for a default list item.
// See DefaultItemView for when these come into play.
type DefaultItemStyles struct {
	// The Normal state.
	NormalTitle lipgloss.Style

	// The selected item state.
	SelectedTitle lipgloss.Style

	// The dimmed state, for when the filter input is initially activated.
	DimmedTitle lipgloss.Style

	// Characters matching the current filter, if any.
	FilterMatch lipgloss.Style

	CheckMark lipgloss.Style

	EmptyCheckMark lipgloss.Style
}

// NewDefaultItemStyles returns style definitions for a default item. See
// DefaultItemView for when these come into play.
func NewDefaultItemStyles() (s DefaultItemStyles) {
	s.NormalTitle = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"}).
		Padding(0, 0, 0, 2) //nolint:mnd

	s.SelectedTitle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"}).
		Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}).
		Padding(0, 0, 0, 1)

	s.DimmedTitle = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"}).
		Padding(0, 0, 0, 2) //nolint:mnd

	s.FilterMatch = lipgloss.NewStyle().Underline(true)

	s.CheckMark = lipgloss.NewStyle().SetString("✓").
		Foreground(lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}).
		PaddingRight(1)

	s.EmptyCheckMark = lipgloss.NewStyle().SetString("").
		Foreground(lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}).
		PaddingRight(2)

	return s
}

// DefaultDelegate is a standard delegate designed to work in lists. It's
// styled by DefaultItemStyles, which can be customized as you like.
//
// The description line can be hidden by setting Description to false, which
// renders the list as single-line-items. The spacing between items can be set
// with the SetSpacing method.
//
// Setting UpdateFunc is optional. If it's set it will be called when the
// ItemDelegate called, which is called when the list's Update function is
// invoked.
//
// Settings ShortHelpFunc and FullHelpFunc is optional. They can be set to
// include items in the list's default short and full help menus.
type DefaultDelegate struct {
	Styles        DefaultItemStyles
	UpdateFunc    func(tea.Msg, *ListScreen) tea.Cmd
	ShortHelpFunc func() []key.Binding
	FullHelpFunc  func() [][]key.Binding
	height        int
	spacing       int
}

// NewDefaultDelegate creates a new delegate with default styles.
func NewDefaultDelegate() DefaultDelegate {
	const defaultHeight = 2
	const defaultSpacing = 1
	return DefaultDelegate{
		Styles:  NewDefaultItemStyles(),
		height:  defaultHeight,
		spacing: defaultSpacing,
	}
}

// SetHeight sets delegate's preferred height.
func (d *DefaultDelegate) SetHeight(i int) {
	d.height = i
}

// Height returns the delegate's preferred height.
// This has effect only if ShowDescription is true,
// otherwise height is always 1.
func (d DefaultDelegate) Height() int {
	return 1
}

// SetSpacing sets the delegate's spacing.
func (d *DefaultDelegate) SetSpacing(i int) {
	d.spacing = i
}

// Spacing returns the delegate's spacing.
func (d DefaultDelegate) Spacing() int {
	return d.spacing
}

// Update checks whether the delegate's UpdateFunc is set and calls it.
func (d DefaultDelegate) Update(msg tea.Msg, m *ListScreen) tea.Cmd {
	if d.UpdateFunc == nil {
		return nil
	}
	return d.UpdateFunc(msg, m)
}

// Render prints an item.
func (d DefaultDelegate) Render(w io.Writer, m ListScreen, index int, item domain.Item) {
	var (
		title        string
		matchedRunes []int
		s            = &d.Styles
	)

	completed := s.EmptyCheckMark.String()
	if item.Completed() {
		completed = s.CheckMark.String()
	}

	title = item.Title()

	if m.width <= 0 {
		// short-circuit
		return
	}

	// Prevent text from exceeding list width
	textwidth := m.width - s.NormalTitle.GetPaddingLeft() - s.NormalTitle.GetPaddingRight()
	title = ansi.Truncate(title, textwidth, cmd.Ellipsis)

	// Conditions
	var (
		isSelected = index == m.Index()
		isFiltered = m.FilterState() == Filtering || m.FilterState() == FilterApplied
	)

	if isFiltered && index < len(m.filteredItems) {
		// Get indices of matched characters
		matchedRunes = m.MatchesForItem(index)
		// Highlight matches
		unmatched := s.SelectedTitle.Inline(true)
		matched := unmatched.Inherit(s.FilterMatch)
		title = lipgloss.StyleRunes(title, matchedRunes, matched, unmatched)
	} else {
		title = s.DimmedTitle.Render(title)
	}

	title = completed + title

	if isSelected && m.FilterState() != Filtering {
		title = s.SelectedTitle.Render(title)
	} else {
		title = s.NormalTitle.Render(title)
	}

	fmt.Fprintf(w, "%s", title) //nolint: errcheck
}
