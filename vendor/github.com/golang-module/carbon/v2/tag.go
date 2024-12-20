package carbon

import (
	"fmt"
	"reflect"
	"strings"
)

var (
	// supported types
	// 支持的类型
	tagTypes = map[string]string{
		"dateTime":           "layout:" + DateTimeLayout,
		"dateTimeMilli":      "layout:" + DateTimeMilliLayout,
		"dateTimeMicro":      "layout:" + DateTimeMicroLayout,
		"dateTimeNano":       "layout:" + DateTimeNanoLayout,
		"shortDateTime":      "layout:" + ShortDateTimeLayout,
		"shortDateTimeMilli": "layout:" + ShortDateTimeMilliLayout,
		"shortDateTimeMicro": "layout:" + ShortDateTimeMicroLayout,
		"shortDateTimeNano":  "layout:" + ShortDateTimeNanoLayout,
		"dayDateTime":        "layout:" + DayDateTimeLayout,

		"date":           "layout:" + DateLayout,
		"dateMilli":      "layout:" + DateMilliLayout,
		"dateMicro":      "layout:" + DateMicroLayout,
		"dateNano":       "layout:" + DateNanoLayout,
		"shortDate":      "layout:" + ShortDateLayout,
		"shortDateMilli": "layout:" + ShortDateMilliLayout,
		"shortDateMicro": "layout:" + ShortDateMicroLayout,
		"shortDateNano":  "layout:" + ShortDateNanoLayout,

		"time":           "layout:" + TimeLayout,
		"timeMilli":      "layout:" + TimeMilliLayout,
		"timeMicro":      "layout:" + TimeMicroLayout,
		"timeNano":       "layout:" + TimeNanoLayout,
		"shortTime":      "layout:" + ShortTimeLayout,
		"shortTimeMilli": "layout:" + ShortTimeMilliLayout,
		"shortTimeMicro": "layout:" + ShortTimeMicroLayout,
		"shortTimeNano":  "layout:" + ShortTimeNanoLayout,

		"atom":     "layout:" + AtomLayout,
		"ansic":    "layout:" + ANSICLayout,
		"cookie":   "layout:" + CookieLayout,
		"kitchen":  "layout:" + KitchenLayout,
		"rss":      "layout:" + RssLayout,
		"rubyDate": "layout:" + RubyDateLayout,
		"unixDate": "layout:" + UnixDateLayout,

		"rfc1036":      "layout:" + RFC1036Layout,
		"rfc1123":      "layout:" + RFC1123Layout,
		"rfc1123Z":     "layout:" + RFC1123ZLayout,
		"rfc2822":      "layout:" + RFC2822Layout,
		"rfc3339":      "layout:" + RFC3339Layout,
		"rfc3339Milli": "layout:" + RFC3339MilliLayout,
		"rfc3339Micro": "layout:" + RFC3339MicroLayout,
		"rfc3339Nano":  "layout:" + RFC3339NanoLayout,
		"rfc7231":      "layout:" + RFC7231Layout,
		"rfc822":       "layout:" + RFC822Layout,
		"rfc822Z":      "layout:" + RFC822ZLayout,
		"rfc850":       "layout:" + RFC850Layout,

		"iso8601":      "layout:" + ISO8601Layout,
		"iso8601Milli": "layout:" + ISO8601MilliLayout,
		"iso8601Micro": "layout:" + ISO8601MicroLayout,
		"iso8601Nano":  "layout:" + ISO8601NanoLayout,

		"timestamp":      "format:U",
		"timestampMilli": "format:V",
		"timestampMicro": "format:X",
		"timestampNano":  "format:Z",
	}

	// invalid pointer error
	// 无效的指针错误
	invalidPtrError = func() error {
		return fmt.Errorf("invalid struct pointer, please make sure the struct is a pointer")
	}

	// invalid tag error
	// 无效的标签错误
	invalidTagError = func(field string) error {
		return fmt.Errorf("invalid carbon tag in %s field, please make sure the tag is valid", field)
	}
)

// tag defines a tag struct.
// 定义 tag 结构体
type tag struct {
	carbon string
	tz     string
}

// SetTag sets tag.
// 设置标签
func (c Carbon) SetTag(tag *tag) Carbon {
	if c.Error != nil {
		return c
	}
	c.tag = tag
	return c
}

// parseTag parses tag.
// 解析标签
func (c Carbon) parseTag() (key, value, tz string) {
	if c.tag == nil {
		return "layout", defaultLayout, defaultTimezone
	}
	tz = strings.TrimSpace(c.tag.tz)
	if tz == "" {
		tz = defaultTimezone
	}
	carbon := strings.TrimSpace(c.tag.carbon)
	if carbon == "" {
		return "layout", defaultLayout, tz
	}
	if !strings.HasPrefix(carbon, "layout:") && !strings.HasPrefix(carbon, "format:") {
		return "", "", tz
	}
	key = strings.TrimSpace(carbon[:6])
	value = strings.TrimSpace(carbon[7:])
	return
}

// LoadTag loads tag.
// 加载标签
func LoadTag(v interface{}) error {
	typeObj, valueObj := reflect.TypeOf(v), reflect.ValueOf(v)
	if typeObj.Kind() != reflect.Ptr {
		return invalidPtrError()
	}
	typeElem, valueElem := typeObj.Elem(), valueObj.Elem()
	params := make([]reflect.Value, 1)
	for i := 0; i < valueElem.NumField(); i++ {
		fieldType, fieldValue := typeElem.Field(i), valueElem.Field(i)
		if reflect.TypeOf(Carbon{}) != fieldValue.Type() {
			continue
		}
		carbon := fieldType.Tag.Get("carbon")
		if carbon == "" {
			carbon = "layout:" + defaultLayout
		}
		if strings.HasPrefix(carbon, "type:") {
			carbon = tagTypes[carbon[5:]]
		}
		if !strings.HasPrefix(carbon, "layout:") && !strings.HasPrefix(carbon, "format:") {
			return invalidTagError(fieldType.Name)
		}
		tz := fieldType.Tag.Get("tz")
		if tz == "" {
			tz = defaultTimezone
		}
		params[0] = reflect.ValueOf(&tag{
			carbon: carbon,
			tz:     tz,
		})
		fieldValue.Set(fieldValue.MethodByName("SetTag").Call(params)[0])
	}
	return nil
}
