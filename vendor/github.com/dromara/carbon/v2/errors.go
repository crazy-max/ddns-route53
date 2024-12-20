package carbon

import (
	"fmt"
)

// returns an invalid timezone error.
// 无效的时区错误
var invalidTimezoneError = func(timezone string) error {
	return fmt.Errorf("invalid timezone %q, please see the file %q for all valid timezones", timezone, "$GOROOT/lib/time/zoneinfo.zip")
}

// returns an invalid location error.
// 无效的地区错误
var invalidLocationError = func() error {
	return fmt.Errorf("invalid location, please make sure the location is valid")
}

// returns an invalid duration error.
// 无效的时长错误
var invalidDurationError = func(duration string) error {
	return fmt.Errorf("invalid duration %q, please make sure the duration is valid", duration)
}

// returns an invalid value error.
// 无效的时间字符串错误
var invalidValueError = func(value string) error {
	return fmt.Errorf("cannot parse string %q as carbon, please make sure the value is valid", value)
}

// returns an invalid layout error.
// 无效的布局模板错误
var invalidLayoutError = func(value, layout string) error {
	return fmt.Errorf("cannot parse string %q as carbon by layout %q, please make sure the value and layout match", value, layout)
}

// returns an invalid format error.
// 无效的格式模板错误
var invalidFormatError = func(value, format string) error {
	return fmt.Errorf("cannot parse string %q as carbon by format %q, please make sure the value and format match", value, format)
}
