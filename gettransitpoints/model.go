package gettransitpoints

type TransitPoint struct {
	Id        int64
	Address   string
	Label     string
	Latitude  float64
	Longitude float64
}

type RouteRequest struct {
	Origin      string   `json:"origin"`
	Destination string   `json:"destination"`
	Scenes      []string `json:"scenes"`
}
