// @Package carbon
// @Description a simple, semantic and developer-friendly golang package for datetime
// @Page github.com/dromara/carbon
// @Developer gouguoyin
// @Blog www.gouguoyin.com
// @Email 245629560@qq.com

// Package carbon is a simple, semantic and developer-friendly golang package for datetime.
package carbon

import (
	"time"
)

// Version current version
// 当前版本号
const Version = "2.5.2"

// timezone constants
// 时区常量
const (
	Local = "Local" // 本地时间
	UTC   = "UTC"   // 世界协调时间
	GMT   = "GMT"   // 格林尼治标准时间
	CST   = "CST"   // 中国标准时间
	EET   = "EET"   // 欧洲东部标准时间
	WET   = "WET"   // 欧洲西部标准时间
	CET   = "CET"   // 欧洲中部标准时间
	EST   = "EST"   // 美国东部标准时间
	MST   = "MST"   // 美国山地标准时间

	Cuba      = "Cuba"      // 古巴
	Egypt     = "Egypt"     // 埃及
	Eire      = "Eire"      // 爱尔兰
	Greenwich = "Greenwich" // 格林尼治
	Iceland   = "Iceland"   // 冰岛
	Iran      = "Iran"      // 伊朗
	Israel    = "Israel"    // 以色列
	Jamaica   = "Jamaica"   // 牙买加
	Japan     = "Japan"     // 日本
	Libya     = "Libya"     // 利比亚
	Poland    = "Poland"    // 波兰
	Portugal  = "Portugal"  // 葡萄牙
	PRC       = "PRC"       // 中国
	Singapore = "Singapore" // 新加坡
	Turkey    = "Turkey"    // 土耳其

	Shanghai   = "Asia/Shanghai"       // 上海
	Chongqing  = "Asia/Chongqing"      // 重庆
	Harbin     = "Asia/Harbin"         // 哈尔滨
	Urumqi     = "Asia/Urumqi"         // 乌鲁木齐
	HongKong   = "Asia/Hong_Kong"      // 香港
	Macao      = "Asia/Macao"          // 澳门
	Taipei     = "Asia/Taipei"         // 台北
	Tokyo      = "Asia/Tokyo"          // 东京
	HoChiMinh  = "Asia/Ho_Chi_Minh"    // 胡志明
	Hanoi      = "Asia/Hanoi"          // 河内
	Saigon     = "Asia/Saigon"         // 西贡
	Seoul      = "Asia/Seoul"          // 首尔
	Pyongyang  = "Asia/Pyongyang"      // 平壤
	Bangkok    = "Asia/Bangkok"        // 曼谷
	Dubai      = "Asia/Dubai"          // 迪拜
	Qatar      = "Asia/Qatar"          // 卡塔尔
	Bangalore  = "Asia/Bangalore"      // 班加罗尔
	Kolkata    = "Asia/Kolkata"        // 加尔各答
	Mumbai     = "Asia/Mumbai"         // 孟买
	MexicoCity = "America/Mexico_City" // 墨西哥
	NewYork    = "America/New_York"    // 纽约
	LosAngeles = "America/Los_Angeles" // 洛杉矶
	Chicago    = "America/Chicago"     // 芝加哥
	SaoPaulo   = "America/Sao_Paulo"   // 圣保罗
	Moscow     = "Europe/Moscow"       // 莫斯科
	London     = "Europe/London"       // 伦敦
	Berlin     = "Europe/Berlin"       // 柏林
	Paris      = "Europe/Paris"        // 巴黎
	Rome       = "Europe/Rome"         // 罗马
	Sydney     = "Australia/Sydney"    // 悉尼
	Melbourne  = "Australia/Melbourne" // 墨尔本
	Darwin     = "Australia/Darwin"    // 达尔文
)

// month constants
// 月份常量
const (
	January   = "January"   // 一月
	February  = "February"  // 二月
	March     = "March"     // 三月
	April     = "April"     // 四月
	May       = "May"       // 五月
	June      = "June"      // 六月
	July      = "July"      // 七月
	August    = "August"    // 八月
	September = "September" // 九月
	October   = "October"   // 十月
	November  = "November"  // 十一月
	December  = "December"  // 十二月
)

// week constants
// 星期常量
const (
	Monday    = "Monday"    // 周一
	Tuesday   = "Tuesday"   // 周二
	Wednesday = "Wednesday" // 周三
	Thursday  = "Thursday"  // 周四
	Friday    = "Friday"    // 周五
	Saturday  = "Saturday"  // 周六
	Sunday    = "Sunday"    // 周日
)

