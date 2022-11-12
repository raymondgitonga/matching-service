package dormain

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

type Partner struct {
	Name       string
	Location   string
	Speciality []byte
	Radius     int
	Rating     float64
}

type PartnerDTO struct {
	Name       string   `json:"name"`
	Location   string   `json:"location"`
	Speciality []string `json:"speciality"`
	Radius     int      `json:"radius,omitempty"`
	Rating     float64  `json:"rating"`
	Distance   int      `json:"distance,omitempty"`
}
