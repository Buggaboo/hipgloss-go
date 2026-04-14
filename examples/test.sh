#!/bin/bash
cd "$(dirname "$0")/.."
go build -o hipgloss ./cmd/hipgloss

echo "=== Testing Pure Lipgloss Implementation ==="

echo "1. Menu with default item (should start on Green):"
./hipgloss --default-item 2 --menu "Select color" 15 40 5 \
    1 "Red" \
    2 "Green" \
    3 "Blue" \
    4 "Yellow" \
    5 "Purple"
echo "Exit: $?"

echo -e "\n2. Yes/No:"
./hipgloss --yesno "Continue?" 7 40
echo "Exit: $?"

echo -e "\n3. Input with default:"
./hipgloss --inputbox "Your name" 8 40 "Anonymous"
echo "Exit: $?"

echo -e "\n4. Password:"
./hipgloss --passwordbox "Secret" 8 40
echo "Exit: $?"

echo -e "\n5. Checklist (pre-selected):"
./hipgloss --checklist "Select" 12 50 4 \
    1 "Option A" on \
    2 "Option B" off \
    3 "Option C" on
echo "Exit: $?"

echo -e "\n6. Radiolist:"
./hipgloss --radiolist "Choose one" 10 40 3 \
    1 "Small" off \
    2 "Medium" on \
    3 "Large" off
echo "Exit: $?"

echo -e "\n7. With backtitle:"
./hipgloss --backtitle "My App v1.0" --menu "Main Menu" 15 40 3 \
    1 "Start" \
    2 "Settings" \
    3 "Quit"
echo "Exit: $?"
