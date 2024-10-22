package newevent

import (
	"strconv"
	"strings"

	calendar "github.com/anotherhadi/markdown-calendar"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"

	"github.com/jasperspahl/calendar/utils"
)

func (m Model) Update(message tea.Msg) (Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.form.WithWidth(msg.Width)
		m.form.WithHeight(msg.Height - 4)
		return m, nil
	case tea.KeyMsg:
		if msg.String() == "esc" {
			return m, utils.ChangeFocusViewCmd(m.previousView)
		}
	}

	form, cmd := m.form.Update(message)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	if m.form.State == huh.StateCompleted {
		event := calendar.Event{}
		event.Name = m.form.GetString("name")
		event.Description = m.form.GetString("description")
		event.StartDate.Day, _ = strconv.Atoi(strings.Split(m.form.GetString("date"), "/")[0])
		event.StartDate.Month, _ = strconv.Atoi(strings.Split(m.form.GetString("date"), "/")[1])
		event.StartDate.Year, _ = strconv.Atoi(strings.Split(m.form.GetString("date"), "/")[2])
		m.calendars.AddEvent(event)
		_ = m.calendars.Write()
		return m, utils.ChangeFocusViewCmd(m.previousView)
	}

	return m, cmd
}
