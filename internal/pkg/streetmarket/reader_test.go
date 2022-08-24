package streetmarket

import (
	"context"
	"testing"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

type stubRepositoryReader struct {
	listFInp  domain.StreetMarketFilter
	listPCInp domain.Pagination
	list      func(context.Context, domain.Pagination, domain.StreetMarketFilter) ([]domain.StreetMarket, *domain.Error)
}

func (s *stubRepositoryReader) List(
	ctx context.Context,
	pc domain.Pagination,
	query domain.StreetMarketFilter,
) ([]domain.StreetMarket, *domain.Error) {
	s.listFInp = query
	s.listPCInp = pc
	return s.list(ctx, pc, query)
}

func TestStreetMarketReader_List(t *testing.T) {
	want := []domain.StreetMarket{{
		ID:            uuid.NewString(),
		Long:          -46548146,
		Lat:           -23568390,
		SectCens:      "355030885000019",
		Area:          "3550308005040",
		IDdist:        "87",
		District:      "VILA FORMOSA",
		IDSubTH:       "26",
		SubTownHall:   "ARICANDUVA",
		Region5:       "Leste",
		Region8:       "Leste 1",
		Name:          "RAPOSO TAVARES",
		Register:      "1129-0",
		Street:        "Rua dos Bobos",
		Number:        "500",
		Neighborhood:  "JARDIM SARAH",
		AddrExtraInfo: "Loren ipsum",
	}}

	repoMock := &stubRepositoryReader{
		list: func(
			ctx context.Context,
			pc domain.Pagination,
			query domain.StreetMarketFilter,
		) ([]domain.StreetMarket, *domain.Error) {
			return want, nil
		},
	}

	srv := NewReader(repoMock)

	wInp := domain.StreetMarketFilter{
		District:     want[0].District,
		Region5:      want[0].Region5,
		Name:         want[0].Name,
		Neighborhood: want[0].Neighborhood,
	}

	t.Run("When page is 0", func(t *testing.T) {
		got, _ := srv.List(context.TODO(), 0, wInp)

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("unexpected return (-want +got):\n%s", diff)
		}

		if diff := cmp.Diff(wInp, repoMock.listFInp); diff != "" {
			t.Errorf("unexpected filter when calls list (-want +got):\n%s", diff)
		}

		wPc := domain.Pagination{
			Offset: 0,
			Limit:  100,
		}
		if diff := cmp.Diff(wPc, repoMock.listPCInp); diff != "" {
			t.Errorf("unexpected page chain when calls list (-want +got):\n%s", diff)
		}
	})

	t.Run("When page is 1", func(t *testing.T) {
		got, _ := srv.List(context.TODO(), 1, wInp)

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("unexpected return (-want +got):\n%s", diff)
		}

		if diff := cmp.Diff(wInp, repoMock.listFInp); diff != "" {
			t.Errorf("unexpected filter when calls list (-want +got):\n%s", diff)
		}

		wPc := domain.Pagination{
			Offset: 0,
			Limit:  100,
		}
		if diff := cmp.Diff(wPc, repoMock.listPCInp); diff != "" {
			t.Errorf("unexpected page chain when calls list (-want +got):\n%s", diff)
		}
	})

	t.Run("When page is less than 2", func(t *testing.T) {
		page := 2
		wPc := domain.Pagination{
			Offset: 101,
			Limit:  100,
		}

		got, _ := srv.List(context.TODO(), page, wInp)

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("unexpected return (-want +got):\n%s", diff)
		}

		if diff := cmp.Diff(wInp, repoMock.listFInp); diff != "" {
			t.Errorf("unexpected filter when calls list (-want +got):\n%s", diff)
		}
		if diff := cmp.Diff(wPc, repoMock.listPCInp); diff != "" {
			t.Errorf("unexpected page chain when calls list (-want +got):\n%s", diff)
		}
	})
}

func TestStreetMarketReader_List_Error(t *testing.T) {
	testCases := map[string]struct {
		wErr domain.KindError
		inp  domain.StreetMarketFilter
		rErr *domain.Error
	}{
		"When a unexpected error occurs in reader repository": {
			wErr: domain.UnexpectedErrKd,
			inp:  domain.StreetMarketFilter{},
			rErr: &domain.Error{Kind: domain.UnexpectedErrKd},
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			repoMock := &stubRepositoryReader{
				list: func(
					ctx context.Context,
					pc domain.Pagination,
					query domain.StreetMarketFilter,
				) ([]domain.StreetMarket, *domain.Error) {
					return nil, tc.rErr
				},
			}

			srv := NewReader(repoMock)

			_, gErr := srv.List(context.TODO(), 0, tc.inp)

			if gErr.Kind != tc.wErr {
				t.Errorf("Want error kind %v, got error %v", tc.wErr, gErr.Kind)
			}
		})
	}
}
