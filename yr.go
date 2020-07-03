package yr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type ForecastUnits struct {
	WindFromDir                 string `json:"wind_from_direction"`
	PrecipitationAmountMax      string `json:"precipitation_amount_max"`
	DewPtTemp                   string `json:"dew_point_temperature"`
	AirTempMin                  string `json:"air_temperature_min"`
	CloudAreaFracHigh           string `json:"cloud_area_fraction_high"`
	UltravioletIndexClearSkyMax string `json:"ultraviolet_index_clear_sky_max"`
	ProbPrecipitation           string `json:"probability_of_precipitation"`
	RelHumidity                 string `json:"relative_humidity"`
	AirTemp                     string `json:"air_temperature"`
	WindSpeedGust               string `json:"wind_speed_of_gust"`
	ProbThunder                 string `json:"probability_of_thunder"`
	FogAreaFrac                 string `json:"fog_area_fraction"`
	CloudAreaFrac               string `json:"cloud_area_fraction"`
	CloudAreaFracMedium         string `json:"cloud_area_fraction_medium"`
	WindSpeed                   string `json:"wind_speed"`
	CloudAreaFracLow            string `json:"cloud_area_fraction_low"`
	AirPressureAtSeaLvl         string `json:"air_pressure_at_sea_level"`
	PrecipitationAmount         string `json:"precipitation_amount"`
	AirTempMax                  string `json:"air_temperature_max"`
	PrecipitationAmountMin      string `json:"precipitation_amount_min"`
}

type ForecastMeta struct {
	UpdatedAt time.Time     `json:"updated_at"`
	Units     ForecastUnits `jsong:"units"`
}

type ForecastTimePeriod struct {
	PrecipitationAmountMin    float64 `json:"precipitation_amount_min"`
	ProbThunder               float64 `json:"probability_of_thunder"`
	AirTempMax                float64 `json:"air_temperature_max"`
	PrecipitationAmount       float64 `json:"precipitation_amount"`
	ProbPrecipitation         float64 `json:"probability_of_precipitation"`
	UltravioletIdxClearSkyMax float64 `json:"ultraviolet_index_clear_sky_max"`
	AirTempMin                float64 `json:"air_temperature_min"`
	PrecipitationAmountMax    float64 `json:"precipitation_amount_max"`
}
type ForecastTimeStepPeriod struct {
	Details ForecastTimePeriod `json:"period"`
}

type ForecastTimeInstant struct {
	AirTemp             float64 `json:"air_temperature"`
	CloudAreaFrac       float64 `json:"cloud_area_fraction"`
	WindSpeedGust       float64 `json:"wind_speed_of_gust"`
	CloudAreaFracLow    float64 `json:"cloud_area_fraction_low"`
	WindSpeed           float64 `json:"wind_speed"`
	CloudAreaFracMedium float64 `json:"cloud_area_fraction_medium"`
	AirPressureAtSeaLvl float64 `json:"air_pressure_at_sea_level"`
	WindFromDir         float64 `json:"wind_from_direction"`
	ForAreaFrac         float64 `json:"fog_area_fraction"`
	DewPTemp            float64 `json:"dew_point_temperature"`
	CloudAreaFracHigh   float64 `json:"cloud_area_fraction_high"`
	RelHumidity         float64 `json:"relative_humidity"`
}

type ForecastTimeStepInstant struct {
	Details ForecastTimeInstant `json:"details"`
}

type ForecastTimeStepData struct {
	Next1H  ForecastTimeStepPeriod  `json:"next_1_hours"`
	Next6H  ForecastTimeStepPeriod  `json:"next_6_hours"`
	Instant ForecastTimeStepInstant `json:"instant"`
}

type ForecastTimeStep struct {
	Time time.Time            `json:"time"`
	Data ForecastTimeStepData `json:"data"`
}

type Forecast struct {
	Meta       ForecastMeta       `json:"meta"`
	TimeSeries []ForecastTimeStep `json:"timeseries"`
}

type PointGeometry struct {
	Coordinates []float64 `json:"coordinates"`
	Type        string    `json:"type"`
}

type METJSONForecast struct {
	Properties Forecast      `json:"properties"`
	Type       string        `json:"features"`
	Geometry   PointGeometry `json:"geometry"`
}

const (
	URI             string = "https://api.met.no/weatherapi/locationforecast/2.0"
	EndpointCompact string = "/compact"
)

func Compact(lat, lon float64) (*METJSONForecast, error) {
	v := url.Values{}
	v.Add("lat", fmt.Sprintf("%.2f", lat))
	v.Add("lon", fmt.Sprintf("%.2f", lon))

	resp, e := http.Get(fmt.Sprintf("%s%s?%s", URI, EndpointCompact, v.Encode()))
	var f METJSONForecast
	if e != nil {
		return nil, e
	} else if e := json.NewDecoder(resp.Body).Decode(&f); e != nil {
		return nil, e
	}

	return &f, nil
}
