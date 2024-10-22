package style

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
	"github.com/jasperspahl/calendar/internal/config"
)

func GetHelpStyles() help.Styles {
	keyStyle := lipgloss.NewStyle().Foreground(config.GetConfig().Colors.LightGray)

	descStyle := lipgloss.NewStyle().Foreground(config.GetConfig().Colors.Muted)

	sepStyle := lipgloss.NewStyle().Foreground(config.GetConfig().Colors.Muted)

	return help.Styles{
		ShortKey:       keyStyle,
		ShortDesc:      descStyle,
		ShortSeparator: sepStyle,
		Ellipsis:       sepStyle,
		FullKey:        keyStyle,
		FullDesc:       descStyle,
		FullSeparator:  sepStyle,
	}
}
