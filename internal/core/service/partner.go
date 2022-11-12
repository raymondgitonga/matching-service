package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/raymondgitonga/matching-service/internal/core/dormain"
)

type PartnerDetails struct {
	partnerID int
	db        *sql.DB
}

func NewPartnerDetails(partnerID int, db *sql.DB) *PartnerDetails {
	return &PartnerDetails{
		partnerID: partnerID,
		db:        db,
	}
}

type Repository interface {
	GetPartner(ctx context.Context, partnerID int) (*dormain.Partner, error)
}

func (p *PartnerDetails) GetPartnerDetails(ctx context.Context, repo Repository) (*dormain.PartnerDTO, error) {
	partner, err := repo.GetPartner(ctx, p.partnerID)
	specialityMap := make(map[string]bool)
	speciality := make([]string, 0)

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
