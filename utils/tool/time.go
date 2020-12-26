package tool

import (
	"math"
	"strconv"
	"strings"
	"time"
)

// DateFormat pattern rules.
// Golang 格式化时间默认字符 2006-01-02 15:04:05
var datePatterns = []string{
	// year
	"Y", "2006", // A full numeric representation of a year, 4 digits   Examples: 1999 or 2003
	"y", "06", //A two digit representation of a year   Examples: 99 or 03

	// month
	"m", "01", // Numeric representation of a month, with leading zeros 01 through 12
	"n", "1", // Numeric representation of a month, without leading zeros   1 through 12
	"M", "Jan", // A short textual representation of a month, three letters Jan through Dec
	"F", "January", // A full textual representation of a month, such as January or March   January through December

	// day
	"d", "02", // Day of the month, 2 digits with leading zeros 01 to 31
	"j", "2", // Day of the month without leading zeros 1 to 31

	// week
	"D", "Mon", // A textual representation of a day, three letters Mon through Sun
	"l", "Monday", // A full textual representation of the day of the week  Sunday through Saturday

	// time
	"g", "3", // 12-hour format of an hour without leading zeros    1 through 12
	"G", "15", // 24-hour format of an hour without leading zeros   0 through 23
	"h", "03", // 12-hour format of an hour with leading zeros  01 through 12
	"H", "15", // 24-hour format of an hour with leading zeros  00 through 23

	"a", "pm", // Lowercase Ante meridiem and Post meridiem am or pm
	"A", "PM", // Uppercase Ante meridiem and Post meridiem AM or PM

	"i", "04", // Minutes with leading zeros    00 to 59
	"s", "05", // Seconds, with leading zeros   00 through 59

	// time zone
	"T", "MST",
	"P", "-07:00",
	"O", "-0700",

	// RFC 2822
	"r", time.RFC1123Z,
}

// 本日起始时间
func TodayTime() time.Time {
	return TimeDayBegin(time.Now())
}

// 本周起始时间
func WeekTime() time.Time {
	return TimeWeekBegin(TodayTime())
}

// 本周结束时间
func WeekEndTime() time.Time {
	return TimeWeekEnd(TodayTime())
}

// 本月起始时间
func MonthTime() time.Time {
	return TimeMonthBegin(TodayTime())
}

// UnixToTime 将时间戳转为time
func UnixToTime(sec int64) time.Time {
	return time.Unix(sec, 0)
}

// Datetime 将时间戳转为字符串日期
func Datetime(sec int64, format ...string) string {
	var ft string
	if len(format) > 0 {
		ft = format[0]
	} else {
		ft = "Y-m-d H:i:s"
	}
	return TimeFormat(time.Unix(sec, 0), ft)
}

