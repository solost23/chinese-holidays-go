package chinese_holidays_go

import (
	"errors"
	"time"
)

const (
	Local          = "Asia/Shanghai"
	TypeHoliday    = "holiday"
	TypeWorkingDay = "workingDay"
)

type event struct {
	Name  string
	Type  string
	Range []string
}

func (e *event) Holiday() bool {
	return e.Type == TypeHoliday
}

func (e *event) WorkingDay() bool {
	return e.Type == TypeWorkingDay
}

func (e *event) holidayTimes() ([]time.Time, error) {
	if len(e.Range) <= 0 || len(e.Range) > 2 {
		return nil, errors.New("原始数据错误")
	}
	if len(e.Range) > 0 && len(e.Range) == 1 {
		t, err := parseTime(e.Range[0])
		if err != nil {
			return nil, err
		}
		return []time.Time{t}, nil
	} else {
		start, err := parseTime(e.Range[0])
		if err != nil {
			return nil, err
		}
		end, err := parseTime(e.Range[1])
		if err != nil {
			return nil, err
		}
		return rangeDates(start, end), nil
	}
}

func parseTime(s string) (time.Time, error) {
	location, err := time.LoadLocation(Local)
	if err != nil {
		return time.Now(), err
	}
	t, err := time.ParseInLocation(time.DateOnly, s, location)
	if err != nil {
		return t, err
	}
	return t, nil
}

func rangeDates(start, end time.Time) []time.Time {
	dates := make([]time.Time, 0)
	location, err := time.LoadLocation(Local)
	if err != nil {
		return dates
	}
	y, m, d := start.Date()
	date := time.Date(y, m, d, 0, 0, 0, 0, location)
	for ; !date.After(end); date = date.AddDate(0, 0, 1) {
		dates = append(dates, date)
	}
	return dates
}
