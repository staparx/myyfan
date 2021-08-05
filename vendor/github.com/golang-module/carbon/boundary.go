package carbon

import (
	"time"
)

// StartOfCentury return a Carbon instance for start of the century
// 本世纪开始时间
func (c Carbon) StartOfCentury() Carbon {
	if c.IsZero() {
		return c
	}
	c.Time = time.Date(c.Year()/YearsPerCentury*YearsPerCentury, 1, 1, 0, 0, 0, 0, c.Loc)
	return c
}

// EndOfCentury return a Carbon instance for end of the century
// 本世纪结束时间
func (c Carbon) EndOfCentury() Carbon {
	if c.IsZero() {
		return c
	}
	c.Time = time.Date(c.Year()/YearsPerCentury*YearsPerCentury+999, 12, 31, 23, 59, 59, 999999999, c.Loc)
	return c
}

// StartOfDecade return a Carbon instance for start of the decade
// 本年代开始时间
func (c Carbon) StartOfDecade() Carbon {
	if c.IsZero() {
		return c
	}
	c.Time = time.Date(c.Year()/YearsPerDecade*YearsPerDecade, 1, 1, 0, 0, 0, 0, c.Loc)
	return c
}

// EndOfDecade return a Carbon instance for end of the decade
// 本年代结束时间
func (c Carbon) EndOfDecade() Carbon {
	if c.IsZero() {
		return c
	}
	c.Time = time.Date(c.Year()/YearsPerDecade*YearsPerDecade+9, 12, 31, 23, 59, 59, 999999999, c.Loc)
	return c
}

// StartOfYear return a Carbon instance for start of the year
// 本年开始时间
func (c Carbon) StartOfYear() Carbon {
	if c.IsZero() {
		return c
	}
	c.Time = time.Date(c.Year(), 1, 1, 0, 0, 0, 0, c.Loc)
	return c
}

// EndOfYear return a Carbon instance for end of the year
// 本年结束时间
func (c Carbon) EndOfYear() Carbon {
	if c.IsZero() {
		return c
	}
	c.Time = time.Date(c.Year(), 12, 31, 23, 59, 59, 999999999, c.Loc)
	return c
}

// StartOfQuarter return a Carbon instance for start of the quarter
// 本季度开始时间
func (c Carbon) StartOfQuarter() Carbon {
	if c.IsZero() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(3*c.Quarter()-2), 1, 0, 0, 0, 0, c.Loc)
	return c
}

// EndOfQuarter return a Carbon instance for end of the quarter
// 本季度结束时间
func (c Carbon) EndOfQuarter() Carbon {
	if c.IsZero() {
		return c
	}
	quarter, day := c.Quarter(), 30
	switch quarter {
	case 1, 4:
		day = 31
	case 2, 3:
		day = 30
	}
	c.Time = time.Date(c.Year(), time.Month(3*quarter), day, 23, 59, 59, 999999999, c.Loc)
	return c
}

// StartOfMonth return a Carbon instance for start of the month
// 本月开始时间
func (c Carbon) StartOfMonth() Carbon {
	if c.IsZero() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()), 1, 0, 0, 0, 0, c.Loc)
	return c
}

// EndOfMonth return a Carbon instance for end of the month
// 本月结束时间
func (c Carbon) EndOfMonth() Carbon {
	if c.IsZero() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()), 1, 23, 59, 59, 999999999, c.Loc).AddDate(0, 1, -1)
	return c
}

// StartOfWeek return a Carbon instance for start of the week
// 本周开始时间
func (c Carbon) StartOfWeek(weekStartDay time.Weekday) Carbon {
	if c.IsZero() {
		return c
	}
	weekDay := c.Time.In(c.Loc).Weekday()
	if weekDay == weekStartDay {
		return c.StartOfDay()
	}
	days := int(weekDay) - int(weekStartDay)
	if weekDay == time.Sunday {
		days = 6
	}
	return c.SubDays(days).StartOfDay()
}

// EndOfWeek return a Carbon instance for end of the week
// 本周结束时间
func (c Carbon) EndOfWeek(weekStartDay time.Weekday) Carbon {
	if c.IsZero() {
		return c
	}
	weekDay := c.Time.In(c.Loc).Weekday()
	if weekStartDay == 0 && weekDay == 6 {
		return c.EndOfDay()
	}
	if weekStartDay == 1 && weekDay == 0 {
		return c.EndOfDay()
	}
	days := 6 - int(weekDay) + int(weekStartDay)
	if weekDay == time.Sunday {
		days = 6
	}
	return c.AddDays(days).EndOfDay()
}

// StartOfDay return a Carbon instance for start of the day
// 本日开始时间
func (c Carbon) StartOfDay() Carbon {
	if c.IsZero() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), 0, 0, 0, 0, c.Loc)
	return c
}

// EndOfDay return a Carbon instance for end of the day
// 本日结束时间
func (c Carbon) EndOfDay() Carbon {
	if c.IsZero() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), 23, 59, 59, 999999999, c.Loc)
	return c
}

// StartOfHour return a Carbon instance for start of the hour
// 小时开始时间
func (c Carbon) StartOfHour() Carbon {
	if c.IsZero() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), 0, 0, 0, c.Loc)
	return c
}

// EndOfHour return a Carbon instance for end of the hour
// 小时结束时间
func (c Carbon) EndOfHour() Carbon {
	if c.IsZero() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), 59, 59, 999999999, c.Loc)
	return c
}

// StartOfMinute return a Carbon instance for start of the minute
// 分钟开始时间
func (c Carbon) StartOfMinute() Carbon {
	if c.IsZero() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), 0, 0, c.Loc)
	return c
}

// EndOfMinute return a Carbon instance for end of the minute
// 分钟结束时间
func (c Carbon) EndOfMinute() Carbon {
	if c.IsZero() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), 59, 999999999, c.Loc)
	return c
}

// StartOfSecond return a Carbon instance for start of the second
// 秒开始时间
func (c Carbon) StartOfSecond() Carbon {
	if c.IsZero() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), 0, c.Loc)
	return c
}

// EndOfSecond return a Carbon instance for end of the second
// 秒结束时间
func (c Carbon) EndOfSecond() Carbon {
	if c.IsZero() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), 999999999, c.Loc)
	return c
}
