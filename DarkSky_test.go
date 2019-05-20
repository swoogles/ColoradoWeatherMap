package ColoradoWeatherMap

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetBasicForecast(t *testing.T) {
	//var darkSkyToken = os.Getenv("DARK_SKY_TOKEN")
	//var forecast = GetBasicForecast(darkSkyToken)
	//fmt.Println(forecast)

	var decodedForeCast ForeCast
	error := json.Unmarshal(json.RawMessage(SampleData()), &decodedForeCast)
	if error != nil {
		panic(error)
	}
	fmt.Println("Typed time: " + decodedForeCast.Currently.TypedTime().String())
	fmt.Println(fmt.Sprintf("%f", decodedForeCast.Currently.Temperature))
	//t.Errorf("Failed forecast: %f", forecast.Currently.Temperature)

	out, err := json.MarshalIndent(decodedForeCast, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
