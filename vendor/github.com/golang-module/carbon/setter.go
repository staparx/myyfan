package carbon

import (
	"time"
)

// SetTimezone set timezone
// 设置时区
func (c Carbon) SetTimezone(name string) Carbon {
	loc, err := getLocationByTimezone(name)
	c.Loc = loc
	c.Error = err
	return c
}

// SetTimezone set timezone
// 设置时区
func SetTimezone(name string) Carbon {
	return NewCarbon().SetTimezone(name)
}

// SetLanguage set language
// 设置语言对象
func (c Carbon) SetLanguage(lang *Language) Carbon {
	c.Lang = lang
	return c
}

// SetLanguage set language
// 设置语言对象
func SetLanguage(lang *Language) Carbon {
	c := NewCarbon()
	err := lang.SetLocale(lang.locale)
	c.Lang = lang
	c.Error = err
	return c
}

// SetLocale set locale
// 设置语言区域
func (c Carbon) SetLocale(locale string) Carbon {
	if c.IsInvalid() {
		return c
	}
	c.Error = c.Lang.SetLocale(locale)
	return c
}

// SetLocale set locale
// 设置语言区域
func SetLocale(locale string) Carbon {
	c := NewCarbon()
	c.Error = c.Lang.SetLocale(locale)
	return c
}

// SetYear set year
// 设置年
func (c Carbon) SetYear(year int) Carbon {
	if c.IsInvalid() {
		return c
	}
	c.Time = time.Date(year, time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Loc)
	return c
}

// SetMonth set month
// 设置月
func (c Carbon) SetMonth(month int) Carbon {
	if c.IsInvalid() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(month), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Loc)
	return c
}

// SetDay set day
// 设置日
func (c Carbon) SetDay(day int) Carbon {
	if c.IsInvalid() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()), day, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Loc)
	return c
}

// SetHour set hour
// 设置时
func (c Carbon) SetHour(hour int) Carbon {
	if c.IsInvalid() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), hour, c.Minute(), c.Second(), c.Nanosecond(), c.Loc)
	return c
}

// SetMinute set minute
// 设置分
func (c Carbon) SetMinute(minute int) Carbon {
	if c.IsInvalid() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), minute, c.Second(), c.Nanosecond(), c.Loc)
	return c
}

// SetSecond set second
// 设置秒
func (c Carbon) SetSecond(second int) Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), second, c.Nanosecond(), c.Loc)
	return c
}

// SetMillisecond set millisecond
// 设置毫秒
func (c Carbon) SetMillisecond(millisecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), millisecond*1e6, c.Loc)
	return c
}

// SetMicrosecond set microsecond
// 设置微秒
func (c Carbon) SetMicrosecond(microsecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), microsecond*1e3, c.Loc)
	return c
}

// SetNanosecond set nanosecond
// 设置纳秒
func (c Carbon) SetNanosecond(nanosecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), nanosecond, c.Loc)
	return c
}
