// Package lunar is part of the carbon package.
package lunar

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-module/carbon/v2/calendar"
)

var (
	minYear, maxYear = 1900, 2100

	numbers = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	months  = []string{"正", "二", "三", "四", "五", "六", "七", "八", "九", "十", "十一", "腊"}
	weeks   = []string{"周日", "周一", "周二", "周三", "周四", "周五", "周六"}
	animals = []string{"猴", "鸡", "狗", "猪", "鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊"}

	festivals = map[string]string{
		// "month-day": "name"
		"1-1":   "春节",
		"1-15":  "元宵节",
		"2-2":   "龙抬头",
		"3-3":   "上巳节",
		"5-5":   "端午节",
		"7-7":   "七夕节",
		"7-15":  "中元节",
		"8-15":  "中秋节",
		"9-9":   "重阳节",
		"10-1":  "寒衣节",
		"10-15": "下元节",
		"12-8":  "腊八节",
	}

	years = []int{
		0x04bd8, 0x04ae0, 0x0a570, 0x054d5, 0x0d260, 0x0d950, 0x16554, 0x056a0, 0x09ad0, 0x055d2, // 1900-1909
		0x04ae0, 0x0a5b6, 0x0a4d0, 0x0d250, 0x1d255, 0x0b540, 0x0d6a0, 0x0ada2, 0x095b0, 0x14977, // 1910-1919
		0x04970, 0x0a4b0, 0x0b4b5, 0x06a50, 0x06d40, 0x1ab54, 0x02b60, 0x09570, 0x052f2, 0x04970, // 1920-1929
		0x06566, 0x0d4a0, 0x0ea50, 0x16a95, 0x05ad0, 0x02b60, 0x186e3, 0x092e0, 0x1c8d7, 0x0c950, // 1930-1939
		0x0d4a0, 0x1d8a6, 0x0b550, 0x056a0, 0x1a5b4, 0x025d0, 0x092d0, 0x0d2b2, 0x0a950, 0x0b557, // 1940-1949
		0x06ca0, 0x0b550, 0x15355, 0x04da0, 0x0a5d0, 0x14573, 0x052d0, 0x0a9a8, 0x0e950, 0x06aa0, // 1950-1959
		0x0aea6, 0x0ab50, 0x04b60, 0x0aae4, 0x0a570, 0x05260, 0x0f263, 0x0d950, 0x05b57, 0x056a0, // 1960-1969
		0x096d0, 0x04dd5, 0x04ad0, 0x0a4d0, 0x0d4d4, 0x0d250, 0x0d558, 0x0b540, 0x0b5a0, 0x195a6, // 1970-1979
		0x095b0, 0x049b0, 0x0a974, 0x0a4b0, 0x0b27a, 0x06a50, 0x06d40, 0x0af46, 0x0ab60, 0x09570, // 1980-1989
		0x04af5, 0x04970, 0x064b0, 0x074a3, 0x0ea50, 0x06b58, 0x05ac0, 0x0ab60, 0x096d5, 0x092e0, // 1990-1999
		0x0c960, 0x0d954, 0x0d4a0, 0x0da50, 0x07552, 0x056a0, 0x0abb7, 0x025d0, 0x092d0, 0x0cab5, // 2000-2009
		0x0a950, 0x0b4a0, 0x0baa4, 0x0ad50, 0x055d9, 0x04ba0, 0x0a5b0, 0x15176, 0x052b0, 0x0a930, // 2010-2019
		0x07954, 0x06aa0, 0x0ad50, 0x05b52, 0x04b60, 0x0a6e6, 0x0a4e0, 0x0d260, 0x0ea65, 0x0d530, // 2020-2029
		0x05aa0, 0x076a3, 0x096d0, 0x04bd7, 0x04ad0, 0x0a4d0, 0x1d0b6, 0x0d250, 0x0d520, 0x0dd45, // 2030-2039
		0x0b5a0, 0x056d0, 0x055b2, 0x049b0, 0x0a577, 0x0a4b0, 0x0aa50, 0x1b255, 0x06d20, 0x0ada0, // 2040-2049
		0x14b63, 0x09370, 0x049f8, 0x04970, 0x064b0, 0x168a6, 0x0ea50, 0x06b20, 0x1a6c4, 0x0aae0, // 2050-2059
		0x0a2e0, 0x0d2e3, 0x0c960, 0x0d557, 0x0d4a0, 0x0da50, 0x05d55, 0x056a0, 0x0a6d0, 0x055d4, // 2060-2069
		0x052d0, 0x0a9b8, 0x0a950, 0x0b4a0, 0x0b6a6, 0x0ad50, 0x055a0, 0x0aba4, 0x0a5b0, 0x052b0, // 2070-2079
		0x0b273, 0x06930, 0x07337, 0x06aa0, 0x0ad50, 0x14b55, 0x04b60, 0x0a570, 0x054e4, 0x0d160, // 2080-2089
		0x0e968, 0x0d520, 0x0daa0, 0x16aa6, 0x056d0, 0x04ae0, 0x0a9d4, 0x0a2d0, 0x0d150, 0x0f252, // 2090-2099
		0x0d520, // 2100
	}

	invalidYearError = func() error {
		return fmt.Errorf("invalid year, outside of range [1900,2100]")
	}
)

