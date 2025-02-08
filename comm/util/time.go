package util

import (
	"regexp"
	"time"
)

const (
	UnixZero = 0
)

func Unix(t time.Time) int64 {
	//CST存入 按照UTC取出 需要处理下
	localtime := t.In(time.Local)
	sec := localtime.Unix()
	if sec < 0 {
		sec = UnixZero
	}
	return sec
}

// Time 时间戳转time
func Time(sec int64) (t time.Time) {
	if sec > 0 {
		t = time.Unix(sec, 0)
	}
	return
}

func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func Format(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func FormatTimeWithoutSeconds(t time.Time) string {
	return t.Format("2006-01-02 15:04")
}

func ParseDate(str string) time.Time {
	t, err := time.ParseInLocation("2006-1-2", str, time.Local)
	if err != nil {
		return time.Now()
	}
	return t
}

type TimeFormat struct {
	Regexp string
	Format string
}

func Parse(str string) time.Time {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	if err != nil {
		return time.Now()
	}

	return t
}

func ParseTimeWithoutSeconds(str string) time.Time {
	t, err := time.ParseInLocation("2006-01-02 15:04", str, time.Local)
	if err != nil {
		return time.Now()
	}

	return t
}

func ParseEx(str string) (time.Time, error) {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	if err != nil {
		return time.Now(), err
	}

	return t, nil
}

func ParseCompatible(str string) time.Time {
	formatLst := []TimeFormat{
		{Regexp: "^(\\d{4})-(\\d{1,2})-(\\d{1,2})$", Format: "2006-1-2"},
		{Regexp: "^(\\d{4})-(\\d{1,2})-(\\d{1,2}) (\\d{1,2})$", Format: "2006-1-2 15"},
		{Regexp: "^(\\d{4})-(\\d{1,2})-(\\d{1,2}) (\\d{1,2}):(\\d{1,2})$", Format: "2006-1-2 15:4"},
		{Regexp: "^(\\d{4})-(\\d{1,2})-(\\d{1,2}) (\\d{1,2}):(\\d{1,2}):(\\d{1,2})$", Format: "2006-1-2 15:4:5"},
		{Regexp: "^(\\d{4})/(\\d{1,2})/(\\d{1,2})$", Format: "2006/1/2"},
		{Regexp: "^(\\d{4})/(\\d{1,2})/(\\d{1,2}) (\\d{1,2})$", Format: "2006/1/2 15"},
		{Regexp: "^(\\d{4})/(\\d{1,2})/(\\d{1,2}) (\\d{1,2}):(\\d{1,2})$", Format: "2006/1/2 15:4"},
		{Regexp: "^(\\d{4})/(\\d{1,2})/(\\d{1,2}) (\\d{1,2}):(\\d{1,2}):(\\d{1,2})$", Format: "2006/1/2 15:4:5"},
	}

	var t time.Time
	var match bool
	for _, v := range formatLst {
		regx, err := regexp.Compile(v.Regexp)
		if err == nil {
			if regx.MatchString(str) {
				t, err = time.ParseInLocation(v.Format, str, time.Local)
				if err == nil {
					match = true
					break
				}
			}
		}
	}

	if !match {
		return time.Now()
	}

	return t
}

func Zero() int64 {
	curTime := time.Now()
	return time.Date(curTime.Year(), curTime.Month(), curTime.Day(), 0, 0, 0, 0, time.Local).Unix()
}

func ZeroByTimestamp(timestamp int64) int64 {
	if timestamp == 0 {
		timestamp = time.Now().Unix()
	}
	t := time.Unix(timestamp, 0)
	timeStr := t.Format("2006-01-02")
	ti, _ := time.Parse("2006-01-02", timeStr)
	return ti.Unix() - (8 * 60 * 60)
}

func Night() int64 {
	return Zero() + 86400
}

func NightByTimestamp(timestamp int64) int64 {
	return ZeroByTimestamp(timestamp) + 86400
}

func DaysIn(t time.Time) int {
	return time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}

func StartOfWeek(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day()-int(t.Weekday()), 0, 0, 0, 0, t.Location())
}

func StartOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func EndOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 0, 23, 59, 59, 999999999, t.Location())
}

func StartOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, t.Location())
}

func EndOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), time.December, 31, 23, 59, 59, 999999999, t.Location())
}

func UnixTostr(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}

// GetLastDayOfMonth 获取本月最后一天
func GetLastDayOfMonth(tTime time.Time) time.Time {
	year, month, _ := tTime.Date()
	if month == time.December { // 如果当前月份是12月
		year++               // 年份加1
		month = time.January // 月份变为1月（January）
	} else {
		month++ // 否则，月份加1
	}
	// 下个月1号
	nextFirstDay := time.Date(year, month, 1, 0, 0, 0, 0, tTime.Location())
	return nextFirstDay.AddDate(0, 0, -1) // 减去一天得到本月最后一天
}
