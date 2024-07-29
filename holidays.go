package chinese_holidays_go

import (
	"time"
)

var Inqr = NewInquirer()

type Inquire struct {
	book *book
}

func NewInquirer() Inquirer {
	return &Inquire{book: &book{}}
}

// MustHoliday checks given date is holiday or not.
func (i *Inquire) MustHoliday(date time.Time) bool {
	holiday, err := i.Holiday(date)
	if err != nil {
		panic(err)
	}
	return holiday
}

// Holiday checks given date is holiday or not.
func (i *Inquire) Holiday(date time.Time) (bool, error) {
	events, err := i.book.updateEvent()
	if err != nil {
		return false, err
	}
	if len(events) != len(i.book.getEvent()) {
		bk, err := newBookFromEvent(events)
		if err != nil {
			return false, err
		}
		i.book = bk
	}
	return i.book.Holiday(date), nil
}

// MustWorkingDay checks given date is working day or not.
func (i *Inquire) MustWorkingDay(date time.Time) bool {
	workingDay, err := i.WorkingDay(date)
	if err != nil {
		panic(err)
	}
	return workingDay
}

// WorkingDay checks given date is working day or not.
func (i *Inquire) WorkingDay(date time.Time) (bool, error) {
	events, err := i.book.updateEvent()
	if err != nil {
		return false, err
	}
	if len(events) != len(i.book.getEvent()) {
		bk, err := newBookFromEvent(events)
		if err != nil {
			return false, err
		}
		i.book = bk
	}

	return i.book.WorkingDay(date), nil
}
