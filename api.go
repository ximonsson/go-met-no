package yr

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

const (
	protocol         string = "http"
	host             string = "yr.no"
	fileForecast     string = "forecast.xml"
	fileHourForecast string = "forecast_hour_by_hour.xml"
	placeURL         string = "%s://%s/place/%s/%s/%s/%s"
)

// Return forecast for city.
func CityForecast(country, region, city string) ([]Forecast, error) {
	url := fmt.Sprintf(placeURL, protocol, host, country, region, city, fileForecast)
	res, e := makeAPICall(url)
	if e != nil {
		return nil, e
	}
	return res.Forecasts, nil
}

// Return hour by hour forecast for city.
func HourlyCityForecast(country, region, city string) ([]Forecast, error) {
	url := fmt.Sprintf(placeURL, protocol, host, country, region, city, fileHourForecast)
	res, e := makeAPICall(url)
	if e != nil {
		return nil, e
	}
	return res.Forecasts, nil
}

// Make API call
func makeAPICall(url string) (Response, error) {
	// make request
	var response Response
	res, e := http.Get(url)
	if e != nil {
		return response, e
	}
	defer res.Body.Close()

	// decode data
	dec := xml.NewDecoder(res.Body)
	e = dec.Decode(&response)
	return response, e
}
