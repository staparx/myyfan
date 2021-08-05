package main

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"
	_ "time/tzdata"
)

type (
	config struct {
		App      app
		Timezone string
		Crons    crons
		Account  account
	}
	// App APP相关信息
	app struct {
		Name string
	}
	// Cron 定时相关信息
	crons struct {
		Lunch  []string
		Dinner []string
	}

	// Account 用户相关信息
	account struct {
		QwWebHook string //企业微信的机器人WebHook
	}
)

var cfg = new(config)

var (
	defaultTimezone = "Asia/Shanghai"
)

func InitConfig() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Error("read config error", zap.Error(err))
		return
	}

	if err := viper.Unmarshal(cfg); err != nil {
		log.Error("viper unmarshal error", zap.Error(err))
		return
	}
	//TODO 必须的参数为空时，载入默认配置
	//加载时区
	_ = initTimezone()

	printCfg(cfg)
}

// 初始化时区
// @author huangsx01
func initTimezone() error {
	if cfg.Timezone == "" {
		cfg.Timezone = defaultTimezone
	}
	var err error
	time.Local, err = time.LoadLocation(cfg.Timezone)
	if err != nil {
		log.Error("load timezone error", zap.Error(err))
		return err
	}
	log.Info("load timezone success", zap.String("timezone", cfg.Timezone))
	return nil
}

// 打印配置信息
// @author huangsx01
func printCfg(c *config) {
	log.Info("程序相关信息",
		zap.String("app_name", c.App.Name))

	log.Info("用户相关信息", zap.String("web_hook", c.Account.QwWebHook))

	log.Info("午餐定时", zap.Any("lunch", c.Crons.Lunch))

	log.Info("晚餐定时", zap.Any("dinner", c.Crons.Dinner))

}
