package satellite

import "image"

const URI string = "https://api.met.no/weatherapi/geosattelite/1.4/"

const (
	AreaAfrica        string = "africa"
	AreaEurope        string = "europe"
	AreaAtlanticOcean string = "atlantic_ocean"
	AreaMediterranean string = "mediterranean"
	AreaGlobal        string = "global"
)

const (
	TypeInfrared string = "infrared"
	TypeVisible  string = "visible"
)

const (
	SizeSmall  string = "small"
	SizeNormal string = "normal"
)

func Image() (*image.Image, error) {
	return nil, nil
}

func Available() error {
	return nil
}
