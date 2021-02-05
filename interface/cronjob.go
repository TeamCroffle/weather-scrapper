package _interface

import (
	"time"
)

type (
	Cronjob interface {
		Run() error
		GetName() string
		GetLatestExecuteTime() time.Time
		SetLatestExecuteTime(t time.Time)
		IsTimeToRun() bool
	}
)

