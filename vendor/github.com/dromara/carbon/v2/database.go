package carbon

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

// timestamp precision constants
// 时间戳精度常量
const (
	PrecisionSecond      = 9
	PrecisionMillisecond = 999
	PrecisionMicrosecond = 999999
	PrecisionNanosecond  = 999999999
)

// returns a failed scan error
// 失败的扫描错误
var failedScanError = func(src interface{}) error {
	return fmt.Errorf("failed to scan value: %v", src)
}

// LayoutFactory defines a LayoutFactory interface
// 定义 LayoutFactory 接口
type LayoutFactory interface {
	~string
	SetLayout() string
}

// LayoutType defines a LayoutType generic struct
// 定义 LayoutType 泛型结构体
type LayoutType[T LayoutFactory] struct {
	*Carbon
}

// FormatFactory defines a FormatFactory interface.
// 定义 FormatFactory 接口
type FormatFactory interface {
	~string
	SetFormat() string
}

// FormatType defines a FormatType generic struct.
// 定义 FormatType 泛型结构体
type FormatType[T FormatFactory] struct {
	*Carbon
}

// TimestampFactory defines a TimestampFactory interface.
// 定义 TimestampFactory 接口
type TimestampFactory interface {
	~int64
	SetPrecision() int64
}

// TimestampType defines a TimestampType generic struct.
// 定义 TimestampType 泛型结构体
type TimestampType[T TimestampFactory] struct {
	*Carbon
}

// NewLayoutType returns a new LayoutType generic instance.
// 返回 LayoutType 泛型实例
func NewLayoutType[T LayoutFactory](carbon *Carbon) LayoutType[T] {
	return LayoutType[T]{
		Carbon: carbon,
	}
}

// NewFormatType returns a new FormatType generic instance.
// 返回 FormatType 泛型实例
func NewFormatType[T FormatFactory](carbon *Carbon) FormatType[T] {
	return FormatType[T]{
		Carbon: carbon,
	}
}

// NewTimestampType returns a new TimestampType generic instance.
// 返回 TimestampType 泛型实例
func NewTimestampType[T TimestampFactory](carbon *Carbon) TimestampType[T] {
	return TimestampType[T]{
		Carbon: carbon,
	}
}

// Scan implements driver.Scanner interface for LayoutType generic struct.
// 实现 driver.Scanner 接口
func (t *LayoutType[T]) Scan(src interface{}) error {
	c := NewCarbon()
	switch v := src.(type) {
	case []byte:
		c = Parse(string(v), DefaultTimezone)
	case string:
		c = Parse(v, DefaultTimezone)
	case time.Time:
		c = CreateFromStdTime(v, DefaultTimezone)
	case int64:
		c = CreateFromTimestamp(v, DefaultTimezone)
	default:
		return failedScanError(v)
	}
	*t = NewLayoutType[T](c)
	return t.Error
}

// Value implements driver.Valuer interface for LayoutType generic struct.
// 实现 driver.Valuer 接口
func (t LayoutType[T]) Value() (driver.Value, error) {
	if t.IsNil() || t.IsZero() {
		return nil, nil
	}
	if t.HasError() {
		return nil, t.Error
	}
	return t.StdTime(), nil
}

// MarshalJSON implements json.Marshal interface for LayoutType generic struct.
// 实现 json.Marshaler 接口
func (t LayoutType[T]) MarshalJSON() ([]byte, error) {
	emptyBytes := []byte(`""`)
	if t.IsNil() || t.IsZero() {
		return emptyBytes, nil
	}
	if t.HasError() {
		return emptyBytes, t.Error
	}
	return []byte(fmt.Sprintf(`"%s"`, t.Layout(t.getLayout(), t.Timezone()))), nil
}

