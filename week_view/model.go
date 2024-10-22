package week

import (
	calendar "github.com/anotherhadi/markdown-calendar"
	"github.com/charmbracelet/bubbles/help"
	"github.com/jasperspahl/calendar/style"
)

type Model struct {
	Help         help.Model
	focusDay     *int
	focusMonth   *int
	focusYear    *int
	calendar     *calendar.Calendar
	keys         keyMap
	currentDay   int
	currentMonth int
	currentYear  int
	width        int
	height       int
}

func NewModel(
	currentDay, currentMonth, currentYear int,
	focusDay, focusMonth, focusYear *int,
	calendar *calendar.Calendar,
) Model {
	help := help.New()
	help.Styles = style.GetHelpStyles()

	m := Model{
		currentDay:   currentDay,
		currentMonth: currentMonth,
		currentYear:  currentYear,
		focusDay:     focusDay,
		focusMonth:   focusMonth,
		focusYear:    focusYear,
		calendar:     calendar,

		keys:   Keys,
		Help:   help,
		width:  0,
		height: 0,
	}

	return m
}
