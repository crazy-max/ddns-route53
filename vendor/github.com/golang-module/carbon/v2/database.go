package carbon

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// Scan an interface used by Scan in package database/sql for Scanning value from database to local golang variable.
func (c *Carbon) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		loc, err := getLocationByTimezone(defaultTimezone)
		if c.loc != nil {
			loc = c.loc
		}
		*c = CreateFromStdTime(value)
		c.loc, c.Error = loc, err
		return nil
	}
	return fmt.Errorf("can not convert %v to carbon", v)
}

// Value the interface providing the Value method for package database/sql/driver.
func (c Carbon) Value() (driver.Value, error) {
	if c.IsZero() {
		return nil, nil
	}
	return c.StdTime(), nil
}

// GormDataType implements the interface GormDataTypeInterface for Carbon struct.
// 实现 GormDataTypeInterface 接口
func (c Carbon) GormDataType() string {
	return "time"
}