// UnmarshalJSON implements json.Unmarshal interface for LayoutType generic struct.
// 实现 json.Unmarshaler 接口
func (t *LayoutType[T]) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" || value == "0" {
		t.Carbon = nil
		return nil
	}
	*t = NewLayoutType[T](ParseByLayout(value, t.getLayout()))
	return t.Error
}

// String implements Stringer interface for LayoutType generic struct.
// 实现 Stringer 接口
func (t LayoutType[T]) String() string {
	if t.IsZero() || t.IsInvalid() {
		return ""
	}
	return t.Layout(t.getLayout(), t.Timezone())
}

// GormDataType sets gorm data type for LayoutType generic struct.
// 设置 gorm 数据类型
func (t LayoutType[T]) GormDataType() string {
	return "time"
}

// getLayout returns the set layout.
// 返回设置的布局模板
func (t LayoutType[T]) getLayout() string {
	var factory T
	return factory.SetLayout()
}

// Scan implements driver.Scanner interface for FormatType generic struct.
// 实现 driver.Scanner 接口
func (t *FormatType[T]) Scan(src interface{}) error {
	c := NewCarbon()
	switch v := src.(type) {
	case []byte:
		c = Parse(string(v), DefaultTimezone)
	case string:
		c = Parse(v, DefaultTimezone)
	case time.Time:
		c = CreateFromStdTime(v, DefaultTimezone)
	case int64:
		c = CreateFromTimestamp(v, DefaultTimezone)
	default:
		return failedScanError(v)
	}
	*t = NewFormatType[T](c)
	return t.Error
}

// Value implements driver.Valuer interface for FormatType generic struct.
// 实现 driver.Valuer 接口
func (t FormatType[T]) Value() (driver.Value, error) {
	if t.IsNil() || t.IsZero() {
		return nil, nil
	}
	if t.HasError() {
		return nil, t.Error
	}
	return t.StdTime(), nil
}

// MarshalJSON implements json.Marshal interface for FormatType generic struct.
// 实现 json.Marshaler 接口
func (t FormatType[T]) MarshalJSON() ([]byte, error) {
	emptyBytes := []byte(`""`)
	if t.IsNil() || t.IsZero() {
		return emptyBytes, nil
	}
	if t.HasError() {
		return emptyBytes, t.Error
	}
	return []byte(fmt.Sprintf(`"%s"`, t.Format(t.getFormat(), t.Timezone()))), nil
}

// UnmarshalJSON implements json.Unmarshal interface for FormatType generic struct.
// 实现 json.Unmarshaler 接口
func (t *FormatType[T]) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" || value == "0" {
		t.Carbon = nil
		return nil
	}
	*t = NewFormatType[T](ParseByFormat(value, t.getFormat()))
	return t.Error
}

// String implements Stringer interface for FormatType generic struct.
// 实现 Stringer 接口
func (t FormatType[T]) String() string {
	if t.IsZero() || t.IsInvalid() {
		return ""
	}
	return t.Format(t.getFormat(), t.Timezone())
}

// GormDataType sets gorm data type for FormatType generic struct.
// 设置 gorm 数据类型
func (t FormatType[T]) GormDataType() string {
	return "time"
}

// getFormat returns the set format.
// 返回设置的格式模板
func (t FormatType[T]) getFormat() string {
	var factory T
	return factory.SetFormat()
}

// Scan implements driver.Scanner interface for TimestampType generic struct.
// 实现 driver.Scanner 接口
func (t *TimestampType[T]) Scan(src interface{}) (err error) {
	ts := int64(0)
	c := NewCarbon()
	switch v := src.(type) {
	case []byte:
		ts, err = strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return invalidTimestampError(string(v))
		}
	case string:
		ts, err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			return invalidTimestampError(v)
		}
	case int64:
		ts = v
	case time.Time:
		c = CreateFromStdTime(v, DefaultTimezone)
		*t = NewTimestampType[T](c)
		return t.Error
	default:
		return failedScanError(src)
	}
	switch t.getPrecision() {
	case PrecisionSecond:
		c = CreateFromTimestamp(ts, DefaultTimezone)
	case PrecisionMillisecond:
		c = CreateFromTimestampMilli(ts, DefaultTimezone)
	case PrecisionMicrosecond:
		c = CreateFromTimestampMicro(ts, DefaultTimezone)
	case PrecisionNanosecond:
		c = CreateFromTimestampNano(ts, DefaultTimezone)
	}
	*t = NewTimestampType[T](c)
	return t.Error
}

