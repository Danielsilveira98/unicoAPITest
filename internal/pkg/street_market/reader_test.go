package streetmarket

import (
	"context"
	"testing"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
	"github.com/google/uuid"
)

type stubRepositoryReader struct {
	getByIDInp domain.StreetMarketGetInput
	getByID    func(ctx context.Context, ID string, query domain.StreetMarketGetQuery) (*domain.StreetMarket, error)
}

func (s *stubRepositoryReader) GetByID(
	ctx context.Context,
	ID string,
	query domain.StreetMarketGetQuery,
) (*domain.StreetMarket, error) {
	s.getByIDInp = domain.StreetMarketGetInput{ID: ID, Query: query}
	return s.getByID(ctx, ID, query)
}

func TestStreetMarketReader_Get(t *testing.T) {
	want := domain.StreetMarket{
		ID:            uuid.NewString(),
		Long:          -46548146,
		Lat:           -23568390,
		SectCens:      "355030885000019",
		Area:          "3550308005040",
		IDdist:        87,
		District:      "VILA FORMOSA",
		IDSubTH:       26,
		SubTownHall:   "ARICANDUVA",
		Region5:       "Leste",
		Region8:       "Leste 1",
		Name:          "RAPOSO TAVARES",
		Register:      "1129-0",
		Street:        "Rua dos Bobos",
		Number:        500,
		Neighborhood:  "JARDIM SARAH",
		AddrExtraInfo: "Loren ipsum",
	}

	repoMock := &stubRepositoryReader{
		getByID: func(ctx context.Context, ID string, query domain.StreetMarketGetQuery) (*domain.StreetMarket, error) {
			return &want, nil
		},
	}

	srv := NewReader(repoMock)

	inp := domain.StreetMarketGetInput{
		ID: want.ID,
		Query: domain.StreetMarketGetQuery{
			District: want.District,
		},
	}

	srv.Get(context.TODO(), inp)

}
