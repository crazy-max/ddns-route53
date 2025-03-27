package carbon

// Timestamp defines a Timestamp type.
// 定义 Timestamp 字段类型
type Timestamp int64

// SetPrecision implements TimestampFactory interface for Timestamp type.
// 实现 TimestampFactory 接口
func (t Timestamp) SetPrecision() int64 {
	return PrecisionSecond
}

// TimestampMilli defines a TimestampMilli type.
// 定义 TimestampMilli 字段类型
type TimestampMilli int64

// SetPrecision implements TimestampFactory interface for TimestampMilli type.
// 实现 TimestampFactory 接口
func (t TimestampMilli) SetPrecision() int64 {
	return PrecisionMillisecond
}

// TimestampMicro defines a TimestampMicro type.
// 定义 TimestampMicro 字段类型
type TimestampMicro int64

// SetPrecision implements TimestampFactory interface for TimestampMicro type.
// 实现 TimestampFactory 接口
func (t TimestampMicro) SetPrecision() int64 {
	return PrecisionMicrosecond
}

// TimestampNano defines a TimestampNano type.
// 定义 TimestampNano 字段类型
type TimestampNano int64

// SetPrecision implements TimestampFactory interface for TimestampNano type.
// 实现 TimestampFactory 接口
func (t TimestampNano) SetPrecision() int64 {
	return PrecisionNanosecond
}

// DateTime defines a DateTime type.
// 定义 DateTime 字段类型
type DateTime string

// SetFormat implements FormatFactory interface for DateTime type.
// 实现 FormatFactory 接口
func (t DateTime) SetFormat() string {
	return DateTimeFormat
}

// SetLayout implements LayoutFactory interface for DateTime type.
// 实现 LayoutFactory 接口
func (t DateTime) SetLayout() string {
	return DateTimeLayout
}

// Date defines a Date type.
// 定义 Date 字段类型
type Date string

// SetFormat implements FormatFactory interface for Date type.
// 实现 FormatFactory 接口
func (t Date) SetFormat() string {
	return DateFormat
}

// SetLayout implements LayoutFactory interface for Date type.
// 实现 LayoutFactory 接口
func (t Date) SetLayout() string {
	return DateLayout
}

// Time defines a Time struct.
// 定义 Time 字段类型
type Time string

// SetFormat implements FormatFactory interface for Time type.
// 实现 FormatFactory 接口
func (t Time) SetFormat() string {
	return TimeFormat
}

// SetLayout implements LayoutFactory interface for Time type.
// 实现 LayoutFactory 接口
func (t Time) SetLayout() string {
	return TimeLayout
}
