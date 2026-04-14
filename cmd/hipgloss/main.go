package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/Buggaboo/hipgloss-go/internal/config"
	"github.com/Buggaboo/hipgloss-go/internal/theme"
	"github.com/Buggaboo/hipgloss-go/internal/widgets"
	"github.com/muesli/termenv"
)

type App struct {
	cfg    config.Config
	widget widgets.Widget
	theme  theme.Theme
}	

func main() {
	cfg := parseArgs()
	
	// Set up lipgloss renderer
	lipgloss.SetColorProfile(termenv.NewOutput(os.Stderr).Profile)
	
	var w widgets.Widget
	
	switch cfg.Widget {
	case config.WidgetMenu:
		w = widgets.NewMenu(cfg)
	case config.WidgetYesNo:
		w = widgets.NewYesNo(cfg)
	case config.WidgetInputBox:
		w = widgets.NewInput(cfg, false)
	case config.WidgetPasswordBox:
		w = widgets.NewInput(cfg, true)
	case config.WidgetChecklist:
		w = widgets.NewChecklist(cfg)
	case config.WidgetRadiolist:
		w = widgets.NewRadiolist(cfg)
	case config.WidgetMsgBox:
		w = widgets.NewMsgBox(cfg)
	case config.WidgetInfoBox:
		fmt.Fprintln(os.Stderr, "Error: --infobox not implemented yet (use --msgbox)")
		os.Exit(255)
	case config.WidgetGauge:
		fmt.Fprintln(os.Stderr, "Error: --gauge not implemented yet")
		os.Exit(255)
	default:
		printUsage()
		os.Exit(255)
	}
	
	p := tea.NewProgram(w, tea.WithAltScreen())
	
	m, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(255)
	}
	
	widget := m.(widgets.Widget)
	
	// Handle clear screen
	if cfg.ClearScreen {
		fmt.Fprint(os.Stderr, "\033[H\033[2J")
	}
	
	// Exit codes
	if widget.Error() != nil {
		if widget.Error() == config.ErrCancelled {
			os.Exit(1)
		}
		os.Exit(255)
	}
	
	if widget.Value() != "" {
		fmt.Fprintln(os.Stderr, widget.Value())
	}
	
	os.Exit(0)
}

func parseArgs() config.Config {
	cfg := config.Config{}
	args := os.Args[1:]
	
	if len(args) == 0 {
		printUsage()
		os.Exit(1)
	}
	
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--menu":
			cfg.Widget = config.WidgetMenu
			cfg.Title = getArg(args, &i)
			cfg.Height = atoi(getArg(args, &i))
			cfg.Width = atoi(getArg(args, &i))
			_ = getArg(args, &i) // menu-height (ignored, we auto-calc)
			
			// Parse tag/item pairs
			for i+2 < len(args) && !strings.HasPrefix(args[i+1], "--") {
				tag := getArg(args, &i)
				text := getArg(args, &i)
				cfg.Items = append(cfg.Items, config.ListItem{Tag: tag, Text: text})
			}
			
		case "--yesno":
			cfg.Widget = config.WidgetYesNo
			cfg.Title = getArg(args, &i)
			cfg.Height = atoi(getArg(args, &i))
			cfg.Width = atoi(getArg(args, &i))
			
		case "--msgbox":
			cfg.Widget = config.WidgetMsgBox
			cfg.Title = getArg(args, &i)
			cfg.Height = atoi(getArg(args, &i))
			cfg.Width = atoi(getArg(args, &i))
			
		case "--inputbox":
			cfg.Widget = config.WidgetInputBox
			cfg.Title = getArg(args, &i)
			cfg.Height = atoi(getArg(args, &i))
			cfg.Width = atoi(getArg(args, &i))
			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "--") {
				cfg.InitialValue = args[i+1]
				i++
			}
			
		case "--passwordbox":
			cfg.Widget = config.WidgetPasswordBox
			cfg.Title = getArg(args, &i)
			cfg.Height = atoi(getArg(args, &i))
			cfg.Width = atoi(getArg(args, &i))
			
		case "--checklist":
			cfg.Widget = config.WidgetChecklist
			cfg.Title = getArg(args, &i)
			cfg.Height = atoi(getArg(args, &i))
			cfg.Width = atoi(getArg(args, &i))
			_ = getArg(args, &i) // list-height ignored
			
			for i+3 < len(args) && !strings.HasPrefix(args[i+1], "--") {
				tag := getArg(args, &i)
				text := getArg(args, &i)
				status := getArg(args, &i) == "on"
				cfg.Items = append(cfg.Items, config.ListItem{Tag: tag, Text: text, Status: status})
			}
			
		case "--radiolist":
			cfg.Widget = config.WidgetRadiolist
			cfg.Title = getArg(args, &i)
			cfg.Height = atoi(getArg(args, &i))
			cfg.Width = atoi(getArg(args, &i))
			_ = getArg(args, &i) // list-height ignored
			
			for i+3 < len(args) && !strings.HasPrefix(args[i+1], "--") {
				tag := getArg(args, &i)
				text := getArg(args, &i)
				status := getArg(args, &i) == "on"
				cfg.Items = append(cfg.Items, config.ListItem{Tag: tag, Text: text, Status: status})
			}
			
		case "--title":
			cfg.Title = getArg(args, &i)
			
		case "--backtitle":
			cfg.BackTitle = getArg(args, &i)
			
		case "--default-item":
			cfg.DefaultItem = getArg(args, &i)
			
		case "--clear":
			cfg.ClearScreen = true
			
		case "--help":
			printUsage()
			os.Exit(0)
			
		default:
			if strings.HasPrefix(args[i], "--") {
				fmt.Fprintf(os.Stderr, "Unknown option: %s\n", args[i])
				os.Exit(255)
			}
		}
	}
	
	return cfg
}

func getArg(args []string, i *int) string {
	*i++
	if *i >= len(args) {
		fmt.Fprintf(os.Stderr, "Missing argument after %s\n", args[*i-1])
		os.Exit(255)
	}
	return args[*i]
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid number: %s\n", s)
		os.Exit(255)
	}
	return n
}

func printUsage() {
	fmt.Fprintln(os.Stderr, `hipgloss-go - Pure Lipgloss/Bubbletea implementation

Usage:
  hipgloss --menu <title> <height> <width> <menu-height> <tag> <item>...
  hipgloss --yesno <text> <height> <width>
  hipgloss --msgbox <text> <height> <width>
  hipgloss --inputbox <text> <height> <width> [init]
  hipgloss --passwordbox <text> <height> <width>
  hipgloss --checklist <text> <height> <width> <list-height> <tag> <item> <status>...
  hipgloss --radiolist <text> <height> <width> <list-height> <tag> <item> <status>...

Options:
  --title <text>       Set dialog title
  --backtitle <text>   Set background title (top of screen)
  --default-item <tag> Set default selected item
  --clear              Clear screen on exit
  
Exit codes:
  0  - OK/Yes
  1  - Cancel/No/ESC
  255 - Error`)
}
