package streetmarket

import (
	"context"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
)

type repositoryWriter interface {
	Create(ctx context.Context, streetMarket domain.StreetMarket) *domain.Error
	Update(ctx context.Context, sm domain.StreetMarket) *domain.Error
}

type uuidGenerator func() string

type StreetMarketWriter struct {
	repo  repositoryWriter
	idGen uuidGenerator
}

func NewWriter(repo repositoryWriter, idGen uuidGenerator) *StreetMarketWriter {
	return &StreetMarketWriter{repo, idGen}
}

func (s *StreetMarketWriter) Create(ctx context.Context, inp domain.StreetMarketCreateInput) (string, *domain.Error) {
	if err := inp.Validate(); err != nil {
		return "", &domain.Error{
			Kind:     domain.InpValidationErrKd,
			Msg:      "Invalid input",
			Previous: err,
		}
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
		return "", &domain.Error{Kind: domain.UnexpectedErrKd, Msg: "Unexpected error when create", Previous: err}
	}

	return sm.ID, nil
}

func (s *StreetMarketWriter) Edit(ctx context.Context, ID domain.SMID, inp domain.StreetMarketEditInput) *domain.Error {
	if err := ID.Validate(); err != nil {
		return &domain.Error{
			Kind:     domain.InpValidationErrKd,
			Msg:      "Invalid input",
			Previous: err,
		}
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
		switch err.Kind {
		case domain.NothingUpdatedErrKd:
			return &domain.Error{Kind: domain.SMNotFoundErrKd, Msg: "Entity not exists", Previous: err}
		default:
			return &domain.Error{Kind: domain.UnexpectedErrKd, Msg: "Unexpected error when edit", Previous: err}
		}
	}

	return nil
}
