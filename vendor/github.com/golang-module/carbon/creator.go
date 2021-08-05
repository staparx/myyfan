package carbon

import (
	"strconv"
	"time"
)

// CreateFromTimestamp create a Carbon instance from a given timestamp, second, millisecond, microsecond and nanosecond are supported
// 从给定的时间戳创建 Carbon 实例，支持秒、毫秒、微秒和纳秒
func (c Carbon) CreateFromTimestamp(timestamp int64, timezone ...string) Carbon {
	if len(timezone) > 0 {
		loc, err := getLocationByTimezone(timezone[len(timezone)-1])
		c.Loc = loc
		c.Error = err
	}
	if c.Error != nil {
		return c
	}
	ts, count := timestamp, len(strconv.FormatInt(timestamp, 10))
	if timestamp < 0 {
		count -= 1
	}
	switch count {
	case 10:
		ts = timestamp
	case 13:
		ts = timestamp / 1e3
	case 16:
		ts = timestamp / 1e6
	case 19:
		ts = timestamp / 1e9
	}
	c.Time = time.Unix(ts, 0)
	return c
}

// CreateFromTimestamp create a Carbon instance from a given timestamp
// 从给定的时间戳创建 Carbon 实例
func CreateFromTimestamp(timestamp int64, timezone ...string) Carbon {
	return NewCarbon().CreateFromTimestamp(timestamp, timezone...)
}

// CreateFromDateTime create a Carbon instance from a given date and time
// 从给定的年月日时分秒创建 Carbon 实例
func (c Carbon) CreateFromDateTime(year int, month int, day int, hour int, minute int, second int, timezone ...string) Carbon {
	if len(timezone) > 0 {
		loc, err := getLocationByTimezone(timezone[len(timezone)-1])
		c.Loc = loc
		c.Error = err
	}
	if c.Error != nil {
		return c
	}
	c.Time = time.Date(year, time.Month(month), day, hour, minute, second, time.Now().Nanosecond(), c.Loc)
	return c
}

// CreateFromDateTime create a Carbon instance from a given date and time
// 从给定的年月日时分秒创建 Carbon 实例
func CreateFromDateTime(year int, month int, day int, hour int, minute int, second int, timezone ...string) Carbon {
	return NewCarbon().CreateFromDateTime(year, month, day, hour, minute, second, timezone...)
}

// CreateFromDate create a Carbon instance from a given date
// 从给定的年月日创建 Carbon 实例
func (c Carbon) CreateFromDate(year int, month int, day int, timezone ...string) Carbon {
	if len(timezone) > 0 {
		loc, err := getLocationByTimezone(timezone[len(timezone)-1])
		c.Loc = loc
		c.Error = err
	}
	if c.Error != nil {
		return c
	}
	hour, minute, second := time.Now().In(c.Loc).Clock()
	c.Time = time.Date(year, time.Month(month), day, hour, minute, second, time.Now().Nanosecond(), c.Loc)
	return c
}

// CreateFromDate create a Carbon instance from a given date
// 从给定的年月日创建 Carbon 实例
func CreateFromDate(year int, month int, day int, timezone ...string) Carbon {
	return NewCarbon().CreateFromDate(year, month, day, timezone...)
}

// CreateFromTime create a Carbon instance from a given time
// 从给定的时分秒创建 Carbon 实例
func (c Carbon) CreateFromTime(hour int, minute int, second int, timezone ...string) Carbon {
	if len(timezone) > 0 {
		loc, err := getLocationByTimezone(timezone[len(timezone)-1])
		c.Loc = loc
		c.Error = err
	}
	if c.Error != nil {
		return c
	}
	year, month, day := time.Now().In(c.Loc).Date()
	c.Time = time.Date(year, month, day, hour, minute, second, time.Now().Nanosecond(), c.Loc)
	return c
}

// CreateFromTime create a Carbon instance from a given time
// 从给定的时分秒创建 Carbon 实例
func CreateFromTime(hour int, minute int, second int, timezone ...string) Carbon {
	return NewCarbon().CreateFromTime(hour, minute, second, timezone...)
}
