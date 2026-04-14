package theme

import (
	catppuccin "github.com/catppuccin/go"
	"github.com/charmbracelet/lipgloss"
)

type Theme struct {
	// Frame
	Frame     lipgloss.Style
	BackTitle lipgloss.Style
	Title     lipgloss.Style
	
	// Content
	Selected    lipgloss.Style
	Unselected  lipgloss.Style
	Disabled    lipgloss.Style
	Help        lipgloss.Style
	Error       lipgloss.Style
	
	// Widgets
	ButtonActive   lipgloss.Style
	ButtonInactive lipgloss.Style
	Input          lipgloss.Style
	Password       lipgloss.Style
	CheckboxOn     lipgloss.Style
	CheckboxOff    lipgloss.Style
	RadioOn        lipgloss.Style
	RadioOff       lipgloss.Style
	GaugeEmpty     lipgloss.Style
	GaugeFilled    lipgloss.Style
	
	// Layout
	Screen lipgloss.Style
}

func Default() Theme {
	// Use Catppuccin Mocha as default (dark theme)
	c := catppuccin.Mocha
	
	base := lipgloss.NewStyle()
	
	return Theme{
		Screen: base,
		
		Frame: base.
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(c.Lavender().Hex)).
			BorderBackground(lipgloss.Color(c.Base().Hex)).
			Background(lipgloss.Color(c.Base().Hex)).
			Foreground(lipgloss.Color(c.Text().Hex)).
			Padding(1, 2),
			
		BackTitle: base.
			Foreground(lipgloss.Color(c.Subtext0().Hex)).
			Background(lipgloss.Color(c.Mantle().Hex)).
			Padding(0, 2).
			Width(80), // Will be adjusted dynamically
			
		Title: base.
			Bold(true).
			Foreground(lipgloss.Color(c.Base().Hex)).
			Background(lipgloss.Color(c.Lavender().Hex)).
			Padding(0, 1).
			MarginBottom(1),
			
		Selected: base.
			Foreground(lipgloss.Color(c.Rosewater().Hex)).
			Background(lipgloss.Color(c.Surface0().Hex)).
			Bold(true),
			
		Unselected: base.
			Foreground(lipgloss.Color(c.Text().Hex)),
			
		Disabled: base.
			Foreground(lipgloss.Color(c.Overlay0().Hex)),
			
		ButtonActive: base.
			Foreground(lipgloss.Color(c.Base().Hex)).
			Background(lipgloss.Color(c.Green().Hex)).
			Padding(0, 2),
			
		ButtonInactive: base.
			Foreground(lipgloss.Color(c.Text().Hex)).
			Background(lipgloss.Color(c.Surface1().Hex)).
			Padding(0, 2),
			
		Input: base.
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color(c.Surface1().Hex)).
			Padding(0, 1).
			Width(30),
			
		Password: base.
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color(c.Surface1().Hex)).
			Padding(0, 1).
			Width(30),
			
		CheckboxOn: base.
			Foreground(lipgloss.Color(c.Green().Hex)).
			SetString("[✓]"),
			
		CheckboxOff: base.
			Foreground(lipgloss.Color(c.Overlay0().Hex)).
			SetString("[ ]"),
			
		RadioOn: base.
			Foreground(lipgloss.Color(c.Blue().Hex)).
			SetString("(●)"),
			
		RadioOff: base.
			Foreground(lipgloss.Color(c.Overlay0().Hex)).
			SetString("(○)"),
			
		GaugeEmpty: base.
			Foreground(lipgloss.Color(c.Surface1().Hex)).
			SetString("░"),
			
		GaugeFilled: base.
			Foreground(lipgloss.Color(c.Mauve().Hex)).
			SetString("█"),
			
		Help: base.
			Foreground(lipgloss.Color(c.Overlay1().Hex)).
			MarginTop(1),
			
		Error: base.
			Foreground(lipgloss.Color(c.Red().Hex)).
			Bold(true),
	}
}

func (t Theme) ApplyFrame(width, height int, hasBackTitle bool) Theme {
	t.Frame = t.Frame.Width(width).Height(height)
	if hasBackTitle {
		t.Frame = t.Frame.MarginTop(1)
	}
	return t
}
