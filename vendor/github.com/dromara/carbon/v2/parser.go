package carbon

import (
	"strconv"
	"time"
)

// Parse parses a standard time string as a Carbon instance.
// 将标准格式时间字符串解析成 Carbon 实例
func Parse(value string, timezone ...string) *Carbon {
	c := NewCarbon()
	if value == "" {
		return nil
	}
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[0])
	}
	if c.HasError() {
		return c
	}
	switch value {
	case "now":
		return Now(c.Timezone())
	case "yesterday":
		return Yesterday(c.Timezone())
	case "tomorrow":
		return Tomorrow(c.Timezone())
	}
	for _, layout := range layouts {
		if tt, err := time.ParseInLocation(layout, value, c.loc); err == nil {
			c.time = tt
			c.layout = layout
			return c
		}
	}
	c.Error = failedParseError(value)
	return c
}

// ParseByFormat parses a time string as a Carbon instance by format.
// 通过格式模板将时间字符串解析成 Carbon 实例
func ParseByFormat(value, format string, timezone ...string) *Carbon {
	c := NewCarbon()
	if value == "" {
		return nil
	}
	if format == "" {
		c.Error = emptyFormatError()
		return c
	}
	c = ParseByLayout(value, format2layout(format), timezone...)
	if c.HasError() {
		c.Error = invalidFormatError(value, format)
	}
	return c
}

// ParseByLayout parses a time string as a Carbon instance by layout.
// 通过布局模板将时间字符串解析成 Carbon 实例
func ParseByLayout(value, layout string, timezone ...string) *Carbon {
	c := NewCarbon()
	if value == "" {
		return nil
	}
	if layout == "" {
		c.Error = emptyLayoutError()
		return c
	}
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[0])
	}
	if c.HasError() {
		return c
	}
	if layout == "timestamp" {
		ts, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			c.Error = invalidTimestampError(value)
			return c
		}
		return CreateFromTimestamp(ts, c.Timezone())
	}
	if layout == "timestampMilli" {
		ts, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			c.Error = invalidTimestampError(value)
			return c
		}
		return CreateFromTimestampMilli(ts, c.Timezone())
	}
	if layout == "timestampMicro" {
		ts, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			c.Error = invalidTimestampError(value)
			return c
		}
		return CreateFromTimestampMicro(ts, c.Timezone())
	}
	if layout == "timestampNano" {
		ts, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			c.Error = invalidTimestampError(value)
			return c
		}
		return CreateFromTimestampNano(ts, c.Timezone())
	}
	tt, err := time.ParseInLocation(layout, value, c.loc)
	if err != nil {
		c.Error = invalidLayoutError(value, layout)
		return c
	}
	c.time = tt
	c.layout = layout
	return c
}
