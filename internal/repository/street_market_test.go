package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
)

var errSome = errors.New("some error")

func TestStreetMarketRepoistory_Delete(t *testing.T) {
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

// func TestStreetMarketRepoistory_Create(t *testing.T) {
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

// func TestStreetMarketRepository_Create_Error(t *testing.T) {
// testCases := map[string]struct {
// 	notFound bool
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

// 		if tc.notFound {
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

// func TestStreetMarketRepoistory_Select(t *testing.T) {
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
// 	notFound bool
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

// 		if tc.notFound {
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

// func TestStreetMarketRepoistory_Edit(t *testing.T) {
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

// func TestStreetMarketRepository_Edit_Error(t *testing.T) {
// testCases := map[string]struct {
// 	notFound bool
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

// 		if tc.notFound {
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
