package handlers

import (
	"net/http"

	groupie "groupie/data"
)

// RelData holds the data for an artist and their relations.
type RelData struct {
	Details groupie.Artist 
	Rel     groupie.Indexxxx 
}

// RelationsHandler handles requests to the relations page for an artist.
func RelationsHandler(w http.ResponseWriter, r *http.Request) {
	var data Info
	var detail RelData

	
	id := GetId(w, r)
	if id <= 0 {
		return  
	}

	
	detail.Details = Artiste[id-1]
	detail.Rel = Rapports.Index[id-1]

	data.Title = "relations"

	
	data.Data = detail

	w.WriteHeader(http.StatusOK)
	
	Temp.ExecuteTemplate(w, "base.html", data)
}
