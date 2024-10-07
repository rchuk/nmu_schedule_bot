package bot

import (
	"fmt"
	"nmu_schedule_bot/schedule"
	"strings"
	"time"
)

func FormatWeek(week *schedule.ScheduleWeek) string {
	res := strings.Builder{}

	startDay := ""
	endDay := ""
	if len(week.Days) != 0 {
		startDay = formatDate(week.Days[0].Date)
		endDay = formatDate(week.Days[len(week.Days)-1].Date)
	}

	fmt.Fprintf(&res, "<b>Тиждень</b> %s-%s\n\n", startDay, endDay)
	days := make([]string, 0, len(week.Days))
	for _, day := range week.Days {
		days = append(days, FormatDay(&day))
	}
	fmt.Fprintf(&res, strings.Join(days, "\n"))

	return res.String()
}

func FormatDay(day *schedule.ScheduleDay) string {
	res := strings.Builder{}

	weekday := day.Date.Weekday()
	padding := strings.Repeat(" ", getWeekdayPadding(weekday))
	weekdayName := getWeekdayName(weekday)
	date := formatDate(day.Date)

	fmt.Fprintf(&res, "⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯\n")
	fmt.Fprintf(&res, "%s<b>%s</b> %s📌\n", padding, weekdayName, date)
	fmt.Fprintf(&res, "⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯\n")

	if len(day.Entries) == 0 {
		fmt.Fprint(&res, "Занять немає")
	} else {
		entries := make([]string, 0, len(day.Entries))
		for _, entry := range day.Entries {
			entries = append(entries, FormatEntry(&entry))
		}
		fmt.Fprintf(&res, strings.Join(entries, "\n"))
	}

	return res.String()
}

func FormatEntry(entry *schedule.ScheduleEntry) string {
	res := strings.Builder{}

	startTime := formatTime(entry.StudyTimeBegin)
	endTime := formatTime(entry.StudyTimeEnd)

	fmt.Fprintf(&res, "  🕙%s-%s | %s\n", startTime, endTime, entry.StudyTimeName)
	fmt.Fprintf(&res, "    <b>%s</b>\n", strings.ReplaceAll(entry.DisciplineName, "\n", " "))
	fmt.Fprintf(&res, "    <i>(%s)</i>\n", entry.StudyTypeName)
	if len(entry.CabinetNumber) != 0 && entry.CabinetNumber != "-" {
		fmt.Fprintf(&res, " | Аудиторія: %s", entry.CabinetNumber)
	}

	return res.String()
}

func formatDate(time time.Time) string {
	return time.Format("02.01")
}

func formatTime(time time.Time) string {
	return time.Format("15:04")
}

func getWeekdayName(weekday time.Weekday) string {
	switch weekday {
	case time.Monday:
		return "Понеділок"
	case time.Tuesday:
		return "Вівторок"
	case time.Wednesday:
		return "Середа"
	case time.Thursday:
		return "Четвер"
	case time.Friday:
		return "П'ятниця"
	case time.Saturday:
		return "Субота"
	case time.Sunday:
		return "Неділя"
	default:
		panic("Unhandled weekday")
	}
}

func getWeekdayPadding(weekday time.Weekday) int {
	switch weekday {
	case time.Monday:
		return 24
	case time.Tuesday:
		return 25
	case time.Wednesday:
		return 26
	case time.Thursday:
		return 26
	case time.Friday:
		return 25
	case time.Saturday:
		return 26
	case time.Sunday:
		return 28
	default:
		panic("Unhandled weekday")
	}
}
