package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
)

type StreetMarketRepository struct {
	db *sql.DB
}

var ()

func NewStreetMarketRepository(db *sql.DB) *StreetMarketRepository {
	return &StreetMarketRepository{db}
}

func (r *StreetMarketRepository) List(
	ctx context.Context,
	query domain.StreetMarketFilter,
) (*domain.StreetMarket, error) {
	return nil, nil
}

func (r *StreetMarketRepository) Create(ctx context.Context, streetMarket domain.StreetMarket) error {
	return nil
}

func (r *StreetMarketRepository) Update(ctx context.Context, sm domain.StreetMarket) error {
	return nil
}

func (r *StreetMarketRepository) DeleteByID(ctx context.Context, ID string) error {
	q := "DELETE FROM street_market WHERE id = $1"

	qr, err := r.db.ExecContext(ctx, q, ID)
	if err != nil {
		return fmt.Errorf("%w", domain.ErrUnexpected)
	}

	ra, err := qr.RowsAffected()
	if err != nil {
		return fmt.Errorf("%w", domain.ErrUnexpected)
	}

	if ra < 1 {
		return domain.ErrNothingDeleted
	}

	return nil
}
