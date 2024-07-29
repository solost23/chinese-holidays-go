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
	date := time.Date(2023, 1, 2, 0, 0, 0, 0, china)
	Inqr.IsHoliday(date) // true
	Inqr.IsWorkingDay(date) // false
}
```