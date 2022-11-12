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

type Speciality map[string]bool
