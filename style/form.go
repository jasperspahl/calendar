package style

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/jasperspahl/calendar/internal/config"
)

func GetFormTheme() *huh.Theme {
	theme := huh.ThemeBase()
	theme.Focused.Title = lipgloss.NewStyle().Foreground(config.GetConfig().Colors.Accent)
	theme.Blurred.Title = lipgloss.NewStyle().Foreground(config.GetConfig().Colors.LightGray)
	theme.Focused.Base = lipgloss.NewStyle().PaddingLeft(1).BorderStyle(lipgloss.ThickBorder()).BorderLeft(true).BorderForeground(config.GetConfig().Colors.Accent)
	theme.Blurred.Description = lipgloss.NewStyle().Foreground(config.GetConfig().Colors.Muted)
	theme.Focused.Description = lipgloss.NewStyle().Foreground(config.GetConfig().Colors.Muted)
	theme.Focused.TextInput.Prompt = lipgloss.NewStyle().Foreground(config.GetConfig().Colors.Accent)
	theme.Blurred.TextInput.Prompt = lipgloss.NewStyle().Foreground(config.GetConfig().Colors.Muted)

	theme.Focused.SelectSelector = lipgloss.NewStyle().Foreground(config.GetConfig().Colors.Muted).SetString("> ")
	theme.Focused.SelectedOption = lipgloss.NewStyle().Foreground(config.GetConfig().Colors.Accent)
	theme.Focused.UnselectedOption = lipgloss.NewStyle().Foreground(config.GetConfig().Colors.Muted)

	theme.Blurred.SelectSelector = lipgloss.NewStyle().Foreground(config.GetConfig().Colors.Muted).SetString("> ")
	theme.Blurred.SelectedOption = lipgloss.NewStyle().Foreground(config.GetConfig().Colors.LightGray)
	theme.Blurred.UnselectedOption = lipgloss.NewStyle().Foreground(config.GetConfig().Colors.Muted)

	return theme
}
