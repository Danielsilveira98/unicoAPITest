package streetmarket

import (
	"context"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
)

type repositoryReader interface {
	List(ctx context.Context, query domain.StreetMarketFilter) (*domain.StreetMarket, error)
}

type StreetMarketReader struct {
	repo repositoryReader
}

func NewReader(repo repositoryReader) *StreetMarketReader {
	return &StreetMarketReader{repo}
}

func (s *StreetMarketReader) List(ctx context.Context, query domain.StreetMarketGetInput) (*domain.StreetMarket, error) {
	filter := domain.StreetMarketFilter{
		District:     query.District,
		Region5:      query.Region5,
		Name:         query.Name,
		Neighborhood: query.Neighborhood,
	}

	sm, err := s.repo.List(ctx, filter)
	if err != nil {
		return nil, domain.ErrUnexpected
	}

	return sm, nil
}
