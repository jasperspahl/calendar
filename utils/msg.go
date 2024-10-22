package utils

import tea "github.com/charmbracelet/bubbletea"

type ChangeFocusViewMsg struct {
	View string
}

func ChangeFocusViewCmd(view string) tea.Cmd {
	return func() tea.Msg {
		return ChangeFocusViewMsg{View: view}
	}
}

type FocusPreviousViewMsg struct{}

func FocusPreviousViewCmd() tea.Msg {
	return FocusPreviousViewMsg{}
}
