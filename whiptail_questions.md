# Whiptail-Go Implementation Questions

Before finalizing the pure Lipgloss/Bubbletea implementation, please answer the following:

## 1. Theming Strategy

Should I add **auto-detection for light/dark terminal backgrounds** (using `termenv` to query the terminal) so it switches between Catppuccin Mocha (dark) and Latte (light) automatically? 

**Options:**
- [ ] Auto-detect based on terminal background
- [ ] Manual flag (e.g., `--theme=light` or `--theme=dark`)
- [ ] Both (auto-detect with override flag)
- [ ] Stick with dark theme only

## 2. The `--infobox` Implementation

Since we're no longer limited by Huh's form model, I can implement `--infobox` as a **non-blocking** dialog that:
- Displays for N seconds (auto-close)
- OR stays until any key is pressed
- Uses a timer-based Bubbletea command

**Options:**
- [ ] Implement auto-closing infobox (with `--timeout <seconds>` flag)
- [ ] Implement "press any key to close" infobox
- [ ] Skip `--infobox` entirely (not needed)
- [ ] Both modes supported

## 3. Scroll Behavior

Currently implemented: vim keys (`j/k`, `g/G`) + arrows. 

**Additional features to add?**
- [ ] Page Up/Down for faster scrolling in long lists
- [ ] Type-ahead search (press `/` then type to filter items)
- [ ] Home/End keys support (already implemented, but confirm)
- [ ] Mouse wheel scrolling (if mouse support enabled)

## 4. Backtitle Rendering

I implemented it as a subtle bar at the top. 

**Visual preference:**
- [ ] **Option A**: Traditional "window title bar" at the very top (separate from dialog)
- [ ] **Option B**: Integrated into the top border of the dialog itself (more compact)
- [ ] **Option C**: Both options via a flag (e.g., `--backtitle-style=inline` vs `--backtitle-style=separate`)

## 5. Mouse Support

Bubbletea supports mouse clicks. 

**Should I enable it so users can:**
- [ ] Click items to select them
- [ ] Click buttons directly (Yes/No, OK)
- [ ] Scroll with mouse wheel
- [ ] No mouse support (keyboard only for strict terminal compatibility)
- [ ] Add `--mouse` flag to enable optionally

## 6. Output Destination

Original whiptail outputs to **stderr**. Some users want **stdout**.

**Preference:**
- [ ] Keep stderr only (strict whiptail compatibility)
- [ ] Add `--output-fd <fd>` flag (e.g., `--output-fd 1` for stdout)
- [ ] Environment variable `WHIPTAIL_OUTPUT=stdout`
- [ ] Both flag and environment variable

## 7. Additional Features

**Any other specific requirements?**

- [ ] `--gauge` support (progress bar with stdin input for percentage)
- [ ] Multi-column menus for many items
- [ ] Built-in help dialog (triggered by F1 or `?` key)
- [ ] Persistent history for input boxes (save to `~/.whiptail_history`)
- [ ] Other: ________________

---

**Please check the boxes and reply, or copy this into a response with your selections.**
