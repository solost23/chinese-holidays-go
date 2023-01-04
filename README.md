# chinese-holidays-go

## Install 
    go get github.com/solost23/chinese-holidays-go
##

## Usage

```go
import (
	chg "github.com/solost23/chinese-holidays-go"
)

func main() {
	inquirer := NewInquirer()
	date := time.Date(2023, 1, 2, 0, 0, 0, 0, china)
	inquirer.IsHoliday(date) // true
	inquirer.IsWorkingDay(date) // false
}
```