package style

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/jasperspahl/calendar/internal/config"
)

var (
	TitleStyle = lipgloss.NewStyle().
			Foreground(config.GetConfig().Colors.Accent).
			Align(lipgloss.Center).
			Bold(true)
	BorderStyle = lipgloss.NewStyle().Foreground(config.GetConfig().Colors.Muted)

	// Cells
	CellStyle        = lipgloss.NewStyle().Padding(0, 1)
	CellStyleHover   = lipgloss.NewStyle().Padding(0, 1).Background(config.GetConfig().Colors.Muted)
	OutsideCellStyle = CellStyle.Foreground(config.GetConfig().Colors.Muted)

	// Event
	EventStyle = lipgloss.NewStyle().Foreground(config.GetConfig().Colors.Muted)

	// Other
	Notice = lipgloss.NewStyle().Foreground(config.GetConfig().Colors.Muted).Align(lipgloss.Center)
)
