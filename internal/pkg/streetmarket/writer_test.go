package streetmarket

import (
	"context"
	"errors"
	"testing"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
	"github.com/google/go-cmp/cmp"
)

var errSome = errors.New("some error")

type stubRepositoryWriter struct {
	createSMInp domain.StreetMarket
	create      func(ctx context.Context, sm domain.StreetMarket) error
	updateInp   domain.StreetMarket
	update      func(ctx context.Context, sm domain.StreetMarket) error
}

func (s *stubRepositoryWriter) Create(
	ctx context.Context,
	sm domain.StreetMarket,
) error {
	s.createSMInp = sm
	return s.create(ctx, sm)
}

func (s *stubRepositoryWriter) Update(ctx context.Context, sm domain.StreetMarket) error {
	s.updateInp = sm
	return s.update(ctx, sm)
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

func TestStreetMarketWriter_Create_Error(t *testing.T) {
	validInp := domain.StreetMarketCreateInput{
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
	}

	testCases := map[string]struct {
		rErr  error
		IDGen string
		inp   domain.StreetMarketCreateInput
		wErr  error
	}{
		"When unexpected erro occurs in writer repository": {
			wErr:  domain.ErrUnexpected,
			rErr:  errSome,
			inp:   validInp,
			IDGen: "70bb2026-9e6a-4dad-9f86-99dbddf3a087",
		},
		"When input is invalid": {
			wErr: domain.ErrInpValidation,
			inp:  domain.StreetMarketCreateInput{},
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

func TestStreetMarketWriter_Edit(t *testing.T) {
	repoMock := &stubRepositoryWriter{
		update: func(ctx context.Context, sm domain.StreetMarket) error {
			return nil
		},
	}

	idGenMock := func() string { return "" }

	srv := NewWriter(repoMock, idGenMock)

	var id domain.SMID = "07468c29-cd01-414d-adcb-68282eb94d9a"
	editInp := domain.StreetMarketEditInput{
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
	}

	err := srv.Edit(context.TODO(), id, editInp)

	if err != nil {
		t.Errorf("expect return nil, got %v", err)
	}

	want := domain.StreetMarket{
		ID:            string(id),
		Long:          editInp.Long,
		Lat:           editInp.Lat,
		SectCens:      editInp.SectCens,
		Area:          editInp.Area,
		IDdist:        editInp.IDdist,
		District:      editInp.District,
		IDSubTH:       editInp.IDSubTH,
		SubTownHall:   editInp.SubTownHall,
		Region5:       editInp.Region5,
		Region8:       editInp.Region8,
		Name:          editInp.Name,
		Register:      editInp.Register,
		Street:        editInp.Street,
		Number:        editInp.Number,
		Neighborhood:  editInp.Neighborhood,
		AddrExtraInfo: editInp.AddrExtraInfo,
	}

	if diff := cmp.Diff(want, repoMock.updateInp); diff != "" {
		t.Errorf("unexpected street market when calls edit (-want +got):\n%s", diff)
	}
}

func TestStreetMarketWriter_Edit_Error(t *testing.T) {
	testCases := map[string]struct {
		rErr error
		wErr error
		id   domain.SMID
	}{
		"When entity not exists": {
			rErr: domain.ErrNothingUpdated,
			wErr: domain.ErrSMNotFound,
			id:   "51557ef2-dfe8-485d-90e0-c7adf4e59581",
		},
		"When a unexpected error occurs in repository": {
			rErr: errSome,
			wErr: domain.ErrUnexpected,
			id:   "c882edc1-c1f3-4b20-b8f6-36156d99bc48",
		},
		"When id is invalid": {
			wErr: domain.ErrInpValidation,
			id:   "invalid",
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			repoMock := &stubRepositoryWriter{
				update: func(ctx context.Context, sm domain.StreetMarket) error {
					return tc.rErr
				},
			}

			idGenMock := func() string { return "" }

			srv := NewWriter(repoMock, idGenMock)

			gErr := srv.Edit(context.TODO(), tc.id, domain.StreetMarketEditInput{})

			if !errors.Is(gErr, tc.wErr) {
				t.Errorf("Want error %v, got error %v", tc.wErr, gErr)
			}
		})
	}
}
