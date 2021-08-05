package carbon

import (
	"fmt"
)

// 无效的时区错误
var invalidTimezoneError = func(timezone string) error {
	return fmt.Errorf("invalid timezone %q, please see the file %q for all valid timezones", timezone, "$GOROOT/lib/time/zoneinfo.zip")
}

// 无效的持续时长错误
var invalidDurationError = func(duration string) error {
	return fmt.Errorf("invalid duration %q, please make sure the duration is valid", duration)
}

// 无效的时间字符串错误
var invalidValueError = func(value string) error {
	return fmt.Errorf("cannot parse %q to carbon, please make sure the value is valid", value)
}

// 无效的布局模板错误
var invalidLayoutError = func(value, layout string) error {
	return fmt.Errorf("cannot parse %q as layout %q to carbon, please make sure the value and layout match", value, layout)
}

// 无效的格式化字符错误
var invalidFormatError = func(value, format string) error {
	return fmt.Errorf("cannot parse %q as format %q to carbon, please make sure the value and format match", value, format)
}

// 无效的区域错误
var invalidLocaleError = func(locale, directory string) error {
	return fmt.Errorf("invalid locale %q, please see the directory %q for all valid locales", locale, directory)
}

// 无效的目录错误
var invalidDirError = func(dir string) error {
	return fmt.Errorf("invalid directory %q, please make sure the directory exists", dir)
}

// 无效的json文件错误
var invalidJsonFileError = func(file string) error {
	return fmt.Errorf("invalid json file %q, please make sure the json file is valid", file)
}
