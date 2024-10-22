package main

import (
	calendar "github.com/anotherhadi/markdown-calendar"
	month "github.com/jasperspahl/calendar/month_view"
	newevent "github.com/jasperspahl/calendar/new_event_view"
	"github.com/jasperspahl/calendar/utils"
	week "github.com/jasperspahl/calendar/week_view"
	year "github.com/jasperspahl/calendar/year_view"
)

type model struct {
	Calendar      calendar.Calendar
	FocusDay      *int
	FocusYear     *int
	FocusMonth    *int
	CurrentView   string
	NewEventModel newevent.Model
	MonthModel    month.Model
	WeekModel     week.Model
	YearModel     year.Model
	CurrentYear   int
	CurrentMonth  int
	CurrentDay    int
	Height        int
	Width         int
}

func initModel() model {
	m := model{
		Width:  0,
		Height: 0,
	}

	m.Calendar = calendar.MergeCalendars(
		utils.PtrCalendarsToCalendars(calendar.GetPurpleCalendars()),
	)
	m.CurrentView = "month"
	m.CurrentDay, m.CurrentMonth, m.CurrentYear = calendar.Today()
	focusDay, focusMonth, focusYear := m.CurrentDay, m.CurrentMonth, m.CurrentYear
	m.FocusDay, m.FocusMonth, m.FocusYear = &focusDay, &focusMonth, &focusYear
	m.MonthModel = month.NewModel(
		m.CurrentDay,
		m.CurrentMonth,
		m.CurrentYear,
		m.FocusDay,
		m.FocusMonth,
		m.FocusYear,
		&m.Calendar,
	)
	m.WeekModel = week.NewModel(
		m.CurrentDay,
		m.CurrentMonth,
		m.CurrentYear,
		m.FocusDay,
		m.FocusMonth,
		m.FocusYear,
		&m.Calendar,
	)
	m.YearModel = year.NewModel(
		m.CurrentDay,
		m.CurrentMonth,
		m.CurrentYear,
		m.FocusDay,
		m.FocusMonth,
		m.FocusYear,
		&m.Calendar,
	)
	m.NewEventModel = newevent.NewModel(&m.Calendar, "", "")

	return m
}
