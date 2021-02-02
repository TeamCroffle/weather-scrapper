package scrape

import (
	"context"
	_interface "github.com/TeamCroffle/weather-scrapper/interface"
	"time"
)

type koreaWeatherCronjob struct {
	name string
	latestExecuteTime time.Time
	dbctx context.Context
}

func (k koreaWeatherCronjob) Run() error {
	panic("implement me")
}

func (k koreaWeatherCronjob) GetName() string {
	return k.name
}

func (k koreaWeatherCronjob) GetLatestExecuteTime() time.Time {
	panic("implement me")
}

func (k koreaWeatherCronjob) SetLatestExecuteTime(t time.Time) {
	panic("implement me")
}

func NewKoreaWeatherSource(name string) _interface.Cronjob{
	return &koreaWeatherCronjob{
		name: name,
	}
}
