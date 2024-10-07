package schedule

import (
	"nmu_schedule_bot/api"
	"time"
)

type ScheduleWeek struct {
	Week int
	Days []ScheduleDay
}

func ScheduleWeekFromEntries(raw []api.RawScheduleEntry) []ScheduleWeek {
	weeks := make([]ScheduleWeek, 0, 2)

	var week *ScheduleWeek
	startIndex := 0
	for i, entry := range raw {
		_, isoWeek := entry.ScheduleDate.ISOWeek()
		if week == nil {
			week = &ScheduleWeek{Week: isoWeek}
		} else if week.Week != isoWeek {
			week.Days = ScheduleDayFromEntries(raw[startIndex:i])
			weeks = append(weeks, *week)
			week = &ScheduleWeek{Week: isoWeek}

			startIndex = i
		}
	}

	if week != nil && len(weeks) == 0 {
		week.Days = ScheduleDayFromEntries(raw)
		weeks = append(weeks, *week)
	}

	return weeks
}

type ScheduleDay struct {
	Date    time.Time
	Entries []ScheduleEntry
}

func ScheduleDayFromEntries(raw []api.RawScheduleEntry) []ScheduleDay {
	days := make([]ScheduleDay, 0, 7)

	var day *ScheduleDay
	for _, entry := range raw {
		if day == nil || day.Date != entry.ScheduleDate {
			if day != nil {
				days = append(days, *day)
			}

			day = &ScheduleDay{
				Date:    entry.ScheduleDate,
				Entries: make([]ScheduleEntry, 0, 4),
			}
		}

		day.Entries = append(day.Entries, ScheduleEntryFromRaw(entry))
	}

	return days
}

type ScheduleEntry struct {
	DisciplineName string
	StudyTimeName  string
	StudyTimeBegin time.Time
	StudyTimeEnd   time.Time
	CabinetNumber  string
	StudyTypeName  string
}

func ScheduleEntryFromRaw(raw api.RawScheduleEntry) ScheduleEntry {
	return ScheduleEntry{
		DisciplineName: raw.DisciplineName,
		StudyTimeName:  raw.StudyTimeName,
		StudyTimeBegin: parseStudyTime(raw.ScheduleDate, raw.StudyTimeBegin),
		StudyTimeEnd:   parseStudyTime(raw.ScheduleDate, raw.StudyTimeEnd),
		CabinetNumber:  raw.CabinetNumber,
		StudyTypeName:  raw.StudyTypeName,
	}
}

func parseStudyTime(scheduleDate time.Time, scheduleTime time.Time) time.Time {
	year, month, day := scheduleDate.Date()
	h, m, s, ns := scheduleTime.Hour(), scheduleTime.Minute(), scheduleTime.Second(), scheduleTime.Nanosecond()

	return time.Date(year, month, day, h, m, s, ns, scheduleDate.Location())
}
