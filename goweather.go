package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const apiKey = "YOUR API KEY HERE"

type WeatherResponse struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: goweather <cityname>")
		os.Exit(1)
	}

	cityName := os.Args[1]

	// Make API request
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", apiKey, cityName)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Error: %s\n", response.Status)
		os.Exit(1)
	}

	// Parse JSON response
	var weather WeatherResponse
	if err := json.NewDecoder(response.Body).Decode(&weather); err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		os.Exit(1)
	}

	// Display weather information
	fmt.Printf("Weather in %s:\n", cityName)
	fmt.Printf("Temperature: %.1f°C\n", weather.Current.TempC)
}
