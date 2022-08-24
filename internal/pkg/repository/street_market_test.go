package repository

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
	"github.com/google/go-cmp/cmp"
)

var errSome = errors.New("some error")

func TestStreetMarketRepository_Delete(t *testing.T) {
	id := "84713a81-0e31-4c14-a62f-7e1f67bc526d"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM street_market WHERE .+").WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewStreetMarketRepository(db)

	if err := repo.DeleteByID(context.TODO(), id); err != nil {
		t.Errorf("expect return nil, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestStreetMarketRepository_Delete_Error(t *testing.T) {
	testCases := map[string]struct {
		wErr   domain.KindError
		mErr   error
		notUpd bool
		id     string
	}{
		"When unexpected error occurs": {
			wErr: domain.UnexpectedErrKd,
			mErr: errSome,
			id:   "822d08fb-6dbd-457b-a467-36ee3a136b13",
		},
		"When update nothing": {
			notUpd: true,
			wErr:   domain.NothingDeletedErrKd,
			id:     "65ea9603-c3bb-4686-9700-8881c8a89374",
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("%v", err)
			}
			defer db.Close()

			if tc.notUpd {
				mock.ExpectExec("DELETE FROM street_market WHERE .+").WithArgs(tc.id).WillReturnResult(sqlmock.NewResult(1, 0))
			} else {
				mock.ExpectExec("DELETE FROM street_market WHERE .+").WithArgs(tc.id).WillReturnError(tc.mErr)
			}

			repo := NewStreetMarketRepository(db)

			gErr := repo.DeleteByID(context.TODO(), tc.id)

			if gErr.Kind != tc.wErr {
				t.Errorf("Want error kind %v, got error %v", tc.wErr, gErr.Kind)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestStreetMarketRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer db.Close()

	inp := domain.StreetMarket{
		ID:           "944ec25d-aac4-4c35-8301-6b35e0d7c05f",
		Name:         "RAPOSO TAVARES",
		Register:     "1129-0",
		Street:       "Rua dos Bobos",
		Number:       "500",
		Neighborhood: "JARDIM SARAH",
	}

	mock.ExpectExec(
		"INSERT INTO street_market (id,name,register,street,number,neighborhood) VALUES ($1,$2,$3,$4,$5,$6)",
	).WithArgs(
		inp.ID,
		inp.Name,
		inp.Register,
		inp.Street,
		inp.Number,
		inp.Neighborhood,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewStreetMarketRepository(db)

	if err := repo.Create(context.TODO(), inp); err != nil {
		t.Errorf("expect return nil, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestStreetMarketRepository_Create_Error(t *testing.T) {
	testCases := map[string]struct {
		createNothing bool
		wErr          domain.KindError
		mErr          error
	}{
		"When unexpected error occurs": {
			wErr: domain.UnexpectedErrKd,
			mErr: errSome,
		},
		"When create nothing": {
			createNothing: true,
			wErr:          domain.NothingCreatedErrKd,
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("%v", err)
			}
			defer db.Close()

			if tc.createNothing {
				mock.ExpectExec(".+").WillReturnResult(sqlmock.NewResult(1, 0))
			} else {
				mock.ExpectExec(".+").WillReturnError(tc.mErr)
			}

			repo := NewStreetMarketRepository(db)

			gErr := repo.Create(context.TODO(), domain.StreetMarket{})

			if gErr.Kind != tc.wErr {
				t.Errorf("Want error kind %v, got error %v", tc.wErr, gErr.Kind)
			}
		})
	}
}

func TestStreetMarketRepository_List(t *testing.T) {
	streetMarket := domain.StreetMarket{
		ID:            "1966d99f-20e8-4e5e-8f68-eb88ca67f95f",
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
		CreatedAt:     &time.Time{},
	}

	columns := []string{
		"id",
		"long",
		"lat",
		"sectcens",
		"area",
		"iddist",
		"district",
		"idsubth",
		"subtownhall",
		"region5",
		"region8",
		"name",
		"register",
		"street",
		"number",
		"neighborhood",
		"addrextrainfo",
		"createdat",
	}

	t.Run("When use filter and return results", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("%v", err)
		}
		defer db.Close()

		want := []domain.StreetMarket{streetMarket}

		rows := sqlmock.NewRows(columns)

		for _, sm := range want {
			rows.AddRow(
				sm.ID,
				fmt.Sprintf("%v", sm.Long),
				fmt.Sprintf("%v", sm.Lat),
				sm.SectCens,
				sm.Area,
				fmt.Sprintf("%v", sm.IDdist),
				sm.District,
				fmt.Sprintf("%v", sm.IDSubTH),
				sm.SubTownHall,
				sm.Region5,
				sm.Region8,
				sm.Name,
				sm.Register,
				sm.Street,
				fmt.Sprintf("%v", sm.Number),
				sm.Neighborhood,
				sm.AddrExtraInfo,
				sm.CreatedAt,
			)
		}

		pg := domain.Pagination{
			Offset: 101,
			Limit:  100,
		}
		inp := domain.StreetMarketFilter{
			District: "district9",
			Region5:  "west",
		}

		wQB := "SELECT * FROM street_market WHERE district = $1 AND region5 = $2 ORDER BY createdat DESC OFFSET %v LIMIT %v"

		wQ := fmt.Sprintf(wQB, pg.Offset, pg.Limit)

		mock.ExpectQuery(wQ).WillReturnRows(rows)

		repo := NewStreetMarketRepository(db)

		r, dErr := repo.List(context.TODO(), pg, inp)

		if dErr != nil {
			t.Errorf("expect return nil, got %v", dErr)
		}

		if diff := cmp.Diff(want, r); diff != "" {
			t.Errorf("unexpected street market when calls create (-want +got):\n%s", diff)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("When don't use filter and return is empty", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("%v", err)
		}
		defer db.Close()

		rows := sqlmock.NewRows(columns)

		pg := domain.Pagination{
			Offset: 101,
			Limit:  100,
		}
		wQB := "SELECT * FROM street_market ORDER BY createdat DESC OFFSET %v LIMIT %v"

		wQ := fmt.Sprintf(wQB, pg.Offset, pg.Limit)
		mock.ExpectQuery(wQ).WillReturnRows(rows)

		repo := NewStreetMarketRepository(db)

		r, dErr := repo.List(context.TODO(), pg, domain.StreetMarketFilter{})

		if dErr != nil {
			t.Errorf("expect return nil, got %v", err)
		}

		want := []domain.StreetMarket{}
		if diff := cmp.Diff(want, r); diff != "" {
			t.Errorf("unexpected street market when calls create (-want +got):\n%s", diff)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestStreetMarketRepository_List_Error(t *testing.T) {
	testCases := map[string]struct {
		wErr domain.KindError
		mErr error
	}{
		"When unexpected error occurs": {
			wErr: domain.UnexpectedErrKd,
			mErr: errSome,
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("%v", err)
			}
			defer db.Close()

			mock.ExpectQuery(".+").WillReturnError(tc.mErr)

			repo := NewStreetMarketRepository(db)

			_, gErr := repo.List(context.TODO(), domain.Pagination{}, domain.StreetMarketFilter{})

			if gErr.Kind != tc.wErr {
				t.Errorf("Want error kind %v, got error %v", tc.wErr, gErr.Kind)
			}
		})
	}
}

func TestStreetMarketRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer db.Close()

	inp := domain.StreetMarket{
		ID:           "944ec25d-aac4-4c35-8301-6b35e0d7c05f",
		Name:         "RAPOSO TAVARES",
		Register:     "1129-0",
		Street:       "Rua dos Bobos",
		Number:       "500",
		Neighborhood: "JARDIM SARAH",
	}

	mock.ExpectExec(
		"UPDATE street_market SET name = $1,register = $2,street = $3,number = $4,neighborhood = $5 WHERE id = $6",
	).WithArgs(
		inp.Name,
		inp.Register,
		inp.Street,
		inp.Number,
		inp.Neighborhood,
		inp.ID,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewStreetMarketRepository(db)

	if err := repo.Update(context.TODO(), inp); err != nil {
		t.Errorf("expect return nil, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestStreetMarketRepository_Update_Error(t *testing.T) {
	testCases := map[string]struct {
		updateNothing bool
		wErr          domain.KindError
		mErr          error
	}{
		"When unexpected error occurs": {
			wErr: domain.UnexpectedErrKd,
			mErr: errSome,
		},
		"When create nothing": {
			updateNothing: true,
			wErr:          domain.NothingUpdatedErrKd,
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("%v", err)
			}
			defer db.Close()

			if tc.updateNothing {
				mock.ExpectExec(".+").WillReturnResult(sqlmock.NewResult(1, 0))
			} else {
				mock.ExpectExec(".+").WillReturnError(tc.mErr)
			}

			repo := NewStreetMarketRepository(db)

			gErr := repo.Update(context.TODO(), domain.StreetMarket{})

			if gErr.Kind != tc.wErr {
				t.Errorf("Want error kind %v, got error %v", tc.wErr, gErr.Kind)
			}
		})
	}
}

func TestStreetMarketRepository_buildArgs(t *testing.T) {
	n := time.Now()

	inp := struct {
		Foo         string
		Bar         int
		EmptyString string
		EmptyInt    int
		EmptyFloat  float64
		Createdat   *time.Time
	}{
		Foo:       "foo",
		Bar:       12,
		Createdat: &n,
	}

	wC := []string{"foo", "bar"}
	wPH := []string{"$1", "$2"}
	wV := []interface{}{"foo", 12}

	gC, gPH, gV := buildArgs(inp)

	if diff := cmp.Diff(wC, gC); diff != "" {
		t.Errorf("unexpected columns returned (-want +got):\n%s", diff)
	}

	if diff := cmp.Diff(wPH, gPH); diff != "" {
		t.Errorf("unexpected place holders returned (-want +got):\n%s", diff)
	}

	if diff := cmp.Diff(wV, gV); diff != "" {
		t.Errorf("unexpected values returned (-want +got):\n%s", diff)
	}
}

func TestStreetMarketRepository_ignore(t *testing.T) {
	if !ignore("a", []any{"a"}) {
		t.Errorf("expect ignore a")
	}

	if ignore("b", []any{"a"}) {
		t.Errorf("expect no ignore b")
	}
}
