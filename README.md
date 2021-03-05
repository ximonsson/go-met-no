# MET.no API Implementations

Implementations for the various APIs from the [Norwegian Meteorlogy Institure](https://api.met.no).


## Weather

```go
import met/weather
```

Only location forecasts are supported, and the `/compact` endpoint from the [API](https://api.met.no/weatherapi/locationforecast/2.0/documentation).

```go
f, e := weather.LocationForecastCompact(...)
```


WIP for the [radar API](https://api.met.no/weatherapi/radar/2.0/documentation).


## Satellite

WIP on the [geosattelite API](https://api.met.no/weatherapi/geosatellite/1.4/documentation).
