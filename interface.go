package chinese_holidays_go

import (
	"time"
)

// Inquirer is the interface that wraps the Inquire method.
type Inquirer interface {
	MustHoliday(date time.Time) bool
	Holiday(date time.Time) (bool, error)
	MustWorkingDay(date time.Time) bool
	WorkingDay(date time.Time) (bool, error)
}
