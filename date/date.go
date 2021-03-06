package date

import "fmt"
import "time"
import "path/filepath"

// 日期类型，年月日三部分小整数，共四字节
type Date struct {
	Year  int16
	Month int8
	Day   int8
}

// 构造函数，从三个常规 int 构造一个日期结构
// 如果日期不合法，返回零值
func NewDate(year, month, day int) Date {
	if !Valid(year, month, day) {
		return Date{}
	}

	return Date{int16(year), int8(month), int8(day)}
}

// 获取今天的日期值
func Today() Date {
	year, month, day := time.Now().Date()
	return NewDate(year, int(month), day)
}

// 月份的英文缩写常量
const (
	Jan = 1 + iota
	Feb
	Mar
	Apr
	May
	Jun
	Jul
	Aug
	Sep
	Oct
	Nov
	Dec
)

// 每月的天数，基数，2月取28天
var MonthDays [13]int8

func init() {
	MonthDays[1] = 31
	MonthDays[2] = 29
	MonthDays[3] = 31
	MonthDays[4] = 30
	MonthDays[5] = 31
	MonthDays[6] = 30
	MonthDays[7] = 31
	MonthDays[8] = 31
	MonthDays[9] = 30
	MonthDays[10] = 31
	MonthDays[11] = 30
	MonthDays[12] = 31
}

// 判断是否闰年，参数为 int 类型的年
func IsLeap(year int) bool {
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

// 判断某个年月的最后一天（也即天数），返回天数及是否闰年
func EndDay(year, month int) (days int, leap bool) {
	if month <= 0 || month > 12 {
		return 0, false
	}
	days = int(MonthDays[month])
	leap = year > 0 && IsLeap(year)
	if month == Feb && !leap {
		days--
	}
	return
}

// 一些分隔符常量
const SEP_DASH string = "-"
const SEP_SLASH string = "/"
const SEP_IOS string = SEP_DASH
const SEP_PATH string = SEP_SLASH

// 年月日字符串化的默认分隔符
var sepField string = SEP_IOS

func init() {
	sepField = string(filepath.Separator)
}

// 设置默认的分隔符，返回旧分割符
func SepField(sep string) (old string) {
	old = sepField
	sepField = sep
	return
}

// 将日期转为字符串
func (d Date) String() string {
	return fmt.Sprintf("%04d%s%02d%s%02d", d.Year, sepField, d.Month, sepField, d.Day)
}

// 将日期转为整数数字
func (d Date) IntNum() int {
	return int(d.Year)*10000 + int(d.Month)*100 + int(d.Day)
}

// 判断日期是否合法
func (d Date) Valid() bool {
	return Valid(int(d.Year), int(d.Month), int(d.Day))
}

func Valid(year, month, day int) bool {
	switch {
	case year <= 0 || year > 9999:
		return false
	case month <= 0 || month > 12:
		return false
	case day <= 0 || int8(day) > MonthDays[month]:
		return false
	default:
		if month == Feb {
			days, _ := EndDay(year, month)
			if day > days {
				return false
			}
		}
	}
	return true
}
