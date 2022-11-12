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
