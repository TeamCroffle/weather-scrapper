package main

import (
	"context"
	"fmt"
	_interface "github.com/TeamCroffle/weather-scrapper/interface"
	. "github.com/TeamCroffle/weather-scrapper/job"
	"time"
)

func main() {
	jobs := []_interface.Cronjob{
		NewKoreaWeatherSource("기상청"),
		// NewOpenWeatherSource("OpenWeather"),
	}
	// Config MongoDB client with connection
	dbClient := MongoConn()
	defer dbClient.Disconnect(context.Background())
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cronHistory := NewCronjobHistory(dbClient, &ctx)

	for _, j := range jobs {
		lastExecuteTime := cronHistory.GetLastTime(j.GetName())
		if !j.IsTimeToRun(lastExecuteTime) {
			fmt.Println("Oh, Hello", j.GetName())
			continue
		}
	}
}

