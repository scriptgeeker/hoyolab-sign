package main

import (
	"time"
)

type Fn func() error

type Timer struct {
	Ticker *time.Ticker
	Runner Fn
}

// NewTimer 构造定时器
func NewTimer(interval int, function Fn) *Timer {
	return &Timer{
		Ticker: time.NewTicker(time.Duration(interval) * time.Second),
		Runner: function,
	}
}

// Start 启动定时任务
func (timer *Timer) Start() {
	for {
		select {
		case <-timer.Ticker.C:
			err := timer.Runner()
			PrintError(err)
		}
	}
}

// SetInterval 循环执行任务
func SetInterval(interval int, function Fn) {
	// 立即执行一次
	err := function()
	if err != nil {
		return
	}
	// 再循环执行
	duration := time.Duration(interval) * time.Second
	timer := time.NewTimer(duration)
	for {
		select {
		case <-timer.C:
			err := function()
			if err == nil {
				timer.Reset(duration)
			}
		}
	}
}
