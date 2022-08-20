package streetmarket

import (
	"context"
	"fmt"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
)

type repositoryEraser interface {
	DeleteByID(ctx context.Context, ID string) error
}

type StreetMarketEraser struct {
	repo repositoryEraser
}

func NewEraser(repo repositoryEraser) *StreetMarketEraser {
	return &StreetMarketEraser{repo}
}

func (s *StreetMarketEraser) Delete(ctx context.Context, ID domain.SMID) error {
	if err := ID.Validate(); err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := s.repo.DeleteByID(ctx, string(ID)); err != nil {
		switch err {
		case domain.ErrNothingDeleted:
			return domain.ErrSMNotFound
		default:
			return domain.ErrUnexpected
		}
	}

	return nil
}
