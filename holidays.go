package chinese_holidays_go

import (
	"time"
)

type Inquire struct {
	book *book
}

// IsHoliday checks given date is holiday or not.
// Deprecated: Use Inquirer.IsHoliday instead.
func (i *Inquire) IsHoliday(date time.Time) bool {
	events, err := i.book.updateEvent()
	if err != nil {
		panic(err)
	}
	if len(events) != len(i.book.getEvent()) {
		bk, err := newBookFromEvent(events)
		if err != nil {
			panic(err)
		}
		i.book = bk
	}
	return i.book.IsHoliday(date)
}

// IsWorkingDay checks given date is working day or not.
// Deprecated: Use Inquirer.IsWorkingDay instead.
func (i *Inquire) IsWorkingDay(date time.Time) bool {
	events, err := i.book.updateEvent()
	if err != nil {
		panic(err)
	}
	if len(events) != len(i.book.getEvent()) {
		bk, err := newBookFromEvent(events)
		if err != nil {
			panic(err)
		}
		i.book = bk
	}

	return i.book.IsWorkingDay(date)
}

func NewInquirer() Inquirer {
	return &Inquire{book: &book{}}
}

// not use
func (i *Inquire) updateEvent() ([]event, error) {
	return nil, nil
}

func (i *Inquire) getEvent() []event {
	return nil
}
