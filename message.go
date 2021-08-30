package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
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
