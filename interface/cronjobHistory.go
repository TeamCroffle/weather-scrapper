package _interface

import "time"

type CronjobHistory interface {
	GetLastTime(name string) time.Time
	LogExecute(name string, executeTime time.Time) bool
}
