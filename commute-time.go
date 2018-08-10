package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
	"googlemaps.github.io/maps"
)

func main() {
	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyCL2wxIPmfmMbzzHY1wyQWc88c0nWqR6Bo"))
	if err != nil {
		log.Fatal("Fatal error: %s", err)
	}

	startTime := roundToNearestHalfHourInFuture()
	endTime := startTime.Add(6 * time.Hour)
	home := "1175 Cypress St, East Palo Alto, CA"
	work := "Apple Park, Cupertino, CA"

	createAndSaveGraph(c, startTime, endTime, "To Work", "to-work.png", home, work)
	createAndSaveGraph(c, startTime, endTime, "From Work", "from-work.png", work, home)
}

func createAndSaveGraph(c *maps.Client, startTime time.Time, endTime time.Time, title string, fileName string, startLoc string, endLoc string) {
	currentTime := startTime

	timestamps := []time.Time{}
	trafficModels := []maps.TrafficModel{maps.TrafficModelBestGuess, maps.TrafficModelPessimistic, maps.TrafficModelOptimistic}
	avgTravelTimes := []float64{}
	worstTravelTimes := []float64{}
	bestTravelTimes := []float64{}

	for currentTime.Before(endTime) {
		var wg sync.WaitGroup
		timestamps = append(timestamps, currentTime)
		wg.Add(len(trafficModels))

		for i, trafficModel := range trafficModels {
			go func(loop int, trafficModel maps.TrafficModel) {
				r := &maps.DistanceMatrixRequest{
					Origins:       []string{startLoc},
					Destinations:  []string{endLoc},
					DepartureTime: strconv.FormatInt(currentTime.Unix(), 10),
					Mode:          maps.TravelModeDriving,
					Units:         maps.UnitsImperial,
					TrafficModel:  trafficModel,
				}

				matrix, err := c.DistanceMatrix(context.Background(), r)

				if err != nil {
					log.Fatal("Fatal error: %s", err)
				}

				switch loop {
				case 0:
					avgTravelTimes = append(avgTravelTimes, matrix.Rows[0].Elements[0].DurationInTraffic.Minutes())
				case 1:
					worstTravelTimes = append(worstTravelTimes, matrix.Rows[0].Elements[0].DurationInTraffic.Minutes())
				case 2:
					bestTravelTimes = append(bestTravelTimes, matrix.Rows[0].Elements[0].DurationInTraffic.Minutes())
				}
				wg.Done()
			}(i, trafficModel)
		}

		wg.Wait()
		currentTime = nextTimeInGraph(currentTime)
	}

	graph := chart.Chart{
		Title:      title,
		TitleStyle: chart.StyleShow(),
		Series: []chart.Series{
			chart.TimeSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: drawing.ColorBlack,
					StrokeWidth: 2,
				},
				Name:    "Best Guess",
				XValues: timestamps,
				YValues: avgTravelTimes,
			},
			chart.TimeSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: drawing.ColorRed,
					StrokeWidth: 2,
				},
				Name:    "Pessimistic",
				XValues: timestamps,
				YValues: worstTravelTimes,
			},
			chart.TimeSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: drawing.ColorGreen,
					StrokeWidth: 2,
				},
				Name:    "Optimistic",
				XValues: timestamps,
				YValues: bestTravelTimes,
			},
		},
		XAxis: chart.XAxis{
			Name:           "Leave Time",
			ValueFormatter: formatTime,
			NameStyle:      chart.StyleShow(),
			Style: chart.Style{
				TextRotationDegrees: 90,
				Show:                true,
			},
			Ticks: generateTicks(timestamps),
		},

		YAxis: chart.YAxis{
			Name:      "Travel Time",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
			Range: &chart.ContinuousRange{
				Min: 15.0,
				Max: 50.0,
			},
			Ticks: []chart.Tick{
				{15.0, "15.0"},
				{20.0, "20.0"},
				{25.0, "25.0"},
				{30.0, "30.0"},
				{35.0, "35.0"},
				{40.0, "40.0"},
				{45.0, "45.0"},
				{50.0, "50.0"},
			},
		},
	}

	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}

	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Fatal error: %s", err)
	}

	err = graph.Render(chart.PNG, f)

	if err != nil {
		log.Fatalf("Fatal error: %s", err)
	}
}
