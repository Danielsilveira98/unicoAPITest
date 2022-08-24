package streetmarket

import (
	"context"
	"testing"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
)

type stubRepositoryEraser struct {
	deleteInp  string
	deleteByID func(ctx context.Context, ID string) *domain.Error
}

func (s *stubRepositoryEraser) DeleteByID(ctx context.Context, ID string) *domain.Error {
	s.deleteInp = ID
	return s.deleteByID(ctx, ID)
}

func TestStreetMarketEraser_Delete(t *testing.T) {
	var wID domain.SMID = "662ea609-2ef5-4026-a402-218dd316cc03"

	repoMock := &stubRepositoryEraser{
		deleteByID: func(ctx context.Context, ID string) *domain.Error {
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
		rErr *domain.Error
		ID   domain.SMID
		wErr domain.KindError
	}{
		"When id is invalid": {
			wErr: domain.InpValidationErrKd,
			ID:   "invalid",
		},
		"When street market not exists": {
			rErr: &domain.Error{Kind: domain.NothingDeletedErrKd},
			wErr: domain.SMNotFoundErrKd,
			ID:   "29336645-6243-4279-b7ff-47f1a64aa781",
		},
		"When a unexpected error occurs in repository": {
			rErr: unexpectedErr,
			wErr: domain.UnexpectedErrKd,
			ID:   "6c34a17f-6330-4625-9184-25eb0a5c6533",
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			repoMock := &stubRepositoryEraser{
				deleteByID: func(ctx context.Context, ID string) *domain.Error {
					return tc.rErr
				},
			}

			srv := NewEraser(repoMock)

			gErr := srv.Delete(context.TODO(), tc.ID)

			if gErr.Kind != tc.wErr {
				t.Errorf("Want error kind %v, got error %v", tc.wErr, gErr.Kind)
			}
		})
	}
}
