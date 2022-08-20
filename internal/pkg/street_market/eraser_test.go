package streetmarket

import (
	"context"
	"errors"
	"testing"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
)

type stubRepositoryEraser struct {
	deleteInp  string
	deleteByID func(ctx context.Context, ID string) error
}

func (s *stubRepositoryEraser) DeleteByID(ctx context.Context, ID string) error {
	s.deleteInp = ID
	return s.deleteByID(ctx, ID)
}

func TestStreetMarketEraser_Delete(t *testing.T) {
	var wID domain.SMID = "662ea609-2ef5-4026-a402-218dd316cc03"

	repoMock := &stubRepositoryEraser{
		deleteByID: func(ctx context.Context, ID string) error {
			return nil
		},
	}

	srv := NewEraser(repoMock)

	err := srv.Delete(context.TODO(), wID)

	if err != nil {
		t.Errorf("want return nil, got %v", err)
	}

	if string(wID) != repoMock.deleteInp {
		t.Errorf("unexpected id when call deletebyid, want %s, got %s", wID, repoMock.deleteInp)
	}
}

func TestStreetMarketEraser_Delete_Error(t *testing.T) {
	testCases := map[string]struct {
		rErr error
		ID   domain.SMID
		wErr error
	}{
		"When id is invalid": {
			wErr: domain.ErrInpValidation,
			ID:   "invalid",
		},
		"When street market not exists": {
			rErr: domain.ErrNothingDeleted,
			wErr: domain.ErrSMNotFound,
			ID:   "29336645-6243-4279-b7ff-47f1a64aa781",
		},
		"When a unexpected error occurs in repository": {
			rErr: errSome,
			wErr: domain.ErrUnexpected,
			ID:   "6c34a17f-6330-4625-9184-25eb0a5c6533",
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			repoMock := &stubRepositoryEraser{
				deleteByID: func(ctx context.Context, ID string) error {
					return tc.rErr
				},
			}

			srv := NewEraser(repoMock)

			gErr := srv.Delete(context.TODO(), tc.ID)

			if !errors.Is(gErr, tc.wErr) {
				t.Errorf("Want error %v, got error %v", tc.wErr, gErr)
			}
		})
	}
}
