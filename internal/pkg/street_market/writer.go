package streetmarket

import (
	"context"
	"fmt"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
)

type repositoryWriter interface {
	Create(ctx context.Context, streetMarket domain.StreetMarket) error
	Update(ctx context.Context, sm domain.StreetMarket) error
}

type uuidGenerator func() string

type StreetMarketWriter struct {
	repo  repositoryWriter
	idGen uuidGenerator
}

func NewWriter(repo repositoryWriter, idGen uuidGenerator) *StreetMarketWriter {
	return &StreetMarketWriter{repo, idGen}
}

func (s *StreetMarketWriter) Create(ctx context.Context, inp domain.StreetMarketCreateInput) (string, error) {
	if err := inp.Validate(); err != nil {
		return "", fmt.Errorf("%w", err)
	}

	sm := domain.StreetMarket{
		ID:            s.idGen(),
		Long:          inp.Long,
		Lat:           inp.Lat,
		SectCens:      inp.SectCens,
		Area:          inp.Area,
		IDdist:        inp.IDdist,
		District:      inp.District,
		IDSubTH:       inp.IDSubTH,
		SubTownHall:   inp.SubTownHall,
		Region5:       inp.Region5,
		Region8:       inp.Region8,
		Name:          inp.Name,
		Register:      inp.Register,
		Street:        inp.Street,
		Number:        inp.Number,
		Neighborhood:  inp.Neighborhood,
		AddrExtraInfo: inp.AddrExtraInfo,
	}

	err := s.repo.Create(ctx, sm)
	if err != nil {
		switch err {
		case domain.ErrForeignKeyNotExists:
			return "", domain.ErrOrgCalNotFound
		default:
			return "", domain.ErrUnexpected
		}
	}

	return sm.ID, nil
}

func (s *StreetMarketWriter) Edit(ctx context.Context, ID domain.SMID, inp domain.StreetMarketEditInput) error {
	if err := ID.Validate(); err != nil {
		return fmt.Errorf("%w", err)
	}

	sm := domain.StreetMarket{
		ID:            string(ID),
		Long:          inp.Long,
		Lat:           inp.Lat,
		SectCens:      inp.SectCens,
		Area:          inp.Area,
		IDdist:        inp.IDdist,
		District:      inp.District,
		IDSubTH:       inp.IDSubTH,
		SubTownHall:   inp.SubTownHall,
		Region5:       inp.Region5,
		Region8:       inp.Region8,
		Name:          inp.Name,
		Register:      inp.Register,
		Street:        inp.Street,
		Number:        inp.Number,
		Neighborhood:  inp.Neighborhood,
		AddrExtraInfo: inp.AddrExtraInfo,
	}

	err := s.repo.Update(ctx, sm)
	if err != nil {
		switch err {
		case domain.ErrNothingUpdated:
			return domain.ErrSMNotFound
		default:
			return domain.ErrUnexpected
		}
	}

	return nil
}
