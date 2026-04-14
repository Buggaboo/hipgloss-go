package widgets

import (
	"strings"

	lipgloss "github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/Buggaboo/hipgloss-go/internal/config"
	"github.com/Buggaboo/hipgloss-go/internal/theme"
)

type MsgBoxWidget struct {
	BaseWidget
	Theme theme.Theme
}

func NewMsgBox(cfg config.Config) *MsgBoxWidget {
	return &MsgBoxWidget{
		BaseWidget: BaseWidget{Cfg: cfg, Width: cfg.Width, Height: cfg.Height},
		Theme:      theme.Default(),
	}
}

func (m *MsgBoxWidget) Init() tea.Cmd { return nil }

func (m *MsgBoxWidget) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			return m, tea.Quit
		case "esc", "ctrl+c":
			m.err = config.ErrCancelled
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)
		m.Theme = m.Theme.ApplyFrame(m.Width, m.Height, m.Cfg.BackTitle != "")
	}
	return m, nil
}

func (m *MsgBoxWidget) View() string {
	var b strings.Builder
	
	b.WriteString(m.Theme.Title.Render(" " + m.Cfg.Title + " "))
	b.WriteString("\n\n")
	
	// Word wrap content
	//words := strings.Fields(m.Cfg.Title) // Using title as content for now
	// Actually we need content field in config, but for now reuse
	content := m.Cfg.Title
	
	b.WriteString(m.Theme.Unselected.Render(content))
	b.WriteString("\n\n")
	
	// OK button centered
	ok := m.Theme.ButtonActive.Render(" OK ")
	centered := lipgloss.PlaceHorizontal(m.Width-4, lipgloss.Center, ok)
	b.WriteString(centered)
	
	return m.Theme.Frame.Render(b.String())
}
