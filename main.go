package main

import (
	"context"
	"strconv"
	"time"

	"github.com/golang-module/carbon"
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
	for _, lunch := range cfg.Crons.Lunch {
		if _, err := schedule.AddFunc(lunch, func() {
			color := NewColor()
			//午餐消息内容组装
			title := titleTemplate.ExecuteString(map[string]interface{}{
				"emoji": Picker.Pick(loveEmojis...),
			})
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
			body := lunchTemplate.ExecuteString(map[string]interface{}{
				"emoji1": Picker.Pick(foodEmojis...),
				"emoji2": Picker.Pick(loveEmojis...),
			})
			markdown := markdownTemplate.ExecuteString(map[string]interface{}{
				"title": color.GetColorWord(color.warning, title),
				"date":  color.GetColorWord(color.comment, date),
				"body":  color.GetColorWord(color.comment, body),
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
	for _, dinner := range cfg.Crons.Dinner {
		if _, err := schedule.AddFunc(dinner, func() {
			color := NewColor()
			//晚餐消息内容组装
			title := titleTemplate.ExecuteString(map[string]interface{}{
				"emoji": Picker.Pick(loveEmojis...),
			})
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
			body := dinnerTemplate.ExecuteString(map[string]interface{}{
				"emoji1": Picker.Pick(foodEmojis...),
				"emoji2": Picker.Pick(loveEmojis...),
			})
			markdown := markdownTemplate.ExecuteString(map[string]interface{}{
				"title": color.GetColorWord(color.warning, title),
				"date":  color.GetColorWord(color.comment, date),
				"body":  color.GetColorWord(color.comment, body),
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
