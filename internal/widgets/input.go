package widgets

import (
	"strings"
	
	lipgloss "github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/Buggaboo/hipgloss-go/internal/config"
	"github.com/Buggaboo/hipgloss-go/internal/theme"
)

type InputWidget struct {
	BaseWidget
	Theme    theme.Theme
	InputText    string
	Mask     bool
	Cursor   int
}

func NewInput(cfg config.Config, mask bool) *InputWidget {
	return &InputWidget{
		BaseWidget: BaseWidget{Cfg: cfg, Width: cfg.Width, Height: cfg.Height},
		Theme:      theme.Default(),
		InputText:      cfg.InitialValue,
		Mask:       mask,
		Cursor:     len(cfg.InitialValue),
	}
}

func (i *InputWidget) Init() tea.Cmd { return nil }

func (i *InputWidget) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			i.value = i.InputText
			if i.value == "" {
				i.err = config.ErrCancelled
			}
			return i, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			i.err = config.ErrCancelled
			return i, tea.Quit
		case tea.KeyBackspace:
			if i.Cursor > 0 {
				i.InputText = i.InputText[:i.Cursor-1] + i.InputText[i.Cursor:]
				i.Cursor--
			}
		case tea.KeyLeft:
			if i.Cursor > 0 {
				i.Cursor--
			}
		case tea.KeyRight:
			if i.Cursor < len(i.InputText) {
				i.Cursor++
			}
		case tea.KeyDelete:
			if i.Cursor < len(i.InputText) {
				i.InputText = i.InputText[:i.Cursor] + i.InputText[i.Cursor+1:]
			}
		case tea.KeyHome:
			i.Cursor = 0
		case tea.KeyEnd:
			i.Cursor = len(i.InputText)
		case tea.KeyRunes:
			i.InputText = i.InputText[:i.Cursor] + string(msg.Runes) + i.InputText[i.Cursor:]
			i.Cursor += len(msg.Runes)
		}
	case tea.WindowSizeMsg:
		i.SetSize(msg.Width, msg.Height)
		i.Theme = i.Theme.ApplyFrame(i.Width, i.Height, i.Cfg.BackTitle != "")
	}
	return i, nil
}

func (i *InputWidget) View() string {
	var b strings.Builder
	
	b.WriteString(i.Theme.Title.Render(" " + i.Cfg.Title + " "))
	b.WriteString("\n\n")
	
	// Render input
	display := i.InputText
	if i.Mask {
		display = strings.Repeat("•", len(i.InputText))
	}
	
	// Show cursor
	if len(display) < i.Cursor {
		display += " "
	}
	if i.Cursor < len(display) {
		cursorChar := string(display[i.Cursor])
		before := display[:i.Cursor]
		after := display[i.Cursor+1:]
		
		cursor := lipgloss.NewStyle().
			Background(lipgloss.Color("#FFF")).
			Foreground(lipgloss.Color("#000")).
			Render(cursorChar)
		
		display = before + cursor + after
	} else {
		display += lipgloss.NewStyle().
			Background(lipgloss.Color("#FFF")).
			Render(" ")
	}
	
	style := i.Theme.Input
	if i.Mask {
		style = i.Theme.Password
	}
	
	b.WriteString(style.Render(display))
	b.WriteString("\n\n")
	b.WriteString(i.Theme.Help.Render("Enter: confirm  Esc: cancel"))
	
	return i.Theme.Frame.Render(b.String())
}
