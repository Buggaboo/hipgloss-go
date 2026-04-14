package widgets

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/Buggaboo/hipgloss-go/internal/config"
)

// Widget is the common interface for all dialogs
type Widget interface {
	tea.Model
	Value() string      // Returns selected value/output
	Error() error       // Returns cancellation or other errors
	SetSize(w, h int)   // Handle terminal resize
}

// BaseWidget provides common functionality
type BaseWidget struct {
	Cfg    config.Config
	Width  int
	Height int
	err    error
	value  string
}

func (b *BaseWidget) Value() string { return b.value }
func (b *BaseWidget) Error() error  { return b.err }
func (b *BaseWidget) SetSize(w, h int) {
	b.Width = w
	b.Height = h
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
