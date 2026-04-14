# hipgloss-go

A pure Go implementation of whiptail/dialog-style terminal UI widgets using [Lipgloss](https://github.com/charmbracelet/lipgloss) and [Bubbletea](https://github.com/charmbracelet/bubbletea). Built with the beautiful [Catppuccin](https://catppuccin.com/) color palette.

## Features

- **Multiple Widget Types**: Menu, Yes/No, Message Box, Input Box, Password Box, Checklist, and Radiolist
- **Beautiful Theming**: Catppuccin Mocha theme with rounded borders and consistent styling
- **Keyboard Navigation**: Full vim-style keybindings (j/k, g/G) plus arrow keys
- **Viewport Scrolling**: Automatic viewport management for long lists
- **Whiptail Compatible**: Drop-in replacement for whiptail with matching exit codes
- **Backtitle Support**: Display context information at the top of the screen
- **Default Selections**: Pre-select items in menus and lists

## Installation

```bash
go install github.com/Buggaboo/hipgloss-go/cmd/hipgloss@latest
```

Or build from source:

```bash
git clone https://github.com/Buggaboo/hipgloss-go
cd hipgloss-go
go build -o hipgloss ./cmd/hipgloss
```

## Usage

### Menu

Display a scrollable menu with selectable items:

```bash
hipgloss --menu "Select an option" 15 40 5 \
    1 "Start Game" \
    2 "Load Save" \
    3 "Settings" \
    4 "Help" \
    5 "Quit"
```

With default item and backtitle:

```bash
hipgloss --backtitle "My App v1.0" \
    --default-item 2 \
    --menu "Main Menu" 15 40 3 \
    1 "Start" \
    2 "Settings" \
    3 "Quit"
```

### Yes/No Dialog

Simple confirmation dialog:

```bash
hipgloss --yesno "Do you want to continue?" 7 40
```

### Message Box

Display a message with OK button:

```bash
hipgloss --msgbox "Operation completed successfully!" 7 50
```

### Input Box

Text input with optional default value:

```bash
# Basic input
hipgloss --inputbox "Enter your name" 8 40

# With default value
hipgloss --inputbox "Enter your name" 8 40 "Anonymous"
```

### Password Box

Masked input for sensitive data:

```bash
hipgloss --passwordbox "Enter password" 8 40
```

### Checklist

Multi-select with checkboxes:

```bash
hipgloss --checklist "Select features" 12 50 4 \
    1 "Enable logging" on \
    2 "Debug mode" off \
    3 "Auto-update" on \
    4 "Notifications" off
```

### Radiolist

Single-select radio buttons:

```bash
hipgloss --radiolist "Choose size" 10 40 3 \
    1 "Small" off \
    2 "Medium" on \
    3 "Large" off
```

## Global Options

| Option | Description |
|--------|-------------|
| `--title <text>` | Set dialog title |
| `--backtitle <text>` | Set background title (top of screen) |
| `--default-item <tag>` | Set default selected item |
| `--clear` | Clear screen on exit |
| `--help` | Display help message |

## Keyboard Controls

### Navigation

| Key | Action |
|-----|--------|
| `↑` / `k` | Move up |
| `↓` / `j` | Move down |
| `Home` / `g` | Go to first item |
| `End` / `G` | Go to last item |
| `Tab` | Switch focus (buttons) |
| `←` / `→` | Switch between buttons |

### Actions

| Key | Action |
|-----|--------|
| `Enter` | Confirm/OK |
| `Space` | Toggle checkbox/radio |
| `Esc` / `Ctrl+C` | Cancel |

### Input Boxes

| Key | Action |
|-----|--------|
| `←` / `→` | Move cursor |
| `Backspace` | Delete previous character |
| `Delete` | Delete current character |
| `Home` | Go to start |
| `End` | Go to end |

## Exit Codes

| Code | Meaning |
|------|---------|
| `0` | OK / Yes |
| `1` | Cancel / No / ESC |
| `255` | Error |

## Theming

hipgloss-go uses the Catppuccin Mocha color palette by default:

- **Frame**: Lavender borders with base background
- **Title**: Bold lavender background with base text
- **Selected**: Rosewater foreground with surface0 background
- **Buttons**: Green for active, Surface1 for inactive
- **Checkboxes**: Green for checked, Overlay0 for unchecked
- **Radio**: Blue for selected, Overlay0 for unselected

## Example Script

```bash
#!/bin/bash

# Menu selection
CHOICE=$(hipgloss --menu "Select action" 15 40 5 \
    "install" "Install package" \
    "remove" "Remove package" \
    "update" "Update system" \
    "search" "Search packages" \
    "quit" "Exit")

case $? in
    0)
        case $CHOICE in
            install)
                hipgloss --msgbox "Installing..." 7 30
                ;;
            remove)
                hipgloss --yesno "Are you sure?" 7 30 && echo "Removing..."
                ;;
            quit)
                exit 0
                ;;
        esac
        ;;
    1)
        echo "Cancelled by user"
        exit 1
        ;;
    255)
        echo "Error occurred"
        exit 255
        ;;
esac
```

## Dependencies

- [Bubbletea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Styling library
- [Catppuccin Go](https://github.com/catppuccin/go) - Color palette
- [termenv](https://github.com/muesli/termenv) - Terminal environment detection

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by the original [whiptail](https://en.wikipedia.org/wiki/Whiptail_(software))
- Built with [Charm](https://charm.sh/)'s excellent TUI libraries
- Color scheme by [Catppuccin](https://catppuccin.com/)