// Value implements driver.Valuer interface for TimestampType generic struct.
// 实现 driver.Valuer 接口
func (t TimestampType[T]) Value() (driver.Value, error) {
	if t.IsNil() || t.IsZero() {
		return nil, nil
	}
	if t.HasError() {
		return nil, t.Error
	}
	v := int64(0)
	switch t.getPrecision() {
	case PrecisionSecond:
		v = t.Timestamp()
	case PrecisionMillisecond:
		v = t.TimestampMilli()
	case PrecisionMicrosecond:
		v = t.TimestampMicro()
	case PrecisionNanosecond:
		v = t.TimestampNano()
	}
	return v, nil
}

// MarshalJSON implements json.Marshal interface for TimestampType generic struct.
// 实现 json.Marshaler 接口
func (t TimestampType[T]) MarshalJSON() ([]byte, error) {
	ts := int64(0)
	if t.IsNil() || t.IsZero() {
		return []byte(fmt.Sprintf(`%d`, ts)), nil
	}
	if t.HasError() {
		return []byte(fmt.Sprintf(`%d`, ts)), t.Error
	}
	switch t.getPrecision() {
	case PrecisionSecond:
		ts = t.Timestamp()
	case PrecisionMillisecond:
		ts = t.TimestampMilli()
	case PrecisionMicrosecond:
		ts = t.TimestampMicro()
	case PrecisionNanosecond:
		ts = t.TimestampNano()
	}
	return []byte(fmt.Sprintf(`%d`, ts)), nil
}

// UnmarshalJSON implements json.Unmarshal interface for TimestampType generic struct.
// 实现 json.Unmarshaler 接口
func (t *TimestampType[T]) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	c := NewCarbon()
	if value == "" || value == "null" || value == "0" {
		t.Carbon = nil
		return nil
	}
	ts, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return invalidTimestampError(value)
	}
	switch t.getPrecision() {
	case PrecisionSecond:
		c = CreateFromTimestamp(ts, DefaultTimezone)
	case PrecisionMillisecond:
		c = CreateFromTimestampMilli(ts, DefaultTimezone)
	case PrecisionMicrosecond:
		c = CreateFromTimestampMicro(ts, DefaultTimezone)
	case PrecisionNanosecond:
		c = CreateFromTimestampNano(ts, DefaultTimezone)
	}
	*t = NewTimestampType[T](c)
	return t.Error
}

// String implements Stringer interface for TimestampType generic struct.
// 实现 Stringer 接口
func (t TimestampType[T]) String() string {
	return strconv.FormatInt(t.Int64(), 10)
}

func (t TimestampType[T]) Int64() int64 {
	ts := int64(0)
	if t.IsZero() || t.IsInvalid() {
		return ts
	}
	switch t.getPrecision() {
	case PrecisionSecond:
		ts = t.Timestamp()
	case PrecisionMillisecond:
		ts = t.TimestampMilli()
	case PrecisionMicrosecond:
		ts = t.TimestampMicro()
	case PrecisionNanosecond:
		ts = t.TimestampNano()
	}
	return ts
}

// GormDataType sets gorm data type for TimestampType generic struct.
// 设置 gorm 数据类型
func (t TimestampType[T]) GormDataType() string {
	return "time"
}

// getPrecision returns the set timestamp precision.
// 返回设置的时间戳精度
func (t TimestampType[T]) getPrecision() int64 {
	var factory T
	return factory.SetPrecision()
}
