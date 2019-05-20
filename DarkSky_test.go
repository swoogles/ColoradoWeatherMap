package ColoradoWeatherMap

import (
	"encoding/json"
	"fmt"
	"github.com/swoogles/hugoprojects/lambdas/weather"
	"os"
	"testing"
)

func TestGetLiveBasicForecast(t *testing.T) {
	var darkSkyToken = os.Getenv("DARK_SKY_TOKEN")
	mountainCoordinates := []weather.GpsCoordinates{
		{38.8697, -106.9878},
	}
	var weatherForecast = weather.GetBasicForecast(darkSkyToken, mountainCoordinates[0], "Crested Butte")
	out, _ := json.MarshalIndent(weatherForecast, "", "  ")
	weatherText := string(out)
	fmt.Println(weatherText)
}

func TestGetBasicForecast(t *testing.T) {
	var decodedForeCast ForeCast
	err := json.Unmarshal(json.RawMessage(SampleData()), &decodedForeCast)
	if err != nil {
		panic(err)
	}
	decodedForeCast.Location = "Crested Butte (Test)"
	fmt.Println("Typed time: " + decodedForeCast.Currently.TypedTime().String())
	fmt.Println(fmt.Sprintf("%f", decodedForeCast.Currently.Temperature))

	out, err := json.MarshalIndent(decodedForeCast, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
