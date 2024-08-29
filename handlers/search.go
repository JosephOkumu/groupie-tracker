package handlers

import (
	"net/http"
	"strings"

	groupie "groupie/data"
)

// SearchHandler handles search requests for artists based on a query parameter.
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	var data Info
	var results []groupie.Artist

	// Retrieve the search query parameter from the URL.
	word := r.URL.Query().Get("query")

	for _, artist := range Artiste {
		// Check if the artist's name contains the search query (case-insensitive).
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(word)) {
			results = append(results, artist) // Add matching artist to results.
		}
	}

	if results == nil {
		ErrorPage(w, r, http.StatusNotFound, "Not found")
		return
	}

	// Set the title for the search results page.
	data.Title = "search"

	// Assign the search results to the page data.
	data.Data = results

	// Render the base template with the search results data.
	Temp.ExecuteTemplate(w, "base.html", data)
}
