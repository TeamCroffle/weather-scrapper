package job


import (
	_interface "github.com/TeamCroffle/weather-scrapper/interface"
)

type koreaWeatherSource struct {
	_interface.WeatherSource
}

func newKoreaWeatherSource() iWeatherSource {
	return &koreaWeatherSource{
		weatherSource: weatherSource{
			name: "한국 기상청",
		},
	}
}
