package carbon

import (
	"bytes"
	"fmt"
)

// MarshalJSON implements the json.Marshaler interface.
// 实现 json.Marshaler 接口
func (c Carbon) MarshalJSON() ([]byte, error) {
	if c.Error != nil {
		return nil, c.Error
	}
	key, value, tz := c.parseTag()
	data := ""
	if key == "layout" {
		data = fmt.Sprintf(`"%s"`, c.Layout(value, tz))
	}
	if key == "format" {
		// timestamp without double quotes in json
		if value == "U" || value == "V" || value == "X" || value == "Z" {
			data = fmt.Sprintf(`%s`, c.Format(value, tz))
		} else {
			data = fmt.Sprintf(`"%s"`, c.Format(value, tz))
		}
	}
	return []byte(data), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// 实现 json.Unmarshaler 接口
func (c *Carbon) UnmarshalJSON(b []byte) error {
	if c.Error != nil {
		return c.Error
	}
	if len(b) == 0 || string(b) == "null" {
		return nil
	}
	key, value, tz := c.parseTag()
	data := fmt.Sprintf("%s", bytes.Trim(b, `"`))
	if key == "layout" {
		*c = ParseByLayout(data, value, tz)
	}
	if key == "format" {
		*c = ParseByFormat(data, value, tz)
	}
	c.tag = &tag{
		carbon: fmt.Sprintf("%s:%s", key, value),
		tz:     tz,
	}
	return c.Error
}
