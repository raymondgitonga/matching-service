package httpclient

type Partner struct {
	Result  []Result `json:"result"`
	Error   bool     `json:"error"`
	Message string   `json:"message"`
}

type Result struct {
	Name     string   `json:"name"`
	Location string   `json:"location"`
	Material []string `json:"material"`
	Radius   int      `json:"radius"`
	Rating   float64  `json:"rating"`
}
