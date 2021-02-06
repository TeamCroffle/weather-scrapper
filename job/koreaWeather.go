package scrape

import (
	_interface "github.com/TeamCroffle/weather-scrapper/interface"
	"time"
)

type koreaWeatherCronjob struct {
	name string
	period time.Duration
}

func (k koreaWeatherCronjob) Run() error {
	panic("implement me")
}

func (k koreaWeatherCronjob) GetName() string {
	return k.name
}

func (k koreaWeatherCronjob) IsTimeToRun(executeTime time.Time) bool {
	return time.Now().After(executeTime.Add(k.period))
}

func NewKoreaWeatherSource(name string) _interface.Cronjob{
	return &koreaWeatherCronjob{
		name: name,
	 	period: time.Minute * 10,
	}
}
