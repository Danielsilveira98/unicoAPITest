package streetmarket

import (
	"context"
	"fmt"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
)

type repositoryReader interface {
	List(context.Context, domain.Pagination, domain.StreetMarketFilter) ([]domain.StreetMarket, error)
}

type StreetMarketReader struct {
	repo repositoryReader
}

func NewReader(repo repositoryReader) *StreetMarketReader {
	return &StreetMarketReader{repo}
}

func (s *StreetMarketReader) List(
	ctx context.Context,
	page int,
	query domain.StreetMarketFilter,
) ([]domain.StreetMarket, error) {
	const perPage = 100

	pc := domain.Pagination{}
	pc.Limit = perPage

	if page > 1 {
		pc.Offset = (page-1)*perPage + 1
	}

	filter := domain.StreetMarketFilter{
		District:     query.District,
		Region5:      query.Region5,
		Name:         query.Name,
		Neighborhood: query.Neighborhood,
	}

	ls, err := s.repo.List(ctx, pc, filter)
	if err != nil {
		err := fmt.Errorf("[repo.List] %s %w", domain.ErrUnexpected.Error(), err)
		return nil, err
	}

	return ls, nil
}
