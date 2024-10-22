package newevent

import (
	"github.com/jasperspahl/calendar/style"
)

func (m Model) View() string {
	var str string

	title := style.TitleStyle.Render("# New Event") + "\n\n"

	form := m.form.View()

	str += title + form

	return str
}
