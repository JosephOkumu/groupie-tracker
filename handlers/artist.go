package handlers

import (
	"net/http"

	groupie "groupie/data"
)

// Data holds the artist and relation information for rendering.
type Data struct {
	Artist   groupie.Artist    
	Relation groupie.Indexxxx   
}

// ArtistHandler handles the HTTP request for displaying artist details.

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	var data Info
	var relate Data

	// Retrieve the artist ID from the query parameters.
	id := GetId(w, r)
	if id <= 0 {
		return 
	}

	// Assign the artist and its relation information to the Data struct.
	relate.Artist = Artiste[id-1]      
	relate.Relation = Rapports.Index[id-1] 

	// Set the title and data for the template.
	data.Title = "Artist"
	data.Data = relate

	// Set HTTP status to OK and render the template with the data.
	w.WriteHeader(http.StatusOK)
	Temp.ExecuteTemplate(w, "base.html", data)
}
