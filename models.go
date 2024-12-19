package main

// City struct to represent city data
type City struct {
	CityID            int     `json:"city_id"`
	CityName          string  `json:"city_name"`
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
	TemperatureUnit   string  `json:"temperature_unit"`
	LatestTemperature float64 `json:"latest_temperature"`
}
