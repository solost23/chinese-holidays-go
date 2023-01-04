package chinese_holidays_go

import (
	"time"
)

// Inquirer is the interface that wraps the Inquire method.
type Inquirer interface {
	IsHoliday(date time.Time) bool
	IsWorkingDay(date time.Time) bool
	// not use
	updateEvent() ([]event, error)
	getEvent() []event
}
