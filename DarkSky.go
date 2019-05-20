package ColoradoWeatherMap

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type TimePeriodData struct {
	Summary string // This isn't handling some symbols, like '<'
	Icon    string
}

type DataPoint struct {
	ApparentTemperature float64
	Temperature         float64
	Summary             string
	CloudCover          float64
	PrecipIntensity     float64
	PrecipProbability   float64
	WindSpeed           float64
	Time                int64
}

func (m DataPoint) TypedTime() time.Time {
	return time.Unix(0, m.Time*int64(time.Millisecond))
}

type ForeCast struct {
	Timezone  string
	Currently DataPoint
	Hourly    TimePeriodData
	Daily     TimePeriodData
	Location  string
}

type GpsCoordinates struct {
	Latitude  float64
	Longitude float64
	Location  string
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func stringOf(coordinates GpsCoordinates) string {
	return fmt.Sprintf("%f", coordinates.Latitude) + "," + fmt.Sprintf("%f", coordinates.Longitude)
}

func GetMultipleForecasts(darkSkyToken string, coordinates []GpsCoordinates) []ForeCast {
	forecasts := []ForeCast{}
	return forecasts
}

func GetBasicForecast(darkSkyToken string, coordinates GpsCoordinates, location string) ForeCast {

	req3, _ := http.NewRequest(
		"GET",
		"https://api.darksky.net/forecast/"+
			darkSkyToken+
			"/"+stringOf(coordinates), nil)
	resp3, error := myClient.Do(req3)
	if error != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", error)
		os.Exit(1)
	}

	defer resp3.Body.Close()
	var weatherForecast ForeCast
	error = json.NewDecoder(resp3.Body).Decode(&weatherForecast)
	if error != nil {
		fmt.Println("Error decoding forecast: " + error.Error())
	}
	weatherForecast.Location = location
	return weatherForecast
}
