package _interface

import (
	"time"
)

type (
	Cronjob interface {
		Run() error
		GetName() string
		IsTimeToRun(executeTime time.Time) bool
	}
)

