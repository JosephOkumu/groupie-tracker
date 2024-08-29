package groupie

// Artist represents information about a musical artist or band.
type Artist struct {
	Id           int      `json:"id"`           
	Image        string   `json:"image"`        
	Name         string   `json:"name"`         
	Members      []string `json:"members"`      
	CreationDate int      `json:"creationDate"` 
	FirstAlbum   string   `json:"firstAlbum"`   
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`    
}

// Indexx represents a location with its ID, associated locations, and related dates.
type Indexx struct {
	Id        int      `json:"id"`        
	Locations []string `json:"locations"` 
	Dates     string   `json:"dates"`     
}

// Location represents the locations data from the API, containing multiple location indices.
type Location struct {
	Index []Indexx `json:"index"` 
}

// Indexxx represents a date with its ID and associated dates.
type Indexxx struct {
	Id    int      `json:"id"`   
	Dates []string `json:"dates"`
}

// Dates represents the concert dates data from the API, containing multiple date indices.
type Dates struct {
	Index []Indexxx `json:"index"` 
}

// Indexxxx represents the relation between dates and locations with an ID and a map of dates to locations.
type Indexxxx struct {
	Id             int                 `json:"id"`             
	DatesLocations map[string][]string `json:"datesLocations"` 
}

// Relations represents the relations data from the API, containing multiple relation indices.
type Relations struct {
	Index []Indexxxx `json:"index"`
}
