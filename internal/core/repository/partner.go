package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/raymondgitonga/matching-service/internal/core/dormain"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetPartner(ctx context.Context, partnerID int) (*dormain.Partner, error) {
	partner := &dormain.Partner{}
	query := `SELECT name, location, speciality, radius, rating FROM partner WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, partnerID)

	err := row.Scan(&partner.Name, &partner.Location, &partner.Speciality, &partner.Radius, &partner.Rating)
	if err != nil {
		return nil, fmt.Errorf("error fetching partner: %w", err)
	}

	return partner, nil
}

func (r *Repository) GetPartners(ctx context.Context, speciality string) (*[]dormain.Partner, error) {
	partner := dormain.Partner{}
	partners := make([]dormain.Partner, 0)
	param := fmt.Sprintf(`{"%s":true}`, speciality)
	query := `select name, location, speciality, radius, rating  from partner where speciality @> $1;`

	rows, err := r.db.QueryContext(ctx, query, param)
	if err != nil {
		return nil, fmt.Errorf("error querying partners: %w", err)
	}

	for rows.Next() {
		err := rows.Scan(&partner.Name, &partner.Location, &partner.Speciality, &partner.Radius, &partner.Rating)

		if err != nil {
			return nil, fmt.Errorf("error scanning partner results: %w", err)
		}
		partners = append(partners, partner)
	}

	return &partners, nil
}
