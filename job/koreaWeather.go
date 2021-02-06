package scrape

import (
	"context"
)

type koreaWeatherCronjob struct {
	name string
	periodMin float64
}

func (k koreaWeatherCronjob) Run() error {
	panic("implement me")
}

func (k koreaWeatherCronjob) GetName() string {
	return k.name
}

func (k koreaWeatherCronjob) IsTimeToRun(executeTime time.Time) bool {
	duration := time.Since(executeTime)
	if k.periodMin < duration.Minutes() {
		return false
	}
}

func NewKoreaWeatherSource(name string) _interface.Cronjob{
	return &koreaWeatherCronjob{
		name: name,
		periodMin: 60, // 1 hour
	}
}
