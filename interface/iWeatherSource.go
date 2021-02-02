package _interface

import "time"

type iWeatherSource interface {
	getName() string
	run() error
	isTimeToRun(lastExecuteTime time.Time) bool
}

type WeatherSource struct {
	name string
}


func (w *WeatherSource) getName() string {
	return w.name
}

func (w *WeatherSource) run() error {
	panic("implement me")
}

func (w *WeatherSource) isTimeToRun(lastExecuteTime time.Time) bool {
	return false
}

