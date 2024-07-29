package chinese_holidays_go

import (
	"time"
)

type book struct {
	events []event
	index  map[string]event
	// 设置一个更新标记，防止12-28一直更新数据
	flag bool
}

func newBookFromEvent(events []event) (*book, error) {
	index := make(map[string]event)
	for _, event := range events {
		times, err := event.holidayTimes()
		if err != nil {
			return nil, err
		}

		for _, time := range times {
			index[_key(time)] = event
		}
	}
	return &book{events: events, index: index}, nil
}

func (b *book) updateEvent() ([]event, error) {
	// TODO: 实现12月28号更新一次数据，当天便不再更新
	currentTime := time.Now()
	if currentTime.Month() == time.December && currentTime.Day() == 29 {
		b.flag = false
	}
	if b.flag == true {
		return b.events, nil
	}
	if !(len(b.events) <= 0 || (currentTime.Month() == time.December && currentTime.Day() == 28)) {
		return b.events, nil
	}
	b.flag = true
	// 从网络更新数据
	return loadData(UrlPrefix)
}

func (b *book) Holiday(date time.Time) bool {
	event := b.findEvent(date)
	if event == nil {
		return b.Weekend(date)
	}
	return event.Holiday()
}

func (b *book) WorkingDay(date time.Time) bool {
	event := b.findEvent(date)
	if event == nil {
		return !b.Weekend(date)
	}
	return event.WorkingDay()
}

func (b *book) findEvent(date time.Time) *event {
	event, ok := b.index[_key(date)]
	if !ok {
		return nil
	}
	return &event
}

func (b *book) Weekend(date time.Time) bool {
	day := date.Weekday()
	return day == time.Sunday || day == time.Saturday
}

func (b *book) getEvent() []event {
	return b.events
}

func _key(date time.Time) string {
	return date.Format(time.DateOnly)
}