// StrToUnix 将日期字符串转为时间戳
func StrToUnix(datetime string, format ...string) int64 {
	var ft string
	if len(format) > 0 {
		ft = format[0]
	} else {
		ft = "Y-m-d H:i:s"
	}

	t, err := StrToTime(datetime, ft)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// StrToTime 将日期字符串转为时间对象
func StrToTime(datetime, format string) (time.Time, error) {
	replacer := strings.NewReplacer(datePatterns...)
	format = replacer.Replace(format)
	return time.ParseInLocation(format, datetime, time.Local)
}

// TimeFormat 格式化时间戳为字符串时间
func TimeFormat(t time.Time, format string) string {
	replacer := strings.NewReplacer(datePatterns...)
	format = replacer.Replace(format)
	return t.Format(format)
}

// FormatDate 输出time对象的日期字符串
func FormatDate(t time.Time) string {
	return TimeFormat(t, "Y-m-d")
}

// FormatTime 输出time对象的时间字符串
func FormatTime(t time.Time) string {
	return TimeFormat(t, "H:i:s")
}

// FormatDatetime 输出time对象的日期时间字符串
func FormatDatetime(t time.Time) string {
	return TimeFormat(t, "Y-m-d H:i:s")
}

// TimeUnix 获取当前时区的系统时间戳
func TimeUnix() int64 {
	return time.Now().Unix()
}

// TimeUnixNano 获取纳秒时间戳
func TimeUnixNano() int64 {
	return time.Now().UnixNano()
}

// TimeUnixNano 获取毫秒时间戳
func TimeUnixMill() int64 {
	return TimeUnixNano() / 1e6
}

// DiffDays 获取两个时间相隔天数
func DiffDays(to, from time.Time) int {
	return int(to.Sub(from).Hours()) / 24
}

// IsSameDay 是否同年同月同日
func IsSameDay(t1, t2 time.Time) bool {
	return IsSameYear(t1, t2) && t1.YearDay() == t2.YearDay()
}

// IsSameMonth 是否同年同月
func IsSameMonth(t1, t2 time.Time) bool {
	return IsSameYear(t1, t2) && t1.Month() == t2.Month()
}

// IsSameDay 是否同年
func IsSameYear(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year()
}

// 取一天的起始时间(当天的00:00:00)
func TimeDayBegin(t time.Time) time.Time {
	r := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	return r
}

// 取一天的特定hour小时的时间(当天的HH:00:00)
func TimeDayHour(t time.Time, hour int) time.Time {
	r := time.Date(t.Year(), t.Month(), t.Day(), hour, 0, 0, 0, time.Local)
	return r
}

// 取一周的起始时间(当周的周一)
func TimeWeekBegin(t time.Time) time.Time {
	wDay := int(t.Weekday())

	// sunday
	if wDay == 0 {
		wDay = 7
	}

	wDay--

	r := t.AddDate(0, 0, -wDay)
	return r
}

func TimeWeekEnd(t time.Time) time.Time {
	cur := int(t.Weekday())
	// sunday
	if cur == 0 {
		cur = 7
	}
	cur--
	r := t.AddDate(0, 0, 7-cur)
	return r
}

// 取一月的起始时间(当月1日)
func TimeMonthBegin(t time.Time) time.Time {
	r := t.AddDate(0, 0, -t.Day()+1)
	return r
}

// 取当月最后一天(当月1日)
func TimeMonthEnd(t time.Time) time.Time {
	r := t.AddDate(0, 1, -t.Day())
	return r
}

// 取指定月的第X个周Y日期
func TimeFindWeekDay(t time.Time, whichWeek int, whichWDay time.Weekday) time.Time {
	if whichWeek < 1 {
		return time.Now()
	}

	// 当月第一天
	firstOfMonth := TimeMonthBegin(t)
	firstOfWDay := firstOfMonth.Weekday()

	offsetDay := int(whichWDay - firstOfWDay)
	if offsetDay < 0 {
		// 差了一轮,补一周天数
		offsetDay += 7
	}

	destTime := firstOfMonth.AddDate(0, 0, offsetDay+(whichWeek-1)*7)
	return destTime
}

// 取指定月的倒数第X个周Y日期
func TimeRFindWeekDay(t time.Time, whichWeek int, whichWDay time.Weekday) time.Time {
	if whichWeek < 1 {
		return time.Now()
	}

	// 当月最后一天
	endOfMonth := TimeMonthEnd(t)
	endOfWDay := endOfMonth.Weekday()

	offsetDay := int(endOfWDay - whichWDay)
	if offsetDay < 0 {
		// 差了一轮,补一周天数
		offsetDay += 7
	}

	destTime := endOfMonth.AddDate(0, 0, -(offsetDay + (whichWeek-1)*7))
	return destTime
}

// 时间戳转字符串切片[YYYY,MM,DD,hh,mm,ss]
func TimeToStrSlice(u int64) []string {
	t := time.Unix(u, 0)
	timeArgs := make([]string, 6)
	timeArgs[0] = strconv.Itoa(t.Year())
	timeArgs[1] = strconv.Itoa(int(t.Month()))
	timeArgs[2] = strconv.Itoa(t.Day())
	timeArgs[3] = strconv.Itoa(t.Hour())
	timeArgs[4] = strconv.Itoa(t.Minute())
	timeArgs[5] = strconv.Itoa(t.Second())
	return timeArgs
}

// 时间戳转成剩余多少分钟，结果向上取整
func TimeToMinute(u int64) int64 {
	return int64(math.Ceil(float64(u) / float64(60)))
}

//获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

//获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

//获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}
