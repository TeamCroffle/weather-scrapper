package main

import (
	"fmt"
	. "github.com/TeamCroffle/weather-scrapper/job"
	_interface "github.com/TeamCroffle/weather-scrapper/interface"

)


func main() {
	jobs := []_interface.Cronjob{
		NewKoreaWeatherSource("기상청"),
		NewOpenWeatherSource("OpenWeather"),
	}

	for _, j := range jobs {
		fmt.Println("Starting job", j.GetName())
	}

}
