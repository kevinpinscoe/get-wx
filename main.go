package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"os"
	"strings"
	"time"
)

// Define JSON decoder output
var v map[string]interface{}

func main() {
	baseurl := "https://api.open-meteo.com/v1/forecast"
	endpoint := config(baseurl)
	get_meteo(endpoint)
}

func config(baseurl string) string {
	// Retrieve our home directory 
	homedir, _ := os.UserHomeDir()

	// Set our config file name
	// config_dir := homedir + "/.config"
	config_dir := homedir + "/.config"
	config_file := config_dir + "/get-wx"
	os.Mkdir(config_dir, os.FileMode(0700))
	var lat string
	var long string

	// See if that file already exists
	_, err := os.Stat(config_file)
	if err != nil {
		if os.IsNotExist(err) {
	// Initial run so create a config
			lat, long = set_config(config_file)
		} else {
			log.Fatalf("Unable to stat config file %s: %s", config_file, err)
		}
	} else {
		// Not our first rodeo so get the config data
		lat, long = read_config(config_file)
	}

	// Set our URL parameters
	// In the future support further weather data to retrieve
	var url string
	lat_uri := "?latitude=" + lat
	long_uri := "&longitude=" + long
	temp_uri := "&current=temperature_2m"
	fahrenheit_uri := "&temperature_unit=fahrenheit"

	url = baseurl + lat_uri + long_uri + temp_uri + fahrenheit_uri

	return url

}

func set_config(config_file string) (string, string) {
	// in the future prompt for lat & long
	// and weather options that can be retrieved
	fmt.Println("Enter your Latitude:")
	var latitude string
	fmt.Scanln(&latitude)

	fmt.Println("Enter your Longitude:")
	var longitude string
	fmt.Scanln(&longitude)

	// Write into config file
	data := latitude + " " + longitude + "\n"
	dw := []byte(data)
	err := os.WriteFile(config_file, dw, 0600)
	if err != nil {
		log.Fatalf("Cannot write config file %s: %s", config_file, err)
	}

	// Since we now have the lat and long just return the data
	// rather than read the config file after just writing it
	return latitude, longitude
}

func read_config(config_file string) (string, string) {
	// Read in our saved config
	d, err := os.ReadFile(config_file)
	if err != nil {
		log.Fatalf("Cannot read config file %s: %s", config_file, err)
	}

	s := strings.Split(string(d), " ")
	lat := s[0]
	long := strings.ReplaceAll(s[1], "\n", "")

	return lat, long

}

func get_meteo(endpoint string) {
	// This function calls the open meteo site and retrieves
	// our weather

	// Use Client so we have a timeout
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(endpoint)
	if err != nil {
		log.Fatalf("Error making HTTP request: %s", err)
	}
	// Close the HTTP response body
	defer resp.Body.Close()

	// Also check for HTTP status
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error making HTTP request: %s", err)
		return
	}

	// Decode the JSON response or fail if it is not good
	decoded := json.NewDecoder(resp.Body)
	if err := decoded.Decode(&v); err != nil {
		log.Fatalf("Error decoding JSON result: %s", err)
		return
	}

	// We are only (currently) interested in the temp
	temp := v["current"].(map[string]interface{})["temperature_2m"]
	fmt.Printf("The current temperature is %g degrees F\n", temp)

}