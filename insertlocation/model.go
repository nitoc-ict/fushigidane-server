package insertlocation

type LocationData struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   string  `json:"address"`
	Scenes    string  `json:"scenes"`
}

type RecvData struct {
	Spot []LocationData `json:"spot"`
}
