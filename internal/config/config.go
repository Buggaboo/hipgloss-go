package config

type WidgetType int

const (
	WidgetMenu WidgetType = iota
	WidgetYesNo
	WidgetMsgBox
	WidgetInfoBox
	WidgetInputBox
	WidgetPasswordBox
	WidgetChecklist
	WidgetRadiolist
	WidgetGauge
)

type Config struct {
	Widget      WidgetType
	Title       string
	BackTitle   string
	Height      int
	Width       int
	MenuHeight  int // list height parameter
	Items       []ListItem
	DefaultItem string
	ClearScreen bool
	ScrollText  bool
	// Widget specific
	InitialValue string
	Percent      int // for gauge
}

type ListItem struct {
	Tag    string
	Text   string
	Status bool // for checklist/radiolist
}

var ErrCancelled = &cancelledError{}

type cancelledError struct{}

func (e *cancelledError) Error() string { return "cancelled" }
