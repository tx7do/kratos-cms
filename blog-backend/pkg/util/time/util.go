package util

import "time"

const (
	TimeLayout = "2006-01-02 15:04:05"
	DateLayout = "2006-01-02"
)

// UnixMilliToStringPtr 毫秒时间戳转字符串
func UnixMilliToStringPtr(tm *int64) *string {
	if tm == nil {
		return nil
	}
	str := time.UnixMilli(*tm).Format(TimeLayout)
	return &str
}

// StringTimeToTime 字符串转时间
func StringTimeToTime(tm *string) *time.Time {
	if tm == nil {
		return nil
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	theTime, err := time.ParseInLocation(TimeLayout, *tm, loc)
	if err != nil {
		return nil
	}
	return &theTime
}

// StringDateToTime 字符串转时间
func StringDateToTime(tm *string) *time.Time {
	if tm == nil {
		return nil
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	theTime, err := time.ParseInLocation(DateLayout, *tm, loc)
	if err != nil {
		return nil
	}
	return &theTime
}

// StringToUnixMilliInt64Ptr 字符串转毫秒时间戳
func StringToUnixMilliInt64Ptr(tm *string) *int64 {
	theTime := StringTimeToTime(tm)
	if theTime == nil {
		return nil
	}
	unixTime := theTime.UnixMilli()
	return &unixTime
}

// GetCurrentDayTime 获取今天的时间
func GetCurrentDayTime() (time.Time, time.Time) {
	now := time.Now()
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	return startDate, endDate
}

// GetCurrentMonthTime 获取本月的时间
func GetCurrentMonthTime() (time.Time, time.Time) {
	now := time.Now()
	firstDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	lastDay := firstDay.AddDate(0, 1, -1)
	endDate := time.Date(lastDay.Year(), lastDay.Month(), lastDay.Day(), 23, 59, 59, 0, now.Location())
	return firstDay, endDate
}

// GetCurrentYearTime 获取今年的时间
func GetCurrentYearTime() (time.Time, time.Time) {
	now := time.Now()
	firstDay := time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location())
	lastDay := firstDay.AddDate(1, 0, -1)
	endDate := time.Date(lastDay.Year(), lastDay.Month(), lastDay.Day(), 23, 59, 59, 0, now.Location())
	return firstDay, endDate
}

// GetCurrentDayDateString 获取今天的日期字符串
func GetCurrentDayDateString() (string, string) {
	firstDay, lastDay := GetCurrentDayTime()

	startDate := firstDay.Format(DateLayout)
	endDate := lastDay.Format(DateLayout)

	return startDate, endDate
}

// GetCurrentMonthDateString 获取本月的日期字符串
func GetCurrentMonthDateString() (string, string) {
	firstDay, lastDay := GetCurrentMonthTime()

	startDate := firstDay.Format(DateLayout)
	endDate := lastDay.Format(DateLayout)

	return startDate, endDate
}

// GetCurrentYearDateString 获取今年的日期字符串
func GetCurrentYearDateString() (string, string) {
	firstDay, lastDay := GetCurrentYearTime()

	startDate := firstDay.Format(DateLayout)
	endDate := lastDay.Format(DateLayout)

	return startDate, endDate
}

// GetCurrentDayTimeString 获取今天的时间字符串
func GetCurrentDayTimeString() (string, string) {
	firstDay, lastDay := GetCurrentDayTime()

	startDate := firstDay.Format(TimeLayout)
	endDate := lastDay.Format(TimeLayout)

	return startDate, endDate
}

// GetCurrentMonthTimeString 获取本月的时间字符串
func GetCurrentMonthTimeString() (string, string) {
	firstDay, lastDay := GetCurrentMonthTime()

	startDate := firstDay.Format(TimeLayout)
	endDate := lastDay.Format(TimeLayout)

	return startDate, endDate
}

// GetCurrentYearTimeString 获取今年的时间字符串
func GetCurrentYearTimeString() (string, string) {
	firstDay, lastDay := GetCurrentYearTime()

	startDate := firstDay.Format(TimeLayout)
	endDate := lastDay.Format(TimeLayout)

	return startDate, endDate
}
