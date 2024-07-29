package chinese_holidays_go

import (
	"errors"
	"testing"
	"time"
)

func TestIsHolidayCnt5(t *testing.T) {
	loc, _ := time.LoadLocation(Local)
	// 封装数据
	type arg struct {
		date time.Time
	}
	type want struct {
		result bool
	}
	type test struct {
		arg  arg
		want want
	}
	startDate, _ := time.ParseInLocation(time.DateOnly, "2023-01-01", loc)
	tests := []test{
		{
			arg: arg{
				date: startDate,
			},
			want: want{
				result: true,
			},
		},
		{
			arg: arg{
				date: startDate.AddDate(0, 0, 1),
			},
			want: want{
				result: true,
			},
		},
		{
			arg: arg{
				date: startDate.AddDate(0, 0, 2),
			},
			want: want{
				result: false,
			},
		},
	}
	for _, test := range tests {
		if test.want.result != Inqr.MustHoliday(test.arg.date) {
			t.Errorf("err: %v", errors.New("发生错误"))
		}
	}
}
