package main

import "go.uber.org/zap"

var log, _ = zap.NewProduction()

type CronZapLog struct {
	CronZapLog *zap.Logger
}

func (c CronZapLog) Info(msg string, keysAndValues ...interface{}) {
	c.CronZapLog.Sugar().Infow(msg, keysAndValues...)
}

func (c CronZapLog) Error(err error, msg string, keysAndValues ...interface{}) {
	c.CronZapLog.With(zap.Error(err)).Sugar().Errorw(msg, keysAndValues...)
}

func getCronLog() *CronZapLog {
	return &CronZapLog{CronZapLog: log}
}
