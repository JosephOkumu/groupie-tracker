package handlers

import (
	"net/http"
)

// HomeHandler handles requests to the home page of the application.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var data Info

	// Set the title for the home page.
	data.Title = "home"

	// Set the data for the home page, which includes the list of artists.
	data.Data = Artiste


	w.WriteHeader(http.StatusOK)

	// Render the base template with the home page data.
	Temp.ExecuteTemplate(w, "base.html", data)
}
