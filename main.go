package main

import (
	"context"
	"time"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	InitConfig()
	schedule := cron.New(
		cron.WithLocation(time.Local),
		cron.WithLogger(getCronLog()),
		cron.WithSeconds(),
	)

	//午餐消息定时推送
	for key, lunch := range cfg.Crons.Lunch {
		sort := key
		if _, err := schedule.AddFunc(lunch, func() {
			color := NewColor()
			//午餐消息内容组装
			title := titleTemplate.ExecuteString(map[string]interface{}{
				"emoji": Picker.Pick(loveEmojis...),
			})
			date, err := GetTodayTime()
			if err != nil {
				log.Error(err.Error())
				return
			}
			//筛选消息，第一条为订餐提醒，第二条为用餐提醒
			var body string
			var extra string
			switch sort {
			case 0:
				body = lunchOrderTemplate.ExecuteString(map[string]interface{}{
					"emoji1": Picker.Pick(foodEmojis...),
					"emoji2": Picker.Pick(loveEmojis...),
				})
			case 1:
				body = lunchTemplate.ExecuteString(map[string]interface{}{
					"emoji1": Picker.Pick(foodEmojis...),
					"emoji2": Picker.Pick(loveEmojis...),
				})
			default:
				body = defaultTemplate.ExecuteString(map[string]interface{}{
					"emoji": "💢",
				})
			}

			markdown := markdownTemplate.ExecuteString(map[string]interface{}{
				"title": color.GetColorWord(color.warning, title),
				"date":  date,
				"body":  body,
				"extra": color.GetColorWord(color.comment, extra),
			})
			log.Info("get markdown text", zap.String("markdown", markdown))
			req := &msgReq{
				Msgtype: "markdown",
				Markdown: struct {
					Content string `json:"content"`
				}{
					Content: markdown,
				},
			}
			pushResp := req.PushMsg(ctx)
			log.Info("push response", zap.Any("response", pushResp))
		}); err != nil {
			log.Error("add lunch cron error", zap.Error(err))
		}
	}

	//晚餐消息定时推送
	for key, dinner := range cfg.Crons.Dinner {
		sort := key
		if _, err := schedule.AddFunc(dinner, func() {
			color := NewColor()
			//晚餐消息内容组装
			title := titleTemplate.ExecuteString(map[string]interface{}{
				"emoji": Picker.Pick(loveEmojis...),
			})
			date, err := GetTodayTime()
			if err != nil {
				log.Error(err.Error())
				return
			}

			//筛选消息，第一条为订餐提醒，第二条为用餐提醒
			var body string
			var extra string
			switch sort {
			case 0:
				body = dinnerOrderTemplate.ExecuteString(map[string]interface{}{
					"emoji1": Picker.Pick(foodEmojis...),
					"emoji2": Picker.Pick(loveEmojis...),
				})
				extra = extraTemplate.ExecuteString(map[string]interface{}{
					"extra": cfg.Account.YouFanURL,
				})
			case 1:
				body = dinnerTemplate.ExecuteString(map[string]interface{}{
					"emoji1": Picker.Pick(foodEmojis...),
					"emoji2": Picker.Pick(loveEmojis...),
				})
			default:
				body = defaultTemplate.ExecuteString(map[string]interface{}{
					"emoji": "💢",
				})
			}
			markdown := markdownTemplate.ExecuteString(map[string]interface{}{
				"title": color.GetColorWord(color.warning, title),
				"date":  date,
				"body":  body,
				"extra": extra,
			})
			log.Info("get markdown text", zap.String("markdown", markdown))
			req := &msgReq{
				Msgtype: "markdown",
				Markdown: struct {
					Content string `json:"content"`
				}{
					Content: markdown,
				},
			}
			pushResp := req.PushMsg(ctx)
			log.Info("push response", zap.Any("response", pushResp))
		}); err != nil {
			log.Error("add dinner cron error", zap.Error(err))
		}
	}
	schedule.Run()

}
