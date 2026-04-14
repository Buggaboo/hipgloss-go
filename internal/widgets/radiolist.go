package widgets

import (
	"strings"
	
    tea "github.com/charmbracelet/bubbletea"
	"github.com/Buggaboo/hipgloss-go/internal/config"
	"github.com/Buggaboo/hipgloss-go/internal/theme"
)

type RadiolistWidget struct {
	BaseWidget
	Theme    theme.Theme
	Cursor   int
	Selected int // -1 if none
	Viewport struct {
		Start int
		End   int
	}
}

func NewRadiolist(cfg config.Config) *RadiolistWidget {
	selected := -1
	for i, item := range cfg.Items {
		if item.Status {
			selected = i
			break
		}
	}
	
	w := &RadiolistWidget{
		BaseWidget: BaseWidget{Cfg: cfg, Width: cfg.Width, Height: cfg.Height},
		Theme:      theme.Default(),
		Cursor:     0,
		Selected:   selected,
	}
	w.updateViewport()
	return w
}

func (r *RadiolistWidget) updateViewport() {
	visible := r.Height - 4
	if visible < 1 {
		visible = 1
	}
	r.Viewport.Start = max(0, r.Cursor - visible/2)
	r.Viewport.End = min(len(r.Cfg.Items), r.Viewport.Start + visible)
}

func (r *RadiolistWidget) Init() tea.Cmd { return nil }

func (r *RadiolistWidget) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if r.Cursor > 0 {
				r.Cursor--
				r.updateViewport()
			}
		case "down", "j":
			if r.Cursor < len(r.Cfg.Items)-1 {
				r.Cursor++
				r.updateViewport()
			}
		case " ":
			r.Selected = r.Cursor
		case "enter":
			if r.Selected >= 0 {
				r.value = r.Cfg.Items[r.Selected].Tag
			} else if len(r.Cfg.Items) > 0 {
				r.value = r.Cfg.Items[0].Tag // Default to first if none selected
			}
			return r, tea.Quit
		case "esc", "ctrl+c":
			r.err = config.ErrCancelled
			return r, tea.Quit
		}
	case tea.WindowSizeMsg:
		r.SetSize(msg.Width, msg.Height)
		r.Theme = r.Theme.ApplyFrame(r.Width, r.Height, r.Cfg.BackTitle != "")
		r.updateViewport()
	}
	return r, nil
}

func (r *RadiolistWidget) View() string {
	var b strings.Builder
	
	b.WriteString(r.Theme.Title.Render(" " + r.Cfg.Title + " "))
	b.WriteString("\n\n")
	
	for i := r.Viewport.Start; i < r.Viewport.End; i++ {
		item := r.Cfg.Items[i]
		
		radio := r.Theme.RadioOff
		if i == r.Selected {
			radio = r.Theme.RadioOn
		}
		
		line := radio.String() + " " + item.Text
		
		if i == r.Cursor {
			b.WriteString(r.Theme.Selected.Render("> " + line))
		} else {
			b.WriteString(r.Theme.Unselected.Render("  " + line))
		}
		if i < r.Viewport.End-1 {
			b.WriteString("\n")
		}
	}
	
	b.WriteString("\n\n")
	b.WriteString(r.Theme.Help.Render("Space: select  Enter: confirm  Esc: cancel"))
	
	return r.Theme.Frame.Render(b.String())
}
