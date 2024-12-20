package carbon

import (
	"bytes"
	"time"
)

// week days
// 工作日
var weekdays = map[string]time.Weekday{
	Monday:    time.Monday,
	Tuesday:   time.Tuesday,
	Wednesday: time.Wednesday,
	Thursday:  time.Thursday,
	Friday:    time.Friday,
	Saturday:  time.Saturday,
	Sunday:    time.Sunday,
}

// common formatting symbols
// 常规格式化符号
var formats = map[byte]string{
	'd': "02",      // Day:    Day of the month, 2 digits with leading zeros. Eg: 01 to 31.
	'D': "Mon",     // Day:    A textual representation of a day, three letters. Eg: Mon through Sun.
	'j': "2",       // Day:    Day of the month without leading zeros. Eg: 1 to 31.
	'l': "Monday",  // Day:    A full textual representation of the day of the week. Eg: Sunday through Saturday.
	'F': "January", // Month:  A full textual representation of a month, such as January or March. Eg: January through December.
	'm': "01",      // Month:  Numeric representation of a month, with leading zeros. Eg: 01 through 12.
	'M': "Jan",     // Month:  A short textual representation of a month, three letters. Eg: Jan through Dec.
	'n': "1",       // Month:  Numeric representation of a month, without leading zeros. Eg: 1 through 12.
	'Y': "2006",    // Year:   A full numeric representation of a year, 4 digits. Eg: 1999 or 2003.
	'y': "06",      // Year:   A two digit representation of a year. Eg: 99 or 03.
	'a': "pm",      // Time:   Lowercase morning or afternoon sign. Eg: am or pm.
	'A': "PM",      // Time:   Uppercase morning or afternoon sign. Eg: AM or PM.
	'g': "3",       // Time:   12-hour format of an hour without leading zeros. Eg: 1 through 12.
	'h': "03",      // Time:   12-hour format of an hour with leading zeros. Eg: 01 through 12.
	'H': "15",      // Time:   24-hour format of an hour with leading zeros. Eg: 00 through 23.
	'i': "04",      // Time:   Minutes with leading zeros. Eg: 00 to 59.
	's': "05",      // Time:   Seconds with leading zeros. Eg: 00 through 59.
	'O': "-0700",   // Zone:   Difference to Greenwich time (GMT) in hours. Eg: +0200.
	'P': "-07:00",  // Zone:   Difference to Greenwich time (GMT) with colon between hours and minutes. Eg: +02:00.
	'T': "MST",     // Zone:   Timezone abbreviation. Eg: UTC, EST, MDT ...

	'U': "timestamp",      // Timestamp with second. Eg: 1699677240.
	'V': "timestampMilli", // TimestampMilli with second. Eg: 1596604455666.
	'X': "timestampMicro", // TimestampMicro with second. Eg: 1596604455666666.
	'Z': "timestampNano",  // TimestampNano with second. Eg: 1596604455666666666.
}

// common layout symbols
// 常规布局模板符号
var layouts = []string{
	DayDateTimeLayout,
	DateTimeLayout, DateTimeNanoLayout, ShortDateTimeLayout, ShortDateTimeNanoLayout,
	DateLayout, DateNanoLayout, ShortDateLayout, ShortDateNanoLayout,
	ISO8601Layout, ISO8601NanoLayout,
	RFC822Layout, RFC822ZLayout, RFC850Layout, RFC1123Layout, RFC1123ZLayout, RFC3339Layout, RFC3339NanoLayout, RFC1036Layout, RFC7231Layout,
	KitchenLayout,
	CookieLayout,
	ANSICLayout,
	UnixDateLayout,
	RubyDateLayout,
	"2006",
	"2006-1", "2006-1-2", "2006-1-2 15", "2006-1-2 15:4", "2006-1-2 15:4:5", "2006-1-2 15:4:5.999999999",
	"2006.1", "2006.1.2", "2006.1.2 15", "2006.1.2 15:4", "2006.1.2 15:4:5", "2006.1.2 15:4:5.999999999",
	"2006/1", "2006/1/2", "2006/1/2 15", "2006/1/2 15:4", "2006/1/2 15:4:5", "2006/1/2 15:4:5.999999999",
	"2006-01-02 15:04:05PM MST", "2006-01-02 15:04:05.999999999PM MST", "2006-1-2 15:4:5PM MST", "2006-1-2 15:4:5.999999999PM MST",
	"2006-01-02 15:04:05 PM MST", "2006-01-02 15:04:05.999999999 PM MST", "2006-1-2 15:4:5 PM MST", "2006-1-2 15:4:5.999999999 PM MST",
	"1/2/2006", "1/2/2006 15", "1/2/2006 15:4", "1/2/2006 15:4:5", "1/2/2006 15:4:5.999999999",
	"2006-1-2 15:4:5 -0700 MST", "2006-1-2 15:4:5.999999999 -0700 MST", "2006-1-2 15:04:05 -0700 MST", "2006-1-2 15:04:05.999999999 -0700 MST",
	"2006-01-02T15:04:05", "2006-01-02T15:04:05.999999999", "2006-1-2T3:4:5", "2006-1-2T3:4:5.999999999",
	"2006-01-02T15:04:05Z07", "2006-01-02T15:04:05.999999999Z07", "2006-1-2T15:4:5Z07", "2006-1-2T15:4:5.999999999Z07",
	"2006-01-02T15:04:05Z07:00", "2006-01-02T15:04:05.999999999Z07:00", "2006-1-2T15:4:5Z07:00", "2006-1-2T15:4:5.999999999Z07:00",
	"2006-01-02T15:04:05-07:00", "2006-01-02T15:04:05.999999999-07:00", "2006-1-2T15:4:5-07:00", "2006-1-2T15:4:5.999999999-07:00",
	"2006-01-02T15:04:05-0700", "2006-01-02T15:04:05.999999999-0700", "2006-1-2T3:4:5-0700", "2006-1-2T3:4:5.999999999-0700",
	"20060102150405-07:00", "20060102150405.999999999-07:00",
	"20060102150405Z07", "20060102150405.999999999Z07",
	"20060102150405Z07:00", "20060102150405.999999999Z07:00",
}

// converts format to layout.
// format 转 layout
func format2layout(format string) string {
	buffer := bytes.NewBuffer(nil)
	for i := 0; i < len(format); i++ {
		if layout, ok := formats[format[i]]; ok {
			buffer.WriteString(layout)
		} else {
			switch format[i] {
			case '\\': // raw output, no parse
				buffer.WriteByte(format[i+1])
				i++
				continue
			default:
				buffer.WriteByte(format[i])
			}
		}
	}
	return buffer.String()
}

// gets a Location instance by a timezone string.
// 通过时区获取 Location 实例
func getLocationByTimezone(timezone string) (*time.Location, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		err = invalidTimezoneError(timezone)
	}
	return loc, err
}

// parses as a Duration instance by a duration string.
// 通过时长解析
func parseByDuration(duration string) (time.Duration, error) {
	td, err := time.ParseDuration(duration)
	if err != nil {
		err = invalidDurationError(duration)
	}
	return td, err
}

// gets absolute value.
// 获取绝对值
func getAbsValue(value int64) int64 {
	return (value ^ value>>31) - value>>31
}
