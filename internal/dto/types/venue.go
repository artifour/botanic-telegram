package types

// Venue
// This object represents a venue.
// https://core.telegram.org/bots/api#venue
type Venue struct {
	Location        Location `json:"location"`
	Title           string   `json:"title"`
	Address         string   `json:"address"`
	FoursquareId    string   `json:"foursquare_id"`
	FoursquareType  string   `json:"foursquare_type"`
	GooglePlaceId   string   `json:"google_place_id"`
	GooglePlaceType string   `json:"google_place_type"`
}
