package main

import (
	"time"

	chart "github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/util"
)

func formatTime(v interface{}) string {
	dateFormat := "Mon - 3:04 PM"
	if typed, isTyped := v.(time.Time); isTyped {
		return typed.Format(dateFormat)
	}
	if typed, isTyped := v.(int64); isTyped {
		return time.Unix(0, typed).Format(dateFormat)
	}
	if typed, isTyped := v.(float64); isTyped {
		return time.Unix(0, int64(typed)).Format(dateFormat)
	}
	return ""
}

func roundToNearestHalfHourInFuture() time.Time {
	now := time.Now()
	var rounded time.Time

	if now.Minute() <= 30 {
		// Round to 30 in future
		rounded = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 30, 0, 0, now.Location())
	} else {
		// Round to 00 in future
		rounded = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
		rounded = rounded.Add(1 * time.Hour)
	}

	if rounded.Before(now) {
		rounded.Add(30 * time.Minute)
	}

	return rounded
}

func nextTimeInGraph(current time.Time) time.Time {
	println(current.String())
	return current.Add(15 * time.Minute)
}

func generateTicks(times []time.Time) []chart.Tick {
	ticks := []chart.Tick{}

	for _, time := range times {
		if time.Minute() == 0 || time.Minute() == 30 {
			ticks = append(ticks, chart.Tick{
				Value: util.Time.ToFloat64(time),
				Label: formatTime(time),
			})
		}
	}

	return ticks
}

func timeOfDay(now time.Time, noon time.Time) string {
	if now.Before(noon) {
		return "this morning"
	}
	return "this afternoon"
}
