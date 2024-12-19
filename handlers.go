package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

// Handler to list all cities
func listCities(w http.ResponseWriter, r *http.Request) {
	cities := getAllCities() // Fetch cities from the database

	// Set the response header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the cities as JSON
	err := json.NewEncoder(w).Encode(cities)
	if err != nil {
		http.Error(w, "Failed to encode cities", http.StatusInternalServerError)
	}
}

// Handler to view a single city
func viewCity(w http.ResponseWriter, r *http.Request) {
	cityID := chi.URLParam(r, "city_id")

	// Fetch city details using cityID
	city := getCityByID(cityID) // Ensure this fetches the correct city

	// Set response header to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode city details into JSON and return it in the response
	err := json.NewEncoder(w).Encode(city)
	if err != nil {
		http.Error(w, "Unable to encode city details", http.StatusInternalServerError)
	}
}

// Handler to edit a city
func editCity(w http.ResponseWriter, r *http.Request) {
	cityID := chi.URLParam(r, "city_id")

	// Fetch city details using cityID
	city := getCityByID(cityID)

	// Set response header to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode city details into JSON and return it in the response
	err := json.NewEncoder(w).Encode(city)
	if err != nil {
		http.Error(w, "Unable to encode city details", http.StatusInternalServerError)
	}
}

// Handler to save the edited city
func saveCity(w http.ResponseWriter, r *http.Request) {
	cityIDStr := chi.URLParam(r, "city_id")
	cityID, err := strconv.Atoi(cityIDStr) // Convert cityID to integer
	if err != nil {
		log.Println("Invalid city ID:", err)
		http.Error(w, "Invalid city ID", http.StatusBadRequest)
		return
	}

	// Parse form data
	city := City{
		CityID:            cityID,
		CityName:          r.FormValue("city_name"),
		Latitude:          parseFloat(r.FormValue("latitude")),
		Longitude:         parseFloat(r.FormValue("longitude")),
		TemperatureUnit:   r.FormValue("temperature_unit"),
		LatestTemperature: parseFloat(r.FormValue("latest_temperature")),
	}

	// Update the city in the database
	_, err = db.Exec(`
        UPDATE cities 
        SET city_name = ?, latitude = ?, longitude = ?, temperature_unit = ?, latest_temperature = ? 
        WHERE city_id = ?`,
		city.CityName, city.Latitude, city.Longitude, city.TemperatureUnit, city.LatestTemperature, city.CityID)

	if err != nil {
		log.Println("Error updating city:", err)
		http.Error(w, "Unable to update city", http.StatusInternalServerError)
		return
	}

	// Redirect after saving
	http.Redirect(w, r, "/cities", http.StatusSeeOther)
}

// Helper function to parse float values
func parseFloat(value string) float64 {
	floatValue, err := strconv.ParseFloat(value, 64) // Convert string to float64
	if err != nil {
		log.Println("Error parsing float:", err)
		return 0.0 // Return 0.0 if there’s an error parsing the value
	}
	return floatValue
}
