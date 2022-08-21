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
	getByIDInp domain.StreetMarketFilter
	list       func(ctx context.Context, query domain.StreetMarketFilter) (*domain.StreetMarket, error)
}

func (s *stubRepositoryReader) List(
	ctx context.Context,
	query domain.StreetMarketFilter,
) (*domain.StreetMarket, error) {
	s.getByIDInp = query
	return s.list(ctx, query)
}

func TestStreetMarketReader_List(t *testing.T) {
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
		Number:        "500",
		Neighborhood:  "JARDIM SARAH",
		AddrExtraInfo: "Loren ipsum",
	}

	repoMock := &stubRepositoryReader{
		list: func(ctx context.Context, query domain.StreetMarketFilter) (*domain.StreetMarket, error) {
			return &want, nil
		},
	}

	srv := NewReader(repoMock)

	wInp := domain.StreetMarketFilter{
		District:     want.District,
		Region5:      want.Region5,
		Name:         want.Name,
		Neighborhood: want.Neighborhood,
	}

	got, _ := srv.List(context.TODO(), wInp)

	if diff := cmp.Diff(&want, got); diff != "" {
		t.Errorf("unexpected return (-want +got):\n%s", diff)
	}

	if diff := cmp.Diff(wInp, repoMock.getByIDInp); diff != "" {
		t.Errorf("unexpected filter when calls getByID (-want +got):\n%s", diff)
	}
}

func TestStreetMarketReader_List_Error(t *testing.T) {
	testCases := map[string]struct {
		wErr error
		inp  domain.StreetMarketFilter
		rErr error
		rRtn *domain.StreetMarket
	}{
		"When a unexpected error occurs in reader repository": {
			wErr: domain.ErrUnexpected,
			inp:  domain.StreetMarketFilter{},
			rErr: errSome,
			rRtn: nil,
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			repoMock := &stubRepositoryReader{
				list: func(ctx context.Context, query domain.StreetMarketFilter) (*domain.StreetMarket, error) {
					return tc.rRtn, tc.rErr
				},
			}

			srv := NewReader(repoMock)

			_, gErr := srv.List(context.TODO(), tc.inp)

			if !errors.Is(gErr, tc.wErr) {
				t.Errorf("Want error %v, got error %v", tc.wErr, gErr)
			}
		})
	}
}
