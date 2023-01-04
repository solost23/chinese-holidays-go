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
	inquirer := NewInquirer() // inquirer应设置成全局变量，避免每次初始化从互联网获取数据初始化
	date := time.Date(2023, 1, 2, 0, 0, 0, 0, china)
	inquirer.IsHoliday(date) // true
	inquirer.IsWorkingDay(date) // false
}
```