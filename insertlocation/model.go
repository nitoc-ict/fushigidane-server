package insertlocation

type LocationData struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"long"`
	Address   string  `json:"address"`
	Scenes    string  `json:"scenes"`
}

type RecvData struct {
	Spot []LocationData `json:"spot"`
}
