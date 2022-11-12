package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/raymondgitonga/matching-service/internal/core/dormain"
)

type PartnerRepository struct {
	db *sql.DB
}

func NewPartnerRepository(db *sql.DB) *PartnerRepository {
	return &PartnerRepository{
		db: db,
	}
}

func (r *PartnerRepository) GetPartner(ctx context.Context, partnerID int) (*dormain.Partner, error) {
	partner := &dormain.Partner{}
	query := `SELECT name, location, material, radius, rating FROM partner WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, partnerID)

	err := row.Scan(&partner.Name, &partner.Location, &partner.Material, &partner.Radius, &partner.Rating)
	if err != nil {
		return nil, fmt.Errorf("error fetching partner: %w", err)
	}

	return partner, nil
}

func (r *PartnerRepository) GetPartners(ctx context.Context, material string) (*[]dormain.Partner, error) {
	partner := dormain.Partner{}
	partners := make([]dormain.Partner, 0)
	param := fmt.Sprintf(`{"%s":true}`, material)
	query := `select name, location, material, radius, rating  from partner where material @> $1;`

	rows, err := r.db.QueryContext(ctx, query, param)
	if err != nil {
		return nil, fmt.Errorf("error querying partners: %w", err)
	}

	for rows.Next() {
		err := rows.Scan(&partner.Name, &partner.Location, &partner.Material, &partner.Radius, &partner.Rating)
		if err != nil {
			return nil, fmt.Errorf("error scanning partner results: %w", err)
		}

		partners = append(partners, partner)
	}

	return &partners, nil
}
