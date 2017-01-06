package yr

import (
	"encoding/xml"
	"fmt"
	"time"
)

const timestampFmt string = "2006-01-02T15:04:05"

// Custom YR timestamp
type Timestamp struct {
	time.Time
}

// Unmarshal function for timestamp
func (timestamp *Timestamp) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var value string
	dec.DecodeElement(&value, &start)
	t, e := time.Parse(timestampFmt, value)
	if e != nil {
		return e
	}
	*timestamp = Timestamp{t}
	return nil
}

// Unmarshal function for attribute timestamps
func (timestamp *Timestamp) UnmarshalXMLAttr(attr xml.Attr) error {
	t, e := time.Parse(timestampFmt, attr.Value)
	if e != nil {
		return e
	}
	*timestamp = Timestamp{t}
	return nil
}

// Forecast response from API
type Forecast struct {
	XMLName xml.Name `xml:"time"`

	// From and To tell the time range for this forecast
	From Timestamp `xml:"from,attr"`
	To   Timestamp `xml:"to,attr"`

	// Temp is the temperature for the time range
	Temp struct {
		XMLName xml.Name `xml:"temperature"`
		Value   int      `xml:"value,attr"`
		Unit    string   `xml:"unit,attr"`
	}

	// Summary does a verbose string summary of the forecast
	Summary struct {
		XMLName xml.Name `xml:"symbol"`
		Name    string   `xml:"name,attr"`
	}

	// WindSpeed in m/s (meters per second)
	WindSpeed struct {
		XMLName xml.Name `xml:"windSpeed"`
		Mps     float32  `xml:"mps,attr"`
	}

	// WindDirection with a code and more verbose name
	WindDirection struct {
		XMLName xml.Name `xml:"windDirection"`
		Degrees float32  `xml:"deg,attr"`
		Code    string   `xml:"code,attr"`
		Name    string   `xml:"name,attr"`
	}

	// Precipitation in mm during the time range.
	Precipitation struct {
		XMLName xml.Name `xml:"precipitation"`
		Value   float32  `xml:"value>attr"`
	}
}

// Debug: string representation of forecast
func (forecast Forecast) String() string {
	return fmt.Sprintf("%s -> %s:\n %d°C [%s] %.0f m/s [%s], %.0f mm",
		forecast.From,
		forecast.To,
		forecast.Temp.Value,
		forecast.Summary.Name,
		forecast.WindSpeed.Mps,
		forecast.WindDirection.Code,
		forecast.Precipitation.Value)
}

// Location represents the location the response data is for.
type Location struct {
	Name    string `xml:"name"`    // Name of the location
	Type    string `xml:"type"`    // Type of the location, e.g. city
	Country string `xml:"country"` // Country of the location
}

// Response from API.
type Response struct {
	XMLName   xml.Name   `xml:"weatherdata"`
	Location  Location   `xml:"location"`
	Forecasts []Forecast `xml:"forecast>tabular>time"`
}