// Gregorian defines a Gregorian struct.
// 定义 Gregorian 结构体
type Gregorian struct {
	calendar.Gregorian
}

// Lunar defines a Lunar struct.
// 定义 Lunar 结构体
type Lunar struct {
	year, month, day, hour, minute, second int
	isLeapMonth                            bool
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

// FromLunar creates a Lunar instance from lunar datetime.
// 从 农历日期 创建 Lunar 实例
func FromLunar(year, month, day, hour, minute, second int, isLeapMonth bool) (l Lunar) {
	l.year, l.month, l.day = year, month, day
	l.hour, l.minute, l.second = hour, minute, second
	l.isLeapMonth = isLeapMonth
	return
}

// ToLunar converts Gregorian instance to Lunar instance.
// 将 Gregorian 实例转化为 Lunar 实例
func (g Gregorian) ToLunar() (l Lunar) {
	if g.IsZero() {
		return
	}
	daysInYear, daysInMonth, leapMonth := 365, 30, 0
	if g.Year() < minYear || g.Year() > maxYear {
		l.Error = invalidYearError()
		return
	}
	// 与 1900-01-31 相差多少天
	offset := g.diffInDays(time.Date(minYear, 1, 31, 0, 0, 0, 0, g.Location()))
	for l.year = minYear; l.year <= maxYear && offset > 0; l.year++ {
		daysInYear = l.getDaysInYear()
		offset -= daysInYear
	}
	if offset < 0 {
		offset += daysInYear
		l.year--
	}
	leapMonth = l.LeapMonth()
	for l.month = 1; l.month <= 12 && offset > 0; l.month++ {
		if leapMonth > 0 && l.month == (leapMonth+1) && !l.isLeapMonth {
			l.month--
			l.isLeapMonth = true
			daysInMonth = l.getDaysInLeapMonth()
		} else {
			daysInMonth = l.getDaysInMonth()
		}
		offset -= daysInMonth
		if l.isLeapMonth && l.month == (leapMonth+1) {
			l.isLeapMonth = false
		}
	}
	// offset为0时，并且刚才计算的月份是闰月，要校正
	if offset == 0 && leapMonth > 0 && l.month == leapMonth+1 {
		if l.isLeapMonth {
			l.isLeapMonth = false
		} else {
			l.isLeapMonth = true
			l.month--
		}
	}
	// offset小于0时，也要校正
	if offset < 0 {
		offset += daysInMonth
		l.month--
	}
	l.day = offset + 1
	l.hour, l.minute, l.second = g.Clock()
	return
}

// ToGregorian converts Lunar instance to Gregorian instance.
// 将 Lunar 实例转化为 Gregorian 实例
func (l Lunar) ToGregorian() (g Gregorian) {
	if l.IsZero() {
		return
	}
	if l.year < minYear || l.year > maxYear {
		g.Error = invalidYearError()
		return
	}

	days := l.getDaysInMonth()
	offset := l.getOffsetInYear()
	offset += l.getOffsetInMonth()
	// 转换闰月农历 需补充该年闰月的前一个月的时差
	if l.isLeapMonth {
		offset += days
	}
	// https://github.com/golang-module/carbon/issues/219
	ts := int64(offset+l.day)*86400 - int64(2206512000) + int64(l.hour)*3600 + int64(l.minute)*60 + int64(l.second)
	g.Time = time.Unix(ts, 0)
	return
}

// Animal gets lunar animal name like "猴".
// 获取农历生肖
func (l Lunar) Animal() string {
	if l.IsZero() {
		return ""
	}
	return animals[l.year%calendar.MonthsPerYear]
}

// Festival gets lunar festival name like "春节".
// 获取农历节日
func (l Lunar) Festival() string {
	if l.IsZero() {
		return ""
	}
	return festivals[fmt.Sprintf("%d-%d", l.month, l.day)]
}

// Year gets lunar year like 2020.
// 获取农历年份
func (l Lunar) Year() int {
	if l.Error != nil {
		return 0
	}
	return l.year
}

// Month gets lunar month like 8.
// 获取农历月份
func (l Lunar) Month() int {
	if l.Error != nil {
		return 0
	}
	return l.month
}

// Day gets lunar day like 5.
// 获取农历日，如 5
func (l Lunar) Day() int {
	if l.Error != nil {
		return 0
	}
	return l.day
}

// Hour gets current hour like 13.
// 获取农历时辰
func (l Lunar) Hour() int {
	if l.Error != nil {
		return 0
	}
	return l.hour
}

// Minute gets current minute like 14.
// 获取农历分钟数
func (l Lunar) Minute() int {
	if l.Error != nil {
		return 0
	}
	return l.minute
}

// Second gets current second like 15.
// 获取农历秒数
func (l Lunar) Second() int {
	if l.Error != nil {
		return 0
	}
	return l.second
}

// LeapMonth gets lunar leap month like 2.
// 获取农历闰月月份，如 2
func (l Lunar) LeapMonth() int {
	if l.Error != nil {
		return 0
	}
	if l.year-minYear < 0 {
		return 0
	}
	return years[l.year-minYear] & 0xf
}

// String implements Stringer interface and outputs a string in YYYY-MM-DD HH::ii::ss format like "2019-12-07 00:00:00".
// 实现 Stringer 接口, 输出 YYYY-MM-DD HH::ii::ss 格式字符串，如 "2019-12-07 00:00:00"
func (l Lunar) String() string {
	if l.IsZero() {
		return ""
	}
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", l.year, l.month, l.day, l.hour, l.minute, l.second)
}

// ToYearString outputs a string in lunar year format like "二零二零".
// 获取农历年份字符串，如 "二零二零"
func (l Lunar) ToYearString() (year string) {
	if l.IsZero() {
		return
	}
	year = fmt.Sprintf("%d", l.year)
	for index, number := range numbers {
		year = strings.Replace(year, fmt.Sprintf("%d", index), number, -1)
	}
	return year
}

// ToMonthString outputs a string in lunar month format like "正月".
// 获取农历月份字符串，如 "正月"
func (l Lunar) ToMonthString() (month string) {
	if l.IsZero() {
		return
	}
	month = months[l.month-1] + "月"
	if l.IsLeapMonth() {
		return "闰" + month
	}
	return
}

// ToWeekString outputs a string in week layout like "周一".
// 输出完整农历星期字符串，如 "周一"
func (l Lunar) ToWeekString() (month string) {
	if l.IsZero() {
		return ""
	}
	return weeks[l.ToGregorian().Week()]
}

// ToDayString outputs a string in lunar day format like "廿一".
// 获取农历日字符串，如 "廿一"
func (l Lunar) ToDayString() (day string) {
	if l.IsZero() {
		return
	}
	num := numbers[l.day%10]
	switch {
	case l.day == 30:
		day = "三十"
	case l.day > 20:
		day = "廿" + num
	case l.day == 20:
		day = "二十"
	case l.day > 10:
		day = "十" + num
	case l.day == 10:
		day = "初十"
	case l.day < 10:
		day = "初" + num
	}
	return
}

// ToDateString outputs a string in lunar date format like "二零二零年腊月初五".
// 获取农历日期字符串，如 "二零二零年腊月初五"
func (l Lunar) ToDateString() string {
	if l.IsZero() {
		return ""
	}
	return l.ToYearString() + "年" + l.ToMonthString() + l.ToDayString()
}

// IsZero reports whether is zero time.
// 是否是农历零值时间
func (l Lunar) IsZero() bool {
	if l.Error != nil {
		return true
	}
	if l.year == 0 || l.month == 0 || l.day == 0 {
		return true
	}
	return false
}

// IsLeapYear reports whether is leap year.
// 是否是农历闰年
func (l Lunar) IsLeapYear() bool {
	if l.Error != nil {
		return false
	}
	return l.LeapMonth() != 0
}

// IsLeapMonth reports whether is leap month.
// 是否是农历闰月
func (l Lunar) IsLeapMonth() bool {
	if l.Error != nil {
		return false
	}
	return l.month == l.LeapMonth()
}

// IsRatYear reports whether is year of Rat.
// 是否是鼠年
func (l Lunar) IsRatYear() bool {
	if l.Error != nil {
		return false
	}
	if l.year%calendar.MonthsPerYear == 4 {
		return true
	}
	return false
}

// IsOxYear reports whether is year of Ox.
// 是否是牛年
func (l Lunar) IsOxYear() bool {
	if l.Error != nil {
		return false
	}
	if l.year%calendar.MonthsPerYear == 5 {
		return true
	}
	return false
}

// IsTigerYear reports whether is year of Tiger.
// 是否是虎年
func (l Lunar) IsTigerYear() bool {
	if l.Error != nil {
		return false
	}
	if l.year%calendar.MonthsPerYear == 6 {
		return true
	}
	return false
}

// IsRabbitYear reports whether is year of Rabbit.
// 是否是兔年
func (l Lunar) IsRabbitYear() bool {
	if l.Error != nil {
		return false
	}
	if l.year%calendar.MonthsPerYear == 7 {
		return true
	}
	return false
}

// IsDragonYear reports whether is year of Dragon.
// 是否是龙年
func (l Lunar) IsDragonYear() bool {
	if l.Error != nil {
		return false
	}
	if l.year%calendar.MonthsPerYear == 8 {
		return true
	}
	return false
}

// IsSnakeYear reports whether is year of Snake.
// 是否是蛇年
func (l Lunar) IsSnakeYear() bool {
	if l.Error != nil {
		return false
	}
	if l.year%calendar.MonthsPerYear == 9 {
		return true
	}
	return false
}

// IsHorseYear reports whether is year of Horse.
// 是否是马年
func (l Lunar) IsHorseYear() bool {
	if l.Error != nil {
		return false
	}
	if l.year%calendar.MonthsPerYear == 10 {
		return true
	}
	return false
}

// IsGoatYear reports whether is year of Goat.
// 是否是羊年
func (l Lunar) IsGoatYear() bool {
	if l.Error != nil {
		return false
	}
	if l.year%calendar.MonthsPerYear == 11 {
		return true
	}
	return false
}

// IsMonkeyYear reports whether is year of Monkey.
// 是否是猴年
func (l Lunar) IsMonkeyYear() bool {
	if l.Error != nil {
		return false
	}
	if l.year%calendar.MonthsPerYear == 0 {
		return true
	}
	return false
}

// IsRoosterYear reports whether is year of Rooster.
// 是否是鸡年
func (l Lunar) IsRoosterYear() bool {
	if l.Error != nil {
		return false
	}
	if l.year%calendar.MonthsPerYear == 1 {
		return true
	}
	return false
}

// IsDogYear reports whether is year of Dog.
// 是否是狗年
func (l Lunar) IsDogYear() bool {
	if l.Error != nil {
		return false
	}
	if l.year%calendar.MonthsPerYear == 2 {
		return true
	}
	return false
}

// IsPigYear reports whether is year of Pig.
// 是否是猪年
func (l Lunar) IsPigYear() bool {
	if l.Error != nil {
		return false
	}
	if l.year%calendar.MonthsPerYear == 3 {
		return true
	}
	return false
}

func (g Gregorian) diffInDays(t time.Time) int {
	return int(g.Time.Sub(t).Hours() / 24)
}

func (l Lunar) getOffsetInYear() int {
	flag := false
	clone, month, offset := l, 0, 0
	for month = 1; month < l.month; month++ {
		leapMonth := l.LeapMonth()
		if !flag {
			// 处理闰月
			if leapMonth <= month && leapMonth > 0 {
				offset += l.getDaysInLeapMonth()
				flag = true
			}
		}
		clone.month = month
		offset += clone.getDaysInMonth()
	}
	return offset
}

func (l Lunar) getOffsetInMonth() int {
	clone, year, offset := l, 0, 0
	for year = minYear; year < l.year; year++ {
		clone.year = year
		offset += clone.getDaysInYear()
	}
	return offset
}

func (l Lunar) getDaysInYear() int {
	var days = 348
	for i := 0x8000; i > 0x8; i >>= 1 {
		if (years[l.year-minYear] & i) != 0 {
			days++
		}
	}
	return days + l.getDaysInLeapMonth()
}

func (l Lunar) getDaysInMonth() int {
	if (years[l.year-minYear] & (0x10000 >> uint(l.month))) != 0 {
		return 30
	}
	return 29
}

func (l Lunar) getDaysInLeapMonth() int {
	if l.LeapMonth() == 0 {
		return 0
	}
	if years[l.year-minYear]&0x10000 != 0 {
		return 30
	}
	return 29
}
