package util

import (
	"fmt"
	"testing"
	"time"

	"kratos-blog/pkg/util/trans"
)

func TestUnixMilliToStringPtr(t *testing.T) {
	now := time.Now().UnixMilli()
	str := UnixMilliToStringPtr(&now)
	fmt.Println(now)
	fmt.Println(*str)

	fmt.Println(*UnixMilliToStringPtr(trans.Int64(1677135885288)))
	fmt.Println(*UnixMilliToStringPtr(trans.Int64(1677647430853)))
	fmt.Println(*UnixMilliToStringPtr(trans.Int64(1677647946234)))

	fmt.Println(*StringToUnixMilliInt64Ptr(trans.String("2023-03-01 00:00:00")))

	fmt.Println(*StringTimeToTime(trans.String("2023-03-01 00:00:00")))
	fmt.Println(*StringDateToTime(trans.String("2023-03-01")))
}

func TestGetCurrentDateString(t *testing.T) {
	fmt.Println(GetCurrentDayDateString())
	fmt.Println(GetCurrentMonthDateString())
	fmt.Println(GetCurrentYearDateString())
}

func TestGetCurrentTime(t *testing.T) {
	fmt.Println(GetCurrentDayTime())
	fmt.Println(GetCurrentMonthTime())
	fmt.Println(GetCurrentYearTime())
}

func TestGetCurrentTimeString(t *testing.T) {
	fmt.Println(GetCurrentDayTimeString())
	fmt.Println(GetCurrentMonthTimeString())
	fmt.Println(GetCurrentYearTimeString())
}
