package handlers

import (
	"net/http"

	groupie "groupie/data"
)

// DateData holds the artist details and their concert dates for rendering.
type DateData struct {
	Details groupie.Artist  
	Dt      groupie.Indexxx 
}

// DatesHandler handles the HTTP request for displaying concert dates for an artist.
func DatesHandler(w http.ResponseWriter, r *http.Request) {
	var data Info
	var details DateData

	// Retrieve the artist ID from the query parameters.
	id := GetId(w, r)
	if id <= 0 {
		return 
	}

	// Assign the artist details and their concert dates to the DateData struct.
	details.Details = Artiste[id-1]    
	details.Dt = Tarehe.Index[id-1]   

	// Set the title and data for the template.
	data.Title = "dates"
	data.Data = details

	// Set HTTP status to OK and render the template with the data.
	w.WriteHeader(http.StatusOK)
	Temp.ExecuteTemplate(w, "base.html", data)
}