// number constants
// 数字常量
const (
	YearsPerMillennium = 1000   // 每千年1000年
	YearsPerCentury    = 100    // 每世纪100年
	YearsPerDecade     = 10     // 每十年10年
	QuartersPerYear    = 4      // 每年4个季度
	MonthsPerYear      = 12     // 每年12月
	MonthsPerQuarter   = 3      // 每季度3月
	WeeksPerNormalYear = 52     // 每常规年52周
	weeksPerLongYear   = 53     // 每长年53周
	WeeksPerMonth      = 4      // 每月4周
	DaysPerLeapYear    = 366    // 每闰年366天
	DaysPerNormalYear  = 365    // 每常规年365天
	DaysPerWeek        = 7      // 每周7天
	HoursPerWeek       = 168    // 每周168小时
	HoursPerDay        = 24     // 每天24小时
	MinutesPerDay      = 1440   // 每天1440分钟
	MinutesPerHour     = 60     // 每小时60分钟
	SecondsPerWeek     = 604800 // 每周604800秒
	SecondsPerDay      = 86400  // 每天86400秒
	SecondsPerHour     = 3600   // 每小时3600秒
	SecondsPerMinute   = 60     // 每分钟60秒
)

// layout constants
// 布局模板常量
const (
	AtomLayout     = RFC3339Layout
	ANSICLayout    = time.ANSIC
	CookieLayout   = "Monday, 02-Jan-2006 15:04:05 MST"
	KitchenLayout  = time.Kitchen
	RssLayout      = time.RFC1123Z
	RubyDateLayout = time.RubyDate
	UnixDateLayout = time.UnixDate
	W3cLayout      = RFC3339Layout

	RFC1036Layout      = "Mon, 02 Jan 06 15:04:05 -0700"
	RFC1123Layout      = time.RFC1123
	RFC1123ZLayout     = time.RFC1123Z
	RFC2822Layout      = time.RFC1123Z
	RFC3339Layout      = "2006-01-02T15:04:05Z07:00"
	RFC3339MilliLayout = "2006-01-02T15:04:05.999Z07:00"
	RFC3339MicroLayout = "2006-01-02T15:04:05.999999Z07:00"
	RFC3339NanoLayout  = "2006-01-02T15:04:05.999999999Z07:00"
	RFC7231Layout      = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC822Layout       = time.RFC822
	RFC822ZLayout      = time.RFC822Z
	RFC850Layout       = time.RFC850

	ISO8601Layout      = "2006-01-02T15:04:05-07:00"
	ISO8601MilliLayout = "2006-01-02T15:04:05.999-07:00"
	ISO8601MicroLayout = "2006-01-02T15:04:05.999999-07:00"
	ISO8601NanoLayout  = "2006-01-02T15:04:05.999999999-07:00"

	ISO8601ZuluLayout      = "2006-01-02T15:04:05Z"
	ISO8601ZuluMilliLayout = "2006-01-02T15:04:05.999Z"
	ISO8601ZuluMicroLayout = "2006-01-02T15:04:05.999999Z"
	ISO8601ZuluNanoLayout  = "2006-01-02T15:04:05.999999999Z"

	FormattedDateLayout    = "Jan 2, 2006"
	FormattedDayDateLayout = "Mon, Jan 2, 2006"

	DayDateTimeLayout        = "Mon, Jan 2, 2006 3:04 PM"
	DateTimeLayout           = "2006-01-02 15:04:05"
	DateTimeMilliLayout      = "2006-01-02 15:04:05.999"
	DateTimeMicroLayout      = "2006-01-02 15:04:05.999999"
	DateTimeNanoLayout       = "2006-01-02 15:04:05.999999999"
	ShortDateTimeLayout      = "20060102150405"
	ShortDateTimeMilliLayout = "20060102150405.999"
	ShortDateTimeMicroLayout = "20060102150405.999999"
	ShortDateTimeNanoLayout  = "20060102150405.999999999"

	DateLayout           = "2006-01-02"
	DateMilliLayout      = "2006-01-02.999"
	DateMicroLayout      = "2006-01-02.999999"
	DateNanoLayout       = "2006-01-02.999999999"
	ShortDateLayout      = "20060102"
	ShortDateMilliLayout = "20060102.999"
	ShortDateMicroLayout = "20060102.999999"
	ShortDateNanoLayout  = "20060102.999999999"

	TimeLayout           = "15:04:05"
	TimeMilliLayout      = "15:04:05.999"
	TimeMicroLayout      = "15:04:05.999999"
	TimeNanoLayout       = "15:04:05.999999999"
	ShortTimeLayout      = "150405"
	ShortTimeMilliLayout = "150405.999"
	ShortTimeMicroLayout = "150405.999999"
	ShortTimeNanoLayout  = "150405.999999999"
)

// format constants
// 格式模板常量
const (
	AtomFormat     = "Y-m-d\\TH:i:sP"
	ANSICFormat    = "D M j H:i:s Y"
	CookieFormat   = "l, d-M-Y H:i:s T"
	KitchenFormat  = "g:iA"
	RssFormat      = "D, d M Y H:i:s O"
	RubyDateFormat = "D M d H:i:s O Y"
	UnixDateFormat = "D M j H:i:s T Y"

	RFC1036Format      = "D, d M y H:i:s O"
	RFC1123Format      = "D, d M Y H:i:s T"
	RFC1123ZFormat     = "D, d M Y H:i:s O"
	RFC2822Format      = "D, d M Y H:i:s O"
	RFC3339Format      = "Y-m-d\\TH:i:sP"
	RFC3339MilliFormat = "Y-m-d\\TH:i:s.vP"
	RFC3339MicroFormat = "Y-m-d\\TH:i:s.uP"
	RFC3339NanoFormat  = "Y-m-d\\TH:i:s.xP"
	RFC7231Format      = "D, d M Y H:i:s T"
	RFC822Format       = "d M y H:i T"
	RFC822ZFormat      = "d M y H:i O"
	RFC850Format       = "l, d-M-y H:i:s T"

	ISO8601Format      = "Y-m-d\\TH:i:sP"
	ISO8601MilliFormat = "Y-m-d\\TH:i:s.vP"
	ISO8601MicroFormat = "Y-m-d\\TH:i:s.uP"
	ISO8601NanoFormat  = "Y-m-d\\TH:i:s.xP"

	ISO8601ZuluFormat      = "Y-m-d\\TH:i:s\\Z"
	ISO8601ZuluMilliFormat = "Y-m-d\\TH:i:s.v\\Z"
	ISO8601ZuluMicroFormat = "Y-m-d\\TH:i:s.u\\Z"
	ISO8601ZuluNanoFormat  = "Y-m-d\\TH:i:s.x\\Z"

	FormattedDateFormat    = "M j, Y"
	FormattedDayDateFormat = "D, M j, Y"

	DayDateTimeFormat        = "D, M j, Y g:i A"
	DateTimeFormat           = "Y-m-d H:i:s"
	DateTimeMilliFormat      = "Y-m-d H:i:s.v"
	DateTimeMicroFormat      = "Y-m-d H:i:s.u"
	DateTimeNanoFormat       = "Y-m-d H:i:s.x"
	ShortDateTimeFormat      = "YmdHis"
	ShortDateTimeMilliFormat = "YmdHis.v"
	ShortDateTimeMicroFormat = "YmdHis.u"
	ShortDateTimeNanoFormat  = "YmdHis.x"

	DateFormat           = "Y-m-d"
	DateMilliFormat      = "Y-m-d.v"
	DateMicroFormat      = "Y-m-d.u"
	DateNanoFormat       = "Y-m-d.x"
	ShortDateFormat      = "Ymd"
	ShortDateMilliFormat = "Ymd.v"
	ShortDateMicroFormat = "Ymd.u"
	ShortDateNanoFormat  = "Ymd.x"

	TimeFormat           = "H:i:s"
	TimeMilliFormat      = "H:i:s.v"
	TimeMicroFormat      = "H:i:s.u"
	TimeNanoFormat       = "H:i:s.x"
	ShortTimeFormat      = "His"
	ShortTimeMilliFormat = "His.v"
	ShortTimeMicroFormat = "His.u"
	ShortTimeNanoFormat  = "His.x"
)

// Carbon defines a Carbon struct.
// 定义 Carbon 结构体
type Carbon struct {
	time         time.Time
	testNow      int64 // nanosecond timestamp of test now time
	layout       string
	weekStartsAt time.Weekday
	loc          *time.Location
	lang         *Language
	Error        error
}

// NewCarbon returns a new Carbon instance.
// 初始化 Carbon 结构体
func NewCarbon() Carbon {
	c := Carbon{lang: NewLanguage()}
	c.loc, c.Error = getLocationByTimezone(defaultTimezone)
	if weekday, ok := weekdays[defaultWeekStartsAt]; ok {
		c.weekStartsAt = weekday
	}
	c.layout = defaultLayout
	return c
}

// DateTime defines a DateTime struct.
// 定义 DateTime 结构体
type DateTime struct {
	Carbon
}

// NewDateTime returns a new DateTime instance.
// 初始化 DateTime 结构体
func NewDateTime(carbon Carbon) DateTime {
	return DateTime{Carbon: carbon}
}

// DateTimeMilli defines a DateTimeMilli struct.
// 定义 DateTimeMilli 结构体
type DateTimeMilli struct {
	Carbon
}

// NewDateTimeMilli returns a new DateTimeMilli instance.
// 初始化 DateTimeMilli 结构体
func NewDateTimeMilli(carbon Carbon) DateTimeMilli {
	return DateTimeMilli{Carbon: carbon}
}

