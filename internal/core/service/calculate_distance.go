package service

import (
	"github.com/raymondgitonga/matching-service/internal/core/dormain"
	"math"
)

func NewCoordinates(latitude float64, longitude float64) *dormain.Coordinates {
	return &dormain.Coordinates{
		Latitude:  latitude,
		Longitude: longitude,
	}
}
func Distance(customerCoordinate dormain.Coordinates, partnerCoordinate dormain.Coordinates) float64 {
	var radius = 6371000.0
	distance := radius * 2 * math.Asin(math.Sqrt(0.5-math.Cos((partnerCoordinate.Latitude-customerCoordinate.Latitude)*math.Pi/180)/2+math.Cos(customerCoordinate.Latitude*math.Pi/180)*math.Cos(partnerCoordinate.Latitude*math.Pi/180)*(1-math.Cos((partnerCoordinate.Longitude-customerCoordinate.Longitude)*math.Pi/180))/2))

	return distance / 1000
}
