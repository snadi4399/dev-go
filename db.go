package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	// Update with correct MySQL credentials
	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/cityapp")
	if err != nil {
		log.Fatal(err)
	}
}

// Fetch all cities
func getAllCities() []City {
	rows, err := db.Query("SELECT city_id, city_name, latitude, longitude, temperature_unit, latest_temperature FROM cities")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var cities []City
	for rows.Next() {
		var city City
		rows.Scan(&city.CityID, &city.CityName, &city.Latitude, &city.Longitude, &city.TemperatureUnit, &city.LatestTemperature)
		cities = append(cities, city)
	}
	return cities
}

// Fetch a city by ID
func getCityByID(cityID string) City {
	var city City
	err := db.QueryRow("SELECT city_id, city_name, latitude, longitude, temperature_unit, latest_temperature FROM cities WHERE city_id = ?", cityID).Scan(
		&city.CityID, &city.CityName, &city.Latitude, &city.Longitude, &city.TemperatureUnit, &city.LatestTemperature)
	if err != nil {
		log.Fatal(err)
	}
	return city
}
