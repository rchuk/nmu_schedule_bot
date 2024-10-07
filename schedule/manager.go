package schedule

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"nmu_schedule_bot/api"
	"sync/atomic"
	"time"
)

type scheduleWeekPair struct {
	CurrWeekSchedule ScheduleWeek
	NextWeekSchedule ScheduleWeek
}

type ScheduleManager struct {
	credentials api.Credentials
	schedule    atomic.Pointer[scheduleWeekPair]
	cron        *cron.Cron
}

func NewScheduleManager(credentials api.Credentials, cronSpec string) (*ScheduleManager, error) {
	manager := &ScheduleManager{
		credentials: credentials,
		cron:        cron.New(),
	}
	_, err := manager.cron.AddFunc(cronSpec, manager.fetchSchedules)
	if err != nil {
		return nil, err
	}

	return manager, nil
}

func (manager *ScheduleManager) Start() {
	manager.fetchSchedules()

	manager.cron.Start()
}

func (manager *ScheduleManager) TodaySchedule() *ScheduleDay {
	return manager.DaySchedule(time.Now())
}

func (manager *ScheduleManager) TomorrowSchedule() *ScheduleDay {
	return manager.DaySchedule(time.Now().AddDate(0, 0, 1))
}

func (manager *ScheduleManager) DaySchedule(date time.Time) *ScheduleDay {
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local)
	for _, day := range manager.CurrWeekSchedule().Days {
		if day.Date == date {
			return &day
		}
	}

	return &ScheduleDay{Date: date}
}

func (manager *ScheduleManager) CurrWeekSchedule() *ScheduleWeek {
	return &manager.schedule.Load().CurrWeekSchedule
}

func (manager *ScheduleManager) NextWeekSchedule() *ScheduleWeek {
	return &manager.schedule.Load().NextWeekSchedule
}

func (manager *ScheduleManager) fetchSchedules() {
	currWeekSchedule, nextWeekSchedule, err := manager.fetchTwoWeekSchedule()
	if err != nil {
		return
	}

	manager.schedule.Store(&scheduleWeekPair{
		CurrWeekSchedule: *currWeekSchedule,
		NextWeekSchedule: *nextWeekSchedule,
	})
}

func (manager *ScheduleManager) fetchTwoWeekSchedule() (*ScheduleWeek, *ScheduleWeek, error) {
	currentTime := time.Now()
	start, _ := getWeekBounds(currentTime)
	_, end := getWeekBounds(currentTime.Add(time.Duration(7*24) * time.Hour))
	entries, err := api.RawGetSchedule(manager.credentials, &api.RawScheduleRequest{
		DateFrom: start,
		DateTo:   end,
	})
	if err != nil {
		return nil, nil, err
	}

	weeks := ScheduleWeekFromEntries(entries)
	if len(weeks) != 2 {
		return nil, nil, fmt.Errorf("expected two weeks, got %d", len(weeks))
	}

	return &weeks[0], &weeks[1], nil
}

func getWeekBounds(date time.Time) (time.Time, time.Time) {
	weekDuration := time.Duration(7*24) * time.Hour
	weekStart := date.Truncate(weekDuration)
	weekEnd := weekStart.Add(weekDuration - time.Second)

	return weekStart, weekEnd
}
