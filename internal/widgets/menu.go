package widgets

import (
	"strings"
	
	tea "github.com/charmbracelet/bubbletea"
	"github.com/Buggaboo/hipgloss-go/internal/config"
	"github.com/Buggaboo/hipgloss-go/internal/theme"
)

type MenuWidget struct {
	BaseWidget
	Theme    theme.Theme
	Cursor   int
	Viewport struct {
		Start int
		End   int
	}
}

func NewMenu(cfg config.Config) *MenuWidget {
	w := &MenuWidget{
		BaseWidget: BaseWidget{Cfg: cfg, Width: cfg.Width, Height: cfg.Height},
		Theme:      theme.Default(),
		Cursor:     0,
	}
	
	// Find default item
	if cfg.DefaultItem != "" {
		for i, item := range cfg.Items {
			if item.Tag == cfg.DefaultItem {
				w.Cursor = i
				break
			}
		}
	}
	
	w.updateViewport()
	return w
}

func (m *MenuWidget) updateViewport() {
	visible := m.Height - 4 // Account for borders, title, padding
	if visible < 1 {
		visible = 1
	}
	
	m.Viewport.Start = max(0, m.Cursor - visible/2)
	m.Viewport.End = min(len(m.Cfg.Items), m.Viewport.Start + visible)
	
	// Adjust if at end
	if m.Viewport.End - m.Viewport.Start < visible {
		m.Viewport.Start = max(0, m.Viewport.End - visible)
	}
}

func (m *MenuWidget) Init() tea.Cmd {
	return nil
}

func (m *MenuWidget) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
				m.updateViewport()
			}
		case "down", "j":
			if m.Cursor < len(m.Cfg.Items)-1 {
				m.Cursor++
				m.updateViewport()
			}
		case "home", "g":
			m.Cursor = 0
			m.updateViewport()
		case "end", "G":
			m.Cursor = len(m.Cfg.Items) - 1
			m.updateViewport()
		case "enter":
			m.value = m.Cfg.Items[m.Cursor].Tag
			return m, tea.Quit
		case "esc", "ctrl+c":
			m.err = config.ErrCancelled
			return m, tea.Quit
		}
		
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)
		m.Theme = m.Theme.ApplyFrame(m.Width, m.Height, m.Cfg.BackTitle != "")
		m.updateViewport()
	}
	
	return m, nil
}

func (m *MenuWidget) View() string {
	var b strings.Builder
	
	// Title
	title := m.Theme.Title.Render(" " + m.Cfg.Title + " ")
	b.WriteString(title + "\n\n")
	
	// Items
	for i := m.Viewport.Start; i < m.Viewport.End; i++ {
		item := m.Cfg.Items[i]
		cursor := "  "
		if i == m.Cursor {
			cursor = "> "
			b.WriteString(m.Theme.Selected.Render(cursor + item.Text))
		} else {
			b.WriteString(m.Theme.Unselected.Render(cursor + item.Text))
		}
		if i < m.Viewport.End-1 {
			b.WriteString("\n")
		}
	}
	
	// Scroll indicators
	if m.Viewport.Start > 0 {
		b.WriteString(m.Theme.Help.Render("\n↑ more"))
	}
	if m.Viewport.End < len(m.Cfg.Items) {
		b.WriteString(m.Theme.Help.Render("\n↓ more"))
	}
	
	content := b.String()
	
	// Apply frame
	return m.Theme.Frame.Render(content)
}
