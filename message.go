package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/valyala/fasttemplate"
	"go.uber.org/zap"
)

type msgReq struct {
	Msgtype  string `json:"msgtype"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`

	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

type msgResp struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (m msgReq) PushMsg(ctx context.Context) interface{} {
	resp, err := resty.NewWithClient(&http.Client{
		Timeout: 10 * time.Second,
		Transport: func() http.RoundTripper {
			transport := http.DefaultTransport.(*http.Transport).Clone()
			transport.MaxIdleConns = 100
			transport.MaxConnsPerHost = 100
			transport.MaxIdleConnsPerHost = 100
			return transport
		}(),
	}).R().SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetBody(m).
		SetResult(new(msgResp)).
		Post(cfg.Account.QwWebHook)
	if err != nil {
		log.Error("send message error", zap.Error(err))
		return nil
	}
	return resp.Result()
}

var (
	//Markdown消息
	markdownTemplate = fasttemplate.New("## {{title}} \n > ### {{date}} \n  \n > ## {{body}}  ", "{{", "}}")
	//消息标题
	titleTemplate = fasttemplate.New("亲爱的主人{{emoji}}~", "{{", "}}")
	//消息推送日期
	dateTemplate = fasttemplate.New("现在是 {{month}} 月 {{day}} 日  {{week}}  {{time}}", "{{", "}}")

	//提醒模板
	lunchTemplate  = fasttemplate.New("{{emoji1}} 别忘了点餐哦~ 不然~~ 可是会耽误午休哦~ {{emoji2}}", "{{", "}}")
	dinnerTemplate = fasttemplate.New("{{emoji1}} 我才不会提醒你点加班餐的 哼~ 人家才不要你加班呢~~ {{emoji2}}", "{{", "}}")
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
