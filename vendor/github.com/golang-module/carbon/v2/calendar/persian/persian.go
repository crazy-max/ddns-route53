// Package persian is part of the carbon package.
package persian

import (
	"fmt"
	"math"
	"time"

	"github.com/golang-module/carbon/v2/calendar"
	"github.com/golang-module/carbon/v2/calendar/julian"
)

var (
	minYear = 1097
	months  = []string{"فروردین", "اردیبهشت", "خرداد", "تیر", "مرداد", "شهریور", "مهر", "آبان", "آذر", "دی", "بهمن", "اسفند"}
	weeks   = []string{"یکشنبه", "دوشنبه", "سه شنبه", "چهارشنبه", "پنجشنبه", "جمعه", "شنبه"}

	invalidYearError = func() error {
		return fmt.Errorf("invalid year, cannot exceed 1079")
	}
)

// Gregorian defines a Gregorian struct.
// 定义 Gregorian 结构体
type Gregorian struct {
	calendar.Gregorian
}

// Persian defines a Persian struct.
// 定义 Persian 结构体
type Persian struct {
	year, month, day, hour, minute, second int
	Error                                  error
}

// FromGregorian creates a Gregorian instance from time.Time.
// 从标准 time.Time 创建 Gregorian 实例
func FromGregorian(t time.Time) (g Gregorian) {
	if t.IsZero() {
		return
	}
	g.Time = t
	return
}

// FromPersian creates a Persian instance from persian datetime.
// 从 波斯日期 创建 Persian 实例
func FromPersian(year, month, day, hour, minute, second int) (p Persian) {
	p.year, p.month, p.day = year, month, day
	p.hour, p.minute, p.second = hour, minute, second
	return p
}

// ToPersian converts Gregorian instance to Persian instance.
// 将 Gregorian 实例转化为 Persian 实例
func (g Gregorian) ToPersian() (p Persian) {
	if g.IsZero() {
		return
	}
	p.hour, p.minute, p.second = g.Hour(), g.Minute(), g.Second()
	if g.Year() < minYear {
		p.Error = invalidYearError()
		return
	}
	gjdn := getGregorianJdn(g.Year(), g.Month(), g.Day())
	pjdn := getPersianJdn(475, 1, 1)

	diff := gjdn - pjdn
	div := diff / 1029983
	mod := diff % 1029983

	p.year = (2134*mod/366+2816*(mod%366)+2815)/1028522 + mod/366 + 1 + 2820*div + 474
	pjdn = getPersianJdn(p.year, 1, 1)
	fjdn := float64(gjdn - pjdn + 1)
	if fjdn <= 186 {
		p.month = int(math.Ceil(fjdn / 31.0))
	} else {
		p.month = int(math.Ceil((fjdn - 6) / 30.0))
	}
	pjdn = getPersianJdn(p.year, p.month, 1)
	p.day = gjdn - pjdn + 1
	return
}

// ToGregorian converts Persian instance to Gregorian instance.
// 将 Persian 实例转化为 Gregorian 实例
func (p Persian) ToGregorian() (g Gregorian) {
	if p.IsZero() {
		return
	}
	jdn := getPersianJdn(p.year, p.month, p.day)

	l := jdn + 68569
	n := 4 * l / 146097
	l = l - (146097*n+3)/4
	i := 4000 * (l + 1) / 1461001
	l = l - 1461*i/4 + 31
	j := 80 * l / 2447
	d := l - 2447*j/80
	l = j / 11
	m := j + 2 - 12*l
	y := 100*(n-49) + i + l

	g.Time = time.Date(y, time.Month(m), d, p.hour, p.minute, p.second, 0, time.Local)
	return
}

// Year gets lunar year like 2020.
// 获取年份，如 2020
func (p Persian) Year() int {
	if p.Error != nil {
		return 0
	}
	return p.year
}

// Month gets lunar month like 8.
// 获取月份，如 8
func (p Persian) Month() int {
	if p.Error != nil {
		return 0
	}
	return p.month
}

// Day gets lunar day like 5.
// 获取日，如 5
func (p Persian) Day() int {
	if p.Error != nil {
		return 0
	}
	return p.day
}

// Hour gets current hour like 13.
// 获取小时，如 13
func (p Persian) Hour() int {
	if p.Error != nil {
		return 0
	}
	return p.hour
}

// Minute gets current minute like 14.
// 获取分钟数，如 14
func (p Persian) Minute() int {
	if p.Error != nil {
		return 0
	}
	return p.minute
}

// Second gets current second like 15.
// 获取秒数，如 15
func (p Persian) Second() int {
	if p.Error != nil {
		return 0
	}
	return p.second
}

// String implements Stringer interface and outputs a string in YYYY-MM-DD HH::ii::ss format like "1402-11-11 00:00:00".
// 实现 Stringer 接口, 输出 YYYY-MM-DD HH::ii::ss 格式字符串，如 "1402-11-11 00:00:00"
func (p Persian) String() string {
	if p.IsZero() {
		return ""
	}
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", p.year, p.month, p.day, p.hour, p.minute, p.second)
}

// ToMonthString outputs a string in persian month format like "فروردین".
// 获取完整月份字符串，如 "فروردین"
func (p Persian) ToMonthString() (month string) {
	if p.IsZero() {
		return ""
	}
	return months[p.month-1]
}

// ToWeekString outputs a string in week layout like "چهارشنبه".
// 输出完整星期字符串，如 "چهارشنبه"
func (p Persian) ToWeekString() (month string) {
	if p.IsZero() {
		return ""
	}
	return weeks[p.ToGregorian().Week()]
}

// IsZero reports whether is zero time.
// 是否是零值时间
func (p Persian) IsZero() bool {
	if p.Error != nil {
		return true
	}
	if p.year == 0 || p.month == 0 || p.day == 0 {
		return true
	}
	return false
}

// IsLeapYear reports whether is a leap year.
// 是否是闰年
func (p Persian) IsLeapYear() bool {
	if p.IsZero() {
		return false
	}
	return (25*p.year+11)%33 < 8
}

// gets Julian day number in Persian calendar
// 获取波斯历儒略日计数
func getPersianJdn(year, month, day int) int {
	year = year - 473
	if year >= 0 {
		year--
	}
	epy := 474 + (year % 2820)
	var md int
	if month <= 7 {
		md = (month - 1) * 31
	} else {
		md = (month-1)*30 + 6
	}
	return day + md + (epy*682-110)/2816 + (epy-1)*365 + year/2820*1029983 + 1948320
}

// gets Julian day number in Gregorian calendar
// 获取公历儒略日计数
func getGregorianJdn(year, month, day int) int {
	jdn := julian.FromGregorian(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)).ToJulian().JD(0)
	return int(jdn)
}
