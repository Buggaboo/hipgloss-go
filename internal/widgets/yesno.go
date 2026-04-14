package widgets

import (
	"strings"
	
	lipgloss "github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/Buggaboo/hipgloss-go/internal/config"
	"github.com/Buggaboo/hipgloss-go/internal/theme"
)

type YesNoWidget struct {
	BaseWidget
	Theme   theme.Theme
	Focused int // 0 = Yes, 1 = No
}

func NewYesNo(cfg config.Config) *YesNoWidget {
	// Default to No (whiptail compatible)
	return &YesNoWidget{
		BaseWidget: BaseWidget{Cfg: cfg, Width: cfg.Width, Height: cfg.Height},
		Theme:      theme.Default(),
		Focused:    1,
	}
}

func (y *YesNoWidget) Init() tea.Cmd { return nil }

func (y *YesNoWidget) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left", "h", "right", "l":
			y.Focused = 1 - y.Focused // Toggle between 0 and 1
		case "tab":
			y.Focused = (y.Focused + 1) % 2
		case "enter":
			if y.Focused == 0 {
				y.value = "yes"
				return y, tea.Quit
			}
			y.value = "no"
			y.err = config.ErrCancelled
			return y, tea.Quit
		case "esc", "ctrl+c":
			y.value = "no"
			y.err = config.ErrCancelled
			return y, tea.Quit
		}
	case tea.WindowSizeMsg:
		y.SetSize(msg.Width, msg.Height)
		y.Theme = y.Theme.ApplyFrame(y.Width, y.Height, y.Cfg.BackTitle != "")
	}
	return y, nil
}

func (y *YesNoWidget) View() string {
	var b strings.Builder
	
	b.WriteString(y.Theme.Title.Render(" " + y.Cfg.Title + " "))
	b.WriteString("\n\n")
	
	// Center buttons
	yesStyle := y.Theme.ButtonInactive
	noStyle := y.Theme.ButtonInactive
	
	if y.Focused == 0 {
		yesStyle = y.Theme.ButtonActive
	} else {
		noStyle = y.Theme.ButtonActive
	}
	
	yesBtn := yesStyle.Render(" Yes ")
	noBtn := noStyle.Render(" No ")
	
	buttons := lipgloss.JoinHorizontal(lipgloss.Center, yesBtn, "  ", noBtn)
	centered := lipgloss.PlaceHorizontal(y.Width-4, lipgloss.Center, buttons)
	b.WriteString(centered)
	
	return y.Theme.Frame.Render(b.String())
}
