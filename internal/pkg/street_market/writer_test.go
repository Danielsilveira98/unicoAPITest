package streetmarket

import (
	"context"
	"errors"
	"testing"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
	"github.com/google/go-cmp/cmp"
)

type stubRepositoryWriter struct {
	createSMInp domain.StreetMarket
	create      func(ctx context.Context, sm domain.StreetMarket) error
}

func (s *stubRepositoryWriter) Create(
	ctx context.Context,
	sm domain.StreetMarket,
) error {
	s.createSMInp = sm
	return s.create(ctx, sm)
}

func TestStreetMarketWriter_Create(t *testing.T) {
	want := "d00443e8-160d-4099-8a93-442a183be369"

	repoMock := &stubRepositoryWriter{
		create: func(ctx context.Context, sm domain.StreetMarket) error {
			return nil
		},
	}

	idGenMock := func() string {
		return want
	}

	inp := domain.StreetMarketCreateInput{
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

	wSM := domain.StreetMarket{
		ID:            want,
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

	srv := NewWriter(repoMock, idGenMock)

	got, _ := srv.Create(context.TODO(), inp)

	if got != want {
		t.Errorf("unexpected return want id %s, got %s", want, got)
	}

	if diff := cmp.Diff(wSM, repoMock.createSMInp); diff != "" {
		t.Errorf("unexpected street market when calls create (-want +got):\n%s", diff)
	}
}

var errSome = errors.New("some error")

func TestStreetMarketWriter_Create_Error(t *testing.T) {
	validInp := domain.StreetMarketCreateInput{
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

	testCases := map[string]struct {
		rErr  error
		IDGen string
		inp   domain.StreetMarketCreateInput
		wErr  error
	}{
		"When input is invalid": {
			wErr:  domain.ErrInpValidation,
			IDGen: "d00443e8-160d-4099-8a93-442a183be369",
			inp:   domain.StreetMarketCreateInput{},
		},
		"When organization calendar not exists": {
			wErr:  domain.ErrOrgCalNotFound,
			rErr:  domain.ErrForeignKeyNotExists,
			IDGen: "f2c69026-8d93-4402-8854-fec51b06e1bb",
			inp:   validInp,
		},
		"When unexpected erro occurs in writer repository": {
			wErr:  domain.ErrUnexpected,
			rErr:  errSome,
			inp:   validInp,
			IDGen: "70bb2026-9e6a-4dad-9f86-99dbddf3a087",
		},
	}
	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			repoMock := &stubRepositoryWriter{
				create: func(ctx context.Context, sm domain.StreetMarket) error {
					return tc.rErr
				},
			}

			idGenMock := func() string {
				return tc.IDGen
			}

			srv := NewWriter(repoMock, idGenMock)

			_, gErr := srv.Create(context.TODO(), tc.inp)

			if !errors.Is(gErr, tc.wErr) {
				t.Errorf("Want error %v, got error %v", tc.wErr, gErr)
			}
		})
	}
}
