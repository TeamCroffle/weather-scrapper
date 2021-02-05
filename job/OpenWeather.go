package scrape

import (
	"context"
	_interface "github.com/TeamCroffle/weather-scrapper/interface"
	"time"
)

type openWeatherCronjob struct {
	name string
	latestExecuteTime time.Time

	dbctx context.Context
}

func (k openWeatherCronjob) Run() error {
	panic("implement me")
}

func (k openWeatherCronjob) GetName() string {
	return k.name
}

func (k openWeatherCronjob) GetLatestExecuteTime() time.Time {
	panic("implement me")
}

func (k openWeatherCronjob) SetLatestExecuteTime(t time.Time) {
	panic("implement me")
}

func (k openWeatherCronjob) IsTimeToRun() bool {
	return true
}
func NewOpenWeatherSource(name string) _interface.Cronjob{
	return &openWeatherCronjob{
		name: name,
	}
}
