package chinese_holidays_go

import (
	"fmt"
	"testing"
)

func TestLoadData(t *testing.T) {
	events, err := loadData(UrlPrefix)
	if err != nil {
		t.Error(err)
	}
	if len(events) != 98 {
		t.Error(len(events))
	}
	fmt.Println(events)
}
