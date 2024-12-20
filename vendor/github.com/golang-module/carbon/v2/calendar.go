package carbon

import (
	"github.com/golang-module/carbon/v2/calendar/julian"
	"github.com/golang-module/carbon/v2/calendar/lunar"
	"github.com/golang-module/carbon/v2/calendar/persian"
)

// Lunar converts Carbon instance to Lunar instance.
// 将 Carbon 实例转化为 Lunar 实例
func (c Carbon) Lunar() (l lunar.Lunar) {
	if c.Error != nil {
		l.Error = c.Error
		return
	}
	return lunar.FromGregorian(c.StdTime()).ToLunar()
}

// CreateFromLunar creates a Carbon instance from Lunar date and time.
// 从 农历日期 创建 Carbon 实例
func CreateFromLunar(year, month, day, hour, minute, second int, isLeapMonth bool) Carbon {
	t := lunar.FromLunar(year, month, day, hour, minute, second, isLeapMonth).ToGregorian().Time
	return CreateFromStdTime(t)
}

// Julian converts Carbon instance to Julian instance.
// 将 Carbon 实例转化为 Julian 实例
func (c Carbon) Julian() (j julian.Julian) {
	if c.Error != nil {
		return
	}
	return julian.FromGregorian(c.StdTime()).ToJulian()
}

// CreateFromJulian creates a Carbon instance from Julian Day or Modified Julian Day.
// 从 儒略日/简化儒略日 创建 Carbon 实例
func CreateFromJulian(f float64) Carbon {
	t := julian.FromJulian(f).ToGregorian().Time
	return CreateFromStdTime(t)
}

// Persian converts Carbon instance to Persian instance.
// 将 Carbon 实例转化为 Persian 实例
func (c Carbon) Persian() (p persian.Persian) {
	if c.Error != nil {
		return
	}
	return persian.FromGregorian(c.StdTime()).ToPersian()
}

// CreateFromPersian creates a Carbon instance from Persian date and time.
// 从 波斯日期 创建 Carbon 实例
func CreateFromPersian(year, month, day, hour, minute, second int) Carbon {
	t := persian.FromPersian(year, month, day, hour, minute, second).ToGregorian().Time
	return CreateFromStdTime(t)
}
