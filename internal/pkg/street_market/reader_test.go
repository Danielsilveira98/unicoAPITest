package streetmarket

import (
	"context"
	"errors"
	"testing"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

type stubRepositoryReader struct {
	getByIDQInp  domain.StreetMarketGetQuery
	getByIDIDInp string
	getByID      func(ctx context.Context, ID string, query domain.StreetMarketGetQuery) (*domain.StreetMarket, error)
}

func (s *stubRepositoryReader) GetByID(
	ctx context.Context,
	ID string,
	query domain.StreetMarketGetQuery,
) (*domain.StreetMarket, error) {
	s.getByIDIDInp = ID
	s.getByIDQInp = query
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
			District:     want.District,
			Region5:      want.Region5,
			Name:         want.Name,
			Neighborhood: want.Neighborhood,
		},
	}

	got, _ := srv.Get(context.TODO(), inp)

	if diff := cmp.Diff(&want, got); diff != "" {
		t.Errorf("unexpected return (-want +got):\n%s", diff)
	}

	if inp.ID != repoMock.getByIDIDInp {
		t.Errorf("expect calls getByID with ID %s, got %s", inp.ID, repoMock.getByIDIDInp)
	}

	if diff := cmp.Diff(inp.Query, repoMock.getByIDQInp); diff != "" {
		t.Errorf("unexpected query when calls getByID (-want +got):\n%s", diff)
	}
}

func TestStreetMarketReader_Get_Error(t *testing.T) {
	testCases := map[string]struct {
		wErr error
		inp  domain.StreetMarketGetInput
		rErr error
		rRtn *domain.StreetMarket
	}{
		"When input is invalid": {
			wErr: domain.ErrInpValidation,
			inp:  domain.StreetMarketGetInput{ID: "invalid"},
			rErr: nil,
			rRtn: nil,
		},
		"When a unexpected error occurs in reader repository": {
			wErr: domain.ErrUnexpected,
			inp:  domain.StreetMarketGetInput{ID: uuid.NewString(), Query: domain.StreetMarketGetQuery{}},
			rErr: errSome,
			rRtn: nil,
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			repoMock := &stubRepositoryReader{
				getByID: func(ctx context.Context, ID string, query domain.StreetMarketGetQuery) (*domain.StreetMarket, error) {
					return tc.rRtn, tc.rErr
				},
			}

			srv := NewReader(repoMock)

			_, gErr := srv.Get(context.TODO(), tc.inp)

			if !errors.Is(gErr, tc.wErr) {
				t.Errorf("Want error %v, got error %v", tc.wErr, gErr)
			}
		})
	}
}
