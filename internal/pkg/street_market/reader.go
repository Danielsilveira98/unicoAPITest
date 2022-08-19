package streetmarket

import (
	"context"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
)

type repositoryReader interface {
	GetByID(ctx context.Context, ID string, query domain.StreetMarketGetQuery) (*domain.StreetMarket, error)
}

type loggerReaderAble interface{}

type StreetMarketReader struct {
	repo repositoryReader
	// logger loggerReaderAble
}

func NewReader(repo repositoryReader /* , logger loggerReaderAble */) *StreetMarketReader {
	return &StreetMarketReader{repo /* , logger */}
}

func (s *StreetMarketReader) Get(ctx context.Context, inp domain.StreetMarketGetInput) (*domain.StreetMarket, error) {
	// if err := inp.Validate(); err != nil {
	// 	return nil, fmt.Errorf("%w", err)
	// }

	return nil, nil
}
