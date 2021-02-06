package _interface

import "time"

type CronjobHistory interface {
	GetLastTime(name string) (time.Time, bool)
	LogExecute(name string, executeTime time.Time) bool
}
