// @Package carbon
// @Description a simple, semantic and developer-friendly time package for golang
// @Page github.com/dromara/carbon
// @Developer gouguoyin
// @Blog www.gouguoyin.com
// @Email 245629560@qq.com

// Package carbon is a simple, semantic and developer-friendly time package for golang.
package carbon

import (
	"time"
)

// Carbon defines a Carbon struct.
// 定义 Carbon 结构体
type Carbon struct {
	time         time.Time
	layout       string
	weekStartsAt time.Weekday
	loc          *time.Location
	lang         *Language
	Error        error
}

// NewCarbon returns a new Carbon instance.
// 返回 Carbon 实例
func NewCarbon(time ...time.Time) *Carbon {
	c := &Carbon{lang: NewLanguage()}
	c.loc, c.Error = getLocationByTimezone(DefaultTimezone)
	if weekday, ok := weekdays[DefaultWeekStartsAt]; ok {
		c.weekStartsAt = weekday
	}
	if len(time) > 0 {
		c.time = time[0]
		c.loc = time[0].Location()
	}
	c.layout = DefaultLayout
	return c
}

// Copy returns a new copy of the current Carbon instance
// 复制 Carbon 实例
func (c *Carbon) Copy() *Carbon {
	newCarbon := NewCarbon()

	newCarbon.time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.loc)
	newCarbon.layout = c.layout
	newCarbon.weekStartsAt = c.weekStartsAt
	newCarbon.loc = c.loc
	newCarbon.Error = c.Error

	newCarbon.lang.dir = c.lang.dir
	newCarbon.lang.locale = c.lang.locale
	newCarbon.lang.Error = c.lang.Error

	return newCarbon
}
