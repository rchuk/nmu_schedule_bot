package api

import "time"

type RawScheduleEntry struct {
	DiscScheduleContentId string    `json:"discScheduleContentId"`
	DisciplineName        string    `json:"disciplineName"`
	StudyTimeName         string    `json:"studyTimeName"`
	StudyTimeBegin        time.Time `json:"studyTimeBegin"`
	StudyTimeEnd          time.Time `json:"studyTimeEnd"`
	ScheduleDate          time.Time `json:"scheduleDate"`
	CabinetNumber         string    `json:"cabinetNumber"`
	PositionName          *string   `json:"positionName"`
	PositionShortName     *string   `json:"positionShortName"`
	EmpFullName           *string   `json:"empFullName"`
	LastName              *string   `json:"lastName"`
	FirstName             *string   `json:"firstName"`
	MiddleName            *string   `json:"middleName"`
	SubgroupName          *string   `json:"subgroupName"`
	ContentNotes          *string   `json:"contentNotes"`
	StudyTypeName         string    `json:"studyTypeName"`
}

type RawScheduleRequest struct {
	DateFrom time.Time `json:"dateFrom"`
	DateTo   time.Time `json:"dateTo"`
}

type RawAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}
