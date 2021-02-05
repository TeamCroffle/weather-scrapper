package main

import (
	"fmt"
	_interface "github.com/TeamCroffle/weather-scrapper/interface"
	. "github.com/TeamCroffle/weather-scrapper/job"
)


func main() {
	jobs := []_interface.Cronjob{
		NewKoreaWeatherSource("기상청"),
		NewOpenWeatherSource("OpenWeather"),
	}

	for _, j := range jobs {
		if !j.IsTimeToRun() {
			continue
		}

		fmt.Println("Starting job", j.GetName())
	}
}
