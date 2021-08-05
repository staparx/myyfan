package carbon

// DaysInYear get days in year
// 获取本年的总天数
func (c Carbon) DaysInYear() int {
	if c.IsInvalid() {
		return 0
	}
	days := DaysPerNormalYear
	if c.IsLeapYear() {
		days = DaysPerLeapYear
	}
	return days
}

// DaysInMonth get days in month
// 获取本月的总天数
func (c Carbon) DaysInMonth() int {
	if c.IsInvalid() {
		return 0
	}
	return c.EndOfMonth().Time.In(c.Loc).Day()
}

// MonthOfYear get month of year
// 获取本年的第几月
func (c Carbon) MonthOfYear() int {
	if c.IsInvalid() {
		return 0
	}
	return int(c.Time.In(c.Loc).Month())
}

// DayOfYear get day of year
// 获取本年的第几天
func (c Carbon) DayOfYear() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Time.In(c.Loc).YearDay()
}

// DayOfMonth get day of month
// 获取本月的第几天
func (c Carbon) DayOfMonth() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Time.In(c.Loc).Day()
}

// DayOfWeek get day of week
// 获取本周的第几天
func (c Carbon) DayOfWeek() int {
	if c.IsInvalid() {
		return 0
	}
	day := int(c.Time.In(c.Loc).Weekday())
	if day == 0 {
		return DaysPerWeek
	}
	return day
}

// WeekOfYear get week of year. see https://en.wikipedia.org/wiki/ISO_8601#Week_dates
// 获取本年的第几周
func (c Carbon) WeekOfYear() int {
	if c.IsInvalid() {
		return 0
	}
	_, week := c.Time.In(c.Loc).ISOWeek()
	return week
}

// WeekOfMonth get week of month
// 获取本月的第几周
func (c Carbon) WeekOfMonth() int {
	if c.IsInvalid() {
		return 0
	}
	days := c.Day() + c.StartOfMonth().DayOfWeek() - 1
	if days%DaysPerWeek == 0 {
		return days / DaysPerWeek
	}
	return days/DaysPerWeek + 1
}

// Century get current century
// 获取当前世纪
func (c Carbon) Century() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Year()/YearsPerCentury + 1
}

// Decade get current decade
// 获取当前年代
func (c Carbon) Decade() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Year() % YearsPerCentury / YearsPerDecade * YearsPerDecade
}

// Year get current year
// 获取当前年
func (c Carbon) Year() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Time.In(c.Loc).Year()
}

// Quarter get current quarter
// 获取当前季度
func (c Carbon) Quarter() (quarter int) {
	if c.IsInvalid() {
		return 0
	}
	switch {
	case c.Month() >= 10:
		quarter = 4
	case c.Month() >= 7:
		quarter = 3
	case c.Month() >= 4:
		quarter = 2
	case c.Month() >= 1:
		quarter = 1
	}
	return
}

// Month get current month
// 获取当前月
func (c Carbon) Month() int {
	if c.IsInvalid() {
		return 0
	}
	return c.MonthOfYear()
}

// Week get current week, start from 0
// 获取当前周(从0开始)
func (c Carbon) Week() int {
	if c.IsInvalid() {
		return -1
	}
	return int(c.Time.In(c.Loc).Weekday())
}

// Day get current day
// 获取当前日
func (c Carbon) Day() int {
	if c.IsInvalid() {
		return 0
	}
	return c.DayOfMonth()
}

// Hour get current hour
// 获取当前小时
func (c Carbon) Hour() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Time.In(c.Loc).Hour()
}

// Minute get current minute
// 获取当前分钟数
func (c Carbon) Minute() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Time.In(c.Loc).Minute()
}

// Second get current second
// 获取当前秒数
func (c Carbon) Second() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Time.In(c.Loc).Second()
}

// Millisecond get current millisecond
// 获取当前毫秒数
func (c Carbon) Millisecond() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Time.In(c.Loc).Nanosecond() / 1e6
}

// Microsecond get current microsecond
// 获取当前微秒数
func (c Carbon) Microsecond() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Time.In(c.Loc).Nanosecond() / 1e3
}

// Nanosecond get current nanosecond
// 获取当前纳秒数，9位数字
func (c Carbon) Nanosecond() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Time.In(c.Loc).Nanosecond()
}

// Location get location name
// 获取位置
func (c Carbon) Location() string {
	return c.Loc.String()
}

// Timezone get timezone name
// 获取时区
func (c Carbon) Timezone() string {
	name, _ := c.Time.Zone()
	return name
}

// Offset get offset seconds from the UTC timezone
// 获取距离UTC时区的偏移量，单位秒
func (c Carbon) Offset() int {
	_, offset := c.Time.Zone()
	return offset
}

// Locale get locale name
// 获取语言区域
func (c Carbon) Locale() string {
	return c.Lang.locale
}

// Age get age
// 获取年龄
func (c Carbon) Age() int {
	if c.IsInvalid() {
		return 0
	}
	now := Now()
	if c.ToTimestamp() > now.ToTimestamp() {
		return 0
	}
	return int(c.DiffInYears(now))
}
