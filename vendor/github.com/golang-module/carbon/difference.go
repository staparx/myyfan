package carbon

import (
	"strings"
)

// DiffInYears get the difference in years
// 相差多少年
func (c Carbon) DiffInYears(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return c.DiffInMonths(end) / MonthsPerYear
}

// DiffInYearsWithAbs get the difference in years with absolute value
// 相差多少年(绝对值)
func (c Carbon) DiffInYearsWithAbs(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return getAbsValue(c.DiffInYears(end))
}

// DiffInMonths get the difference in months
// 相差多少月
func (c Carbon) DiffInMonths(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	dy, dm, dd := end.Year()-c.Year(), end.Month()-c.Month(), end.Day()-c.Day()
	if dd < 0 {
		dm = dm - 1
	}
	if dy == 0 && dm == 0 {
		return 0
	}
	if dy == 0 && dm != 0 && dd != 0 {
		if int(end.DiffInHoursWithAbs(c)) < c.DaysInMonth()*HoursPerDay {
			return 0
		}
		return int64(dm)
	}
	return int64(dy*MonthsPerYear + dm)
}

// DiffInMonthsWithAbs get the difference in months with absolute value
// 相差多少月(绝对值)
func (c Carbon) DiffInMonthsWithAbs(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return getAbsValue(c.DiffInMonths(end))
}

// DiffInWeeks get the difference in weeks
// 相差多少周
func (c Carbon) DiffInWeeks(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return c.DiffInDays(end) / DaysPerWeek
}

// DiffInWeeksWithAbs get the difference in weeks with absolute value
// 相差多少周(绝对值)
func (c Carbon) DiffInWeeksWithAbs(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return getAbsValue(c.DiffInWeeks(end))
}

// DiffInDays get the difference in days
// 相差多少天
func (c Carbon) DiffInDays(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return c.DiffInSeconds(end) / SecondsPerDay
}

// DiffInDaysWithAbs get the difference in days with absolute value
// 相差多少天(绝对值)
func (c Carbon) DiffInDaysWithAbs(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return getAbsValue(c.DiffInDays(end))
}

// DiffInHours get the difference in hours
// 相差多少小时
func (c Carbon) DiffInHours(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return c.DiffInSeconds(end) / SecondsPerHour
}

// DiffInHoursWithAbs get the difference in hours with absolute value
// 相差多少小时(绝对值)
func (c Carbon) DiffInHoursWithAbs(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return getAbsValue(c.DiffInHours(end))
}

// DiffInMinutes get the difference in minutes
// 相差多少分钟
func (c Carbon) DiffInMinutes(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return c.DiffInSeconds(end) / SecondsPerMinute
}

// DiffInMinutesWithAbs get the difference in minutes with absolute value
// 相差多少分钟(绝对值)
func (c Carbon) DiffInMinutesWithAbs(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return getAbsValue(c.DiffInMinutes(end))
}

// DiffInSeconds get the difference in seconds
// 相差多少秒
func (c Carbon) DiffInSeconds(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return end.ToTimestamp() - c.ToTimestamp()
}

// DiffInSecondsWithAbs get the difference in seconds with absolute value
// 相差多少秒(绝对值)
func (c Carbon) DiffInSecondsWithAbs(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return getAbsValue(c.DiffInSeconds(end))
}

// DiffForHumans get the difference in a human readable format, i18n is supported
// 获取对人类友好的可读格式时间差，支持i18n
func (c Carbon) DiffForHumans(carbon ...Carbon) string {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	unit, diff := "", int64(0)
	switch true {
	case c.DiffInYearsWithAbs(end) > 0:
		unit = "year"
		diff = c.DiffInYearsWithAbs(end)
		break
	case c.DiffInMonthsWithAbs(end) > 0:
		unit = "month"
		diff = c.DiffInMonthsWithAbs(end)
		break
	case c.DiffInWeeksWithAbs(end) > 0:
		unit = "week"
		diff = c.DiffInWeeksWithAbs(end)
		break
	case c.DiffInDaysWithAbs(end) > 0:
		unit = "day"
		diff = c.DiffInDaysWithAbs(end)
		break
	case c.DiffInHoursWithAbs(end) > 0:
		unit = "hour"
		diff = c.DiffInHoursWithAbs(end)
		break
	case c.DiffInMinutesWithAbs(end) > 0:
		unit = "minute"
		diff = c.DiffInMinutesWithAbs(end)
	case c.DiffInSecondsWithAbs(end) > 0:
		unit = "second"
		diff = c.DiffInSecondsWithAbs(end)
	case c.DiffInSecondsWithAbs(end) == 0:
		unit = "now"
		diff = 0
		return c.Lang.translate(unit, diff)
	}
	translation := c.Lang.translate(unit, diff)
	if c.Lt(end) && len(carbon) == 0 {
		return strings.Replace(c.Lang.resources["ago"], "%s", translation, 1)
	}
	if c.Lt(end) && len(carbon) > 0 {
		return strings.Replace(c.Lang.resources["before"], "%s", translation, 1)
	}
	if c.Gt(end) && len(carbon) == 0 {
		return strings.Replace(c.Lang.resources["from_now"], "%s", translation, 1)
	}
	if c.Gt(end) && len(carbon) > 0 {
		return strings.Replace(c.Lang.resources["after"], "%s", translation, 1)
	}
	return translation
}
