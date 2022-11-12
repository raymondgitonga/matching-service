package dormain

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

type Partner struct {
	Name     string
	Location string
	Material []byte
	Radius   int
	Rating   float64
}

type CustomerRequest struct {
	Material  string  `json:"material"`
	Lat       float64 `json:"lat"`
	Long      float64 `json:"long"`
	FloorSize float64 `json:"floor_size"`
	Phone     string  `json:"phone"`
}

type PartnerResponse struct {
	Result  []PartnerDTO `json:"result"`
	Error   bool         `json:"error"`
	Message string       `json:"message"`
}

type PartnerDTO struct {
	Name     string   `json:"name"`
	Location string   `json:"location"`
	Material []string `json:"material"`
	Radius   int      `json:"radius,omitempty"`
	Rating   float64  `json:"rating"`
	Distance int      `json:"distance,omitempty"`
}
