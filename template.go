package main

import (
	"strconv"
	"time"

	"github.com/golang-module/carbon"
	"github.com/valyala/fasttemplate"
)

var (
	//Markdown消息
	markdownTemplate = fasttemplate.New("## {{title}} \n > ### {{date}} \n  \n > ## {{body}} \n \n ###### {{extra}}  ", "{{", "}}")
	//消息标题
	titleTemplate = fasttemplate.New("亲爱的主人{{emoji}}~", "{{", "}}")
	//消息推送日期
	dateTemplate = fasttemplate.New("现在是 {{month}} 月 {{day}} 日  {{week}}  {{time}}", "{{", "}}")

	//提醒模板
	lunchOrderTemplate  = fasttemplate.New("{{emoji1}} 想好吃什么了吗？想好记得提前点餐哦~ 不然~~ 可是会耽误午休的呢~ {{emoji2}}", "{{", "}}")
	lunchTemplate       = fasttemplate.New("{{emoji1}} 到点吃饭啦~ 到点吃饭啦~ 早点吃饭哟…… 免得要排队哦~ {{emoji2}}", "{{", "}}")
	dinnerOrderTemplate = fasttemplate.New("{{emoji1}} 我才不会提醒你点加班餐的 人家才不要你加班呢~~ 哼~{{emoji2}}", "{{", "}}")
	//todo 获取悠饭的订单数据
	dinnerTemplate = fasttemplate.New("{{emoji1}} 加班餐到啦！早点吃饭~ 不然饭菜可是会冷了的哦~ {{emoji2}}", "{{", "}}")

	//附件信息
	extraTemplate = fasttemplate.New("[(悄咪咪的递上小纸条💌)]({{extra}})", "{{", "}}")

	//默认消息
	defaultTemplate = fasttemplate.New("我只顺路过来看看而已… 你…… 你…… 你可别想多了！哼~ {{emoji}} ", "{{", "}}")
)

var (
	foodEmojis = []string{"🍏", "🍎", "🍐", "🍊", "🍋", "🍌", "🍉", "🍇", "🍓", "🍈", "🍒", "🍑", "🍍", "🥥", "🥝",
		"🍅", "🥑", "🍆", "🌶", "🥒", "🥦", "🌽", "🥕", "🥗", "🥔", "🍠", "🥜", "🍯", "🍞", "🥐", "🥖", "🥨", "🥞",
		"🧀", "🍗", "🍖", "🥩", "🍤", "🥚", "🥚", "🍳", "🥓", "🍔", "🍟", "🌭", "🍕", "🍝", "🥪", "🌮", "🌯", "🥙",
		"🍜", "🍲", "🥘", "🍥", "🍱", "🍣", "🍙", "🍛", "🍘", "🍚", "🥟", "🍢", "🍡", "🍧", "🍨", "🍦", "🍰", "🎂",
		"🥧", "🍮", "🍭", "🍬", "🍫", "🍿", "🍩", "🍪", "🥠", "☕", "🍵", "🥣", "🍼", "🥤", "🥛", "🍺", "🍻", "🍷",
		"🥂", "🥃", "🍸", "🍹", "🍾", "🍶", "🥄", "🍴", "🍽", "🥢"}
	loveEmojis = []string{"💝", "💞", "💟", "💘", "❤"}
)

func GetWeekStr(dayOfWeek int) string {
	switch dayOfWeek {
	case 1:
		return "星期一"
	case 2:
		return "星期二"
	case 3:
		return "星期三"
	case 4:
		return "星期四"
	case 5:
		return "星期五"
	case 6:
		return "星期六"
	case 7:
		return "星期日"
	}
	return ""
}

// 字体颜色(只支持3种内置颜色)
// info:绿色
// comment:灰色
// waring:橙红色
// @author huangsx01
type color struct {
	info    string
	comment string
	warning string
}

func NewColor() *color {
	return &color{
		info:    "info",
		comment: "comment",
		warning: "warning",
	}
}

func (c *color) GetColorWord(color string, word string) string {
	return fasttemplate.New(" <font color=\"{{color}}\"> {{word}} </font>", "{{", "}}").
		ExecuteString(map[string]interface{}{
			"color": color,
			"word":  word,
		})
}

// GetTodayTime 获取现在时间
func GetTodayTime() string {
	month := carbon.CreateFromTimestamp(time.Now().Unix()).Month()
	weekdayStr := GetWeekStr(carbon.CreateFromTimestamp(time.Now().Unix()).DayOfWeek())
	day := carbon.CreateFromTimestamp(time.Now().Unix()).Day()
	timer := time.Now().Format("15:04:05")
	date := dateTemplate.ExecuteString(map[string]interface{}{
		"month": strconv.Itoa(month),
		"day":   strconv.Itoa(day),
		"week":  weekdayStr,
		"time":  timer,
	})
	return date
}
