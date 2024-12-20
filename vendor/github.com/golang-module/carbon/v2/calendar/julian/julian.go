// Package julian is part of the carbon package.
package julian

import (
	"math"
	"strconv"
	"time"

	"github.com/golang-module/carbon/v2/calendar"
)

var (
	// julian day or modified julian day decimal precision
	// 儒略日或简化儒略日小数精度
	decimalPrecision = 6

	// difference between Julian Day and Modified Julian Day
	// 儒略日和简化儒略日之间的差值
	diffJdFromMjd = 2400000.5
)

// Gregorian defines a Gregorian struct.
// 定义 Gregorian 结构体
type Gregorian struct {
	calendar.Gregorian
}

// Julian defines a Julian struct.
// 定义 Julian 结构体
type Julian struct {
	jd, mjd float64
}

// FromGregorian creates a Gregorian instance from time.Time.
// 从标准 time.Time 创建 Gregorian 实例
func FromGregorian(t time.Time) (g Gregorian) {
	if t.IsZero() {
		return
	}
	g.Time = t
	return g
}

// FromJulian creates a Julian instance from julian day or modified julian day.
// 从 儒略日 或 简化儒略日 创建 Julian 实例
func FromJulian(f float64) (j Julian) {
	// get length of the integer part
	l := len(strconv.Itoa(int(math.Ceil(f))))
	switch l {
	// modified julian day
	case 5:
		j.mjd = f
		j.jd = f + diffJdFromMjd
	// julian day
	case 7:
		j.jd = f
		j.mjd = f - diffJdFromMjd
	default:
		j.jd = 0
		j.mjd = 0
	}
	return
}

// ToJulian converts Gregorian instance to Julian instance.
// 将 Gregorian 实例转化为 Julian 实例
func (g Gregorian) ToJulian() (j Julian) {
	if g.IsZero() {
		return
	}
	y := g.Year()
	m := g.Month()
	d := float64(g.Day()) + ((float64(g.Second())/60+float64(g.Minute()))/60+float64(g.Hour()))/24
	n := 0
	f := false
	if y*372+m*31+int(d) >= 588829 {
		f = true
	}
	if m <= 2 {
		m += 12
		y--
	}
	if f {
		n = y / 100
		n = 2 - n + n/4
	}
	jd := float64(int(365.25*(float64(y)+4716))) + float64(int(30.6001*(float64(m)+1))) + d + float64(n) - 1524.5
	return FromJulian(jd)
}

// ToGregorian converts Julian instance to Gregorian instance.
// 将 Julian 实例转化为 Gregorian 实例
func (j Julian) ToGregorian() (g Gregorian) {
	if j.IsZero() {
		return
	}
	d := int(j.jd + 0.5)
	f := j.jd + 0.5 - float64(d)
	if d >= 2299161 {
		c := int((float64(d) - 1867216.25) / 36524.25)
		d += 1 + c - c/4
	}
	d += 1524
	year := int((float64(d) - 122.1) / 365.25)
	d -= int(365.25 * float64(year))
	month := int(float64(d) / 30.601)
	d -= int(30.601 * float64(month))
	day := d
	if month > 13 {
		month -= 13
		year -= 4715
	} else {
		month -= 1
		year -= 4716
	}
	f *= 24
	hour := int(f)

	f -= float64(hour)
	f *= 60
	minute := int(f)

	f -= float64(minute)
	f *= 60
	second := int(math.Round(f))
	return FromGregorian(time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local))
}

// JD gets julian day like 2460332.5
// 获取儒略日
func (j Julian) JD(precision ...int) float64 {
	if len(precision) > 0 {
		decimalPrecision = precision[0]
	}
	return parseFloat64(j.jd, decimalPrecision)
}

// MJD gets modified julian day like 60332
// 获取简化儒略日
func (j Julian) MJD(precision ...int) float64 {
	if len(precision) > 0 {
		decimalPrecision = precision[0]
	}
	return parseFloat64(j.mjd, decimalPrecision)
}

// IsZero reports whether is zero time.
// 是否是零值时间
func (j Julian) IsZero() bool {
	if j.jd == 0 || j.mjd == 0 {
		return true
	}
	return false
}

// parseFloat64 round to n decimal places
// 四舍五入保留 n 位小数点
func parseFloat64(f float64, n int) float64 {
	p10 := math.Pow10(n)
	return math.Round(f*p10) / p10
}
