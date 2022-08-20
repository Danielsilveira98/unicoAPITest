package streetmarket

import (
	"context"
	"fmt"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
)

type repositoryReader interface {
	GetByID(ctx context.Context, ID string, query domain.StreetMarketGetQuery) (*domain.StreetMarket, error)
}

type StreetMarketReader struct {
	repo repositoryReader
}

func NewReader(repo repositoryReader) *StreetMarketReader {
	return &StreetMarketReader{repo}
}

func (s *StreetMarketReader) Get(ctx context.Context, inp domain.StreetMarketGetInput) (*domain.StreetMarket, error) {
	if err := inp.Validate(); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	sm, err := s.repo.GetByID(ctx, inp.ID, inp.Query)
	if err != nil {
		return nil, domain.ErrUnexpected
	}

	return sm, nil
}