// DateTimeMicro defines a DateTimeMicro struct.
// 定义 DateTimeMicro 结构体
type DateTimeMicro struct {
	Carbon
}

// NewDateTimeMicro returns a new DateTimeMicro instance.
// 初始化 DateTimeMicro 结构体
func NewDateTimeMicro(carbon Carbon) DateTimeMicro {
	return DateTimeMicro{Carbon: carbon}
}

// DateTimeNano defines a DateTimeNano struct.
// 定义 DateTimeNano 结构体
type DateTimeNano struct {
	Carbon
}

// NewDateTimeNano returns a new DateTimeNano instance.
// 初始化 DateTimeNano 结构体
func NewDateTimeNano(carbon Carbon) DateTimeNano {
	return DateTimeNano{Carbon: carbon}
}

// Date defines a Date struct.
// 定义 Date 结构体
type Date struct {
	Carbon
}

// NewDate returns a new Date instance.
// 初始化 Date 结构体
func NewDate(carbon Carbon) Date {
	return Date{Carbon: carbon}
}

// DateMilli defines a DateMilli struct.
// 定义 DateMilli 结构体
type DateMilli struct {
	Carbon
}

// NewDateMilli returns a new DateMilli instance.
// 初始化 DateMilli 结构体
func NewDateMilli(carbon Carbon) DateMilli {
	return DateMilli{Carbon: carbon}
}

// DateMicro defines a DateMicro struct.
// 定义 DateMicro 结构体
type DateMicro struct {
	Carbon
}

// NewDateMicro returns a new DateMicro instance.
// 初始化 DateMicro 结构体
func NewDateMicro(carbon Carbon) DateMicro {
	return DateMicro{Carbon: carbon}
}

// DateNano defines a DateNano struct.
// 定义 DateNano 结构体
type DateNano struct {
	Carbon
}

// NewDateNano returns a new DateNano instance.
// 初始化 DateNano 结构体
func NewDateNano(carbon Carbon) DateNano {
	return DateNano{Carbon: carbon}
}

// Time defines a Time struct.
// 定义 Time 结构体
type Time struct {
	Carbon
}

// NewTime returns a new Time instance.
// 初始化 Time 结构体
func NewTime(carbon Carbon) Time {
	return Time{Carbon: carbon}
}

// TimeMilli defines a TimeMilli struct.
// 定义 TimeMilli 结构体
type TimeMilli struct {
	Carbon
}

// NewTimeMilli returns a new TimeMilli instance.
// 初始化 TimeMilli 结构体
func NewTimeMilli(carbon Carbon) TimeMilli {
	return TimeMilli{Carbon: carbon}
}

// TimeMicro defines a TimeMicro struct.
// 定义 TimeMicro 结构体
type TimeMicro struct {
	Carbon
}

// NewTimeMicro returns a new TimeMicro instance.
// 初始化 TimeMicro 结构体
func NewTimeMicro(carbon Carbon) TimeMicro {
	return TimeMicro{Carbon: carbon}
}

// TimeNano defines a TimeNano struct.
// 定义 TimeNano 结构体
type TimeNano struct {
	Carbon
}

// NewTimeNano returns a new TimeNano instance.
// 初始化 TimeNano 结构体
func NewTimeNano(carbon Carbon) TimeNano {
	return TimeNano{Carbon: carbon}
}

// Timestamp defines a Timestamp struct.
// 定义 Timestamp 结构体
type Timestamp struct {
	Carbon
}

// NewTimestamp returns a new Timestamp instance.
// 初始化 Timestamp 结构体
func NewTimestamp(carbon Carbon) Timestamp {
	return Timestamp{Carbon: carbon}
}

// TimestampMilli defines a TimestampMilli struct.
// 定义 TimestampMilli 结构体
type TimestampMilli struct {
	Carbon
}

// NewTimestampMilli returns a new TimestampMilli instance.
// 初始化 TimestampMilli 结构体
func NewTimestampMilli(carbon Carbon) TimestampMilli {
	return TimestampMilli{Carbon: carbon}
}

// TimestampMicro defines a TimestampMicro struct.
// 定义 TimestampMicro 结构体
type TimestampMicro struct {
	Carbon
}

// NewTimestampMicro returns a new TimestampMicro instance.
// 初始化 TimestampMicro 结构体
func NewTimestampMicro(carbon Carbon) TimestampMicro {
	return TimestampMicro{Carbon: carbon}
}

// TimestampNano defines a TimestampNano struct.
// 定义 TimestampNano 结构体
type TimestampNano struct {
	Carbon
}

// NewTimestampNano returns a new TimestampNano instance.
// 初始化 TimestampNano 结构体
func NewTimestampNano(carbon Carbon) TimestampNano {
	return TimestampNano{Carbon: carbon}
}
