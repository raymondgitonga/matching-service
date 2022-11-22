package service

import (
	"math"
)

type Coordinates struct {
	latitude  float64
	longitude float64
}

func NewCoordinates(latitude float64, longitude float64) *Coordinates {
	return &Coordinates{
		latitude:  latitude,
		longitude: longitude,
	}
}

// Distance :Calculates distance between two coordinates
func Distance(customerCoordinate Coordinates, partnerCoordinate Coordinates) float64 {
	var radius = 6371000.0
	distance := radius * 2 * math.Asin(math.Sqrt(0.5-math.Cos((partnerCoordinate.latitude-customerCoordinate.latitude)*math.Pi/180)/2+math.Cos(customerCoordinate.latitude*math.Pi/180)*math.Cos(partnerCoordinate.latitude*math.Pi/180)*(1-math.Cos((partnerCoordinate.longitude-customerCoordinate.longitude)*math.Pi/180))/2))

	return distance / 1000
}
