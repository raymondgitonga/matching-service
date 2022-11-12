package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/raymondgitonga/matching-service/internal/core/dormain"
	"sort"
	"strconv"
	"strings"
)

type PartnerService struct {
	repo Repository
}

func NewPartnerService(repo Repository) *PartnerService {
	return &PartnerService{
		repo: repo,
	}
}

type Repository interface {
	GetPartner(ctx context.Context, partnerID int) (*dormain.Partner, error)
	GetPartners(ctx context.Context, speciality string) (*[]dormain.Partner, error)
}

func (p *PartnerService) GetPartnerDetails(ctx context.Context, partnerID int) (*dormain.PartnerDTO, error) {
	specialityMap := make(map[string]bool)
	speciality := make([]string, 0)

	partner, err := p.repo.GetPartner(ctx, partnerID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(partner.Speciality, &specialityMap)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling speciality: %w", err)
	}

	for key, val := range specialityMap {
		if val {
			speciality = append(speciality, key)
		}
	}

	partnerLocation := partner.Location[1 : len(partner.Location)-1]

	partnerDTO := &dormain.PartnerDTO{
		Name:       partner.Name,
		Location:   partnerLocation,
		Speciality: speciality,
		Radius:     partner.Radius,
		Rating:     partner.Rating,
	}

	return partnerDTO, nil
}

func (p *PartnerService) GetMatchingPartners(ctx context.Context, speciality string, lat float64, lon float64) (*[]dormain.PartnerDTO, error) {
	partners, err := p.repo.GetPartners(ctx, speciality)
	if err != nil {
		return nil, err
	}

	return sortAndFilterPartners(partners, lat, lon)
}

func ComputeDistance(partnerLocation string, customerLat float64, customerLon float64) (int, error) {
	location := strings.Split(partnerLocation, ",")
	partnerLat, err := strconv.ParseFloat(location[0], 64)
	if err != nil {
		return -1, fmt.Errorf("error parsing partnerLat: %w", err)
	}

	partnerLon, err := strconv.ParseFloat(location[1], 64)
	if err != nil {
		return -1, fmt.Errorf("error parsing partnerLon: %w", err)
	}

	customerCoordinates := NewCoordinates(customerLat, customerLon)
	partnerCoordinates := NewCoordinates(partnerLat, partnerLon)
	distance := Distance(*customerCoordinates, *partnerCoordinates)

	return int(distance), nil
}

func sortAndFilterPartners(partners *[]dormain.Partner, lat float64, lon float64) (*[]dormain.PartnerDTO, error) {
	specialityMap := make(map[string]bool)
	partnersDTO := make([]dormain.PartnerDTO, 0)
	for i := range *partners {
		specialities := make([]string, 0)
		current := (*partners)[i]

		err := json.Unmarshal(current.Speciality, &specialityMap)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling speciality: %w", err)
		}

		for key, val := range specialityMap {
			if val {
				specialities = append(specialities, key)
			}
		}

		partnerLocation := current.Location[1 : len(current.Location)-1]

		//Check if customer's distance is within partners' range
		distance, err := ComputeDistance(partnerLocation, lat, lon)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling speciality: %w", err)
		}

		if distance >= 0 && distance <= current.Radius {
			partnerDTO := dormain.PartnerDTO{
				Name:       current.Name,
				Location:   partnerLocation,
				Speciality: specialities,
				Radius:     current.Radius,
				Distance:   distance,
				Rating:     current.Rating,
			}
			partnersDTO = append(partnersDTO, partnerDTO)
		}
	}

	sort.Slice(partnersDTO, func(i, j int) bool {
		return partnersDTO[i].Rating > partnersDTO[j].Rating
	})

	return &partnersDTO, nil
}
