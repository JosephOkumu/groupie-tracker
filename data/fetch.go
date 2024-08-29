package groupie

import (
	"encoding/json"
	"net/http"
)

// Response performs an HTTP GET request to the specified API endpoint and returns the response.
func Response(str string) (*http.Response, error) {
	
	url := "https://groupietrackers.herokuapp.com/api/" + str

	// Perform the HTTP GET request.
	response, err := http.Get(url)
	if err != nil {
		return nil, err 
	}
	return response, nil 
}

// FetchArtist fetches the artist data from the API and returns a slice of Artist structs.
func FetchArtist() ([]Artist, error) {
	var artist []Artist
	response, err := Response("artists")
	if err != nil {
		return nil, err 
	}
	defer response.Body.Close()

	// Create a JSON decoder to read from the response body.
	decode := json.NewDecoder(response.Body)
	// Decode the JSON data into the artist slice.
	jsErr := decode.Decode(&artist)
	if jsErr != nil {
		return nil, jsErr
	}
	return artist, nil 
}

// FetchLocation fetches the location data from the API and returns a Location struct.
func FetchLocation() (Location, error) {
	var locations Location
	response, err := Response("locations")
	if err != nil {
		return locations, err 
	}
	defer response.Body.Close() 

	// Create a JSON decoder to read from the response body.
	decode := json.NewDecoder(response.Body)
	// Decode the JSON data into the Location struct.
	jsErr := decode.Decode(&locations)
	if jsErr != nil {
		return locations, jsErr 
	}
	return locations, nil 
}

// FetchDates fetches the dates data from the API and returns a Dates struct.
func FetchDates() (Dates, error) {
	var dates Dates
	response, err := Response("dates")
	if err != nil {
		return dates, err 
	}
	defer response.Body.Close() 

	decode := json.NewDecoder(response.Body)

	jsErr := decode.Decode(&dates)
	if jsErr != nil {
		return dates, jsErr
	}
	return dates, nil 
}

// FetchRelation fetches the relation data from the API and returns a Relations struct.
func FetchRelation() (Relations, error) {
	var relation Relations
	response, err := Response("relation")
	if err != nil {
		return relation, err 
	}
	defer response.Body.Close() 

	decode := json.NewDecoder(response.Body)
	// Decode the JSON data into the Relations struct.
	jsErr := decode.Decode(&relation)
	if jsErr != nil {
		return relation, jsErr 
	}
	return relation, nil 
}