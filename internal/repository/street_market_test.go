package repository

import (
	"context"
	"errors"
	"testing"

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
		wErr   error
		mErr   error
		notUpd bool
		id     string
	}{
		"When unexpected error occurs": {
			wErr: domain.ErrUnexpected,
			mErr: errSome,
			id:   "822d08fb-6dbd-457b-a467-36ee3a136b13",
		},
		"When update nothing": {
			notUpd: true,
			wErr:   domain.ErrNothingDeleted,
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

			if !errors.Is(gErr, tc.wErr) {
				t.Errorf("Want error %v, got error %v", tc.wErr, gErr)
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
		Number:       500,
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
		wErr          error
		mErr          error
	}{
		"When unexpected error occurs": {
			wErr: domain.ErrUnexpected,
			mErr: errSome,
		},
		"When create nothing": {
			createNothing: true,
			wErr:          domain.ErrNothingCreated,
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

			if !errors.Is(gErr, tc.wErr) {
				t.Errorf("Want error %v, got error %v", tc.wErr, gErr)
			}
		})
	}
}

// func TestStreetMarketRepository_Select(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("%v", err)
// 	}
// 	defer db.Close()

// 	mock.ExpectExec("")

// 	repo := NewStreetMarketRepository(db)

// 		if err := mock.ExpectationsWereMet(); err != nil {
// 			t.Errorf("there were unfulfilled expectations: %s", err)
// 		}
// }

// func TestStreetMarketRepository_Select_Error(t *testing.T) {
// testCases := map[string]struct {
// 	not bool
// 	wErr     error
// 	mErr     error
// }{
// 	"": {},
// }

// for title, tc := range testCases {
// 	t.Run(title, func(t *testing.T) {
// 		db, mock, err := sqlmock.New()
// 		if err != nil {
// 			t.Fatalf("%v", err)
// 		}
// 		defer db.Close()

// 		if tc.not {
// 			mock.ExpectExec().WithArgs().WillReturnResult(sqlmock.NewResult(1, 0))
// 		} else {
// 			mock.ExpectExec().WithArgs().WillReturnError(tc.mErr)
// 		}

// 		repo := NewStreetMarketRepository(db)

// 		_, gErr := repo.(context.TODO(), )

// 		if !errors.Is(gErr, tc.wErr) {
// 			t.Errorf("Want error %v, got error %v", tc.wErr, gErr)
// 		}
// 	})
// }

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
		Number:       500,
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
		wErr          error
		mErr          error
	}{
		"When unexpected error occurs": {
			wErr: domain.ErrUnexpected,
			mErr: errSome,
		},
		"When create nothing": {
			updateNothing: true,
			wErr:          domain.ErrNothingUpdated,
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
				mock.ExpectExec(".+").WithArgs().WillReturnResult(sqlmock.NewResult(1, 0))
			} else {
				mock.ExpectExec(".+").WithArgs().WillReturnError(tc.mErr)
			}

			repo := NewStreetMarketRepository(db)

			gErr := repo.Update(context.TODO(), domain.StreetMarket{})

			if !errors.Is(gErr, tc.wErr) {
				t.Errorf("Want error %v, got error %v", tc.wErr, gErr)
			}
		})
	}
}

func TestStreetMarketRepository_buildArgs(t *testing.T) {
	inp := struct {
		Foo         string
		Bar         int
		EmptyString string
		EmptyInt    int
	}{
		Foo: "foo",
		Bar: 12,
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
