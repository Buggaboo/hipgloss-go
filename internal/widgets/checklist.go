package widgets

import (
	"strings"
	
	tea "github.com/charmbracelet/bubbletea"
	"github.com/Buggaboo/hipgloss-go/internal/config"
	"github.com/Buggaboo/hipgloss-go/internal/theme"
)

type ChecklistWidget struct {
	BaseWidget
	Theme    theme.Theme
	Cursor   int
	Selected map[int]bool
	Viewport struct {
		Start int
		End   int
	}
}

func NewChecklist(cfg config.Config) *ChecklistWidget {
	selected := make(map[int]bool)
	for i, item := range cfg.Items {
		if item.Status {
			selected[i] = true
		}
	}
	
	w := &ChecklistWidget{
		BaseWidget: BaseWidget{Cfg: cfg, Width: cfg.Width, Height: cfg.Height},
		Theme:      theme.Default(),
		Cursor:     0,
		Selected:   selected,
	}
	w.updateViewport()
	return w
}

func (c *ChecklistWidget) updateViewport() {
	visible := c.Height - 4
	if visible < 1 {
		visible = 1
	}
	c.Viewport.Start = max(0, c.Cursor - visible/2)
	c.Viewport.End = min(len(c.Cfg.Items), c.Viewport.Start + visible)
}

func (c *ChecklistWidget) Init() tea.Cmd { return nil }

func (c *ChecklistWidget) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if c.Cursor > 0 {
				c.Cursor--
				c.updateViewport()
			}
		case "down", "j":
			if c.Cursor < len(c.Cfg.Items)-1 {
				c.Cursor++
				c.updateViewport()
			}
		case " ":
			c.Selected[c.Cursor] = !c.Selected[c.Cursor]
		case "enter":
			// Build output string of selected tags
			var selected []string
			for i, item := range c.Cfg.Items {
				if c.Selected[i] {
					selected = append(selected, item.Tag)
				}
			}
			c.value = strings.Join(selected, " ")
			return c, tea.Quit
		case "esc", "ctrl+c":
			c.err = config.ErrCancelled
			return c, tea.Quit
		}
	case tea.WindowSizeMsg:
		c.SetSize(msg.Width, msg.Height)
		c.Theme = c.Theme.ApplyFrame(c.Width, c.Height, c.Cfg.BackTitle != "")
		c.updateViewport()
	}
	return c, nil
}

func (c *ChecklistWidget) View() string {
	var b strings.Builder
	
	b.WriteString(c.Theme.Title.Render(" " + c.Cfg.Title + " "))
	b.WriteString("\n\n")
	
	for i := c.Viewport.Start; i < c.Viewport.End; i++ {
		item := c.Cfg.Items[i]
		
		// Checkbox
		checkbox := c.Theme.CheckboxOff
		if c.Selected[i] {
			checkbox = c.Theme.CheckboxOn
		}
		
		line := checkbox.String() + " " + item.Text
		
		if i == c.Cursor {
			b.WriteString(c.Theme.Selected.Render("> " + line))
		} else {
			b.WriteString(c.Theme.Unselected.Render("  " + line))
		}
		if i < c.Viewport.End-1 {
			b.WriteString("\n")
		}
	}
	
	// Help
	b.WriteString("\n\n")
	b.WriteString(c.Theme.Help.Render("Space: toggle  Enter: confirm  Esc: cancel"))
	
	return c.Theme.Frame.Render(b.String())
}
