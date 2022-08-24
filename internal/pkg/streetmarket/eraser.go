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
		return fmt.Errorf("[ID.Validate] %s %w", domain.ErrInpValidation.Error(), err)
	}

	if err := s.repo.DeleteByID(ctx, string(ID)); err != nil {
		var msg string
		switch err {
		case domain.ErrNothingDeleted:
			msg = domain.ErrSMNotFound.Error()
		default:
			msg = domain.ErrUnexpected.Error()
		}

		return fmt.Errorf("[repo.DeleteByID] %s %w", msg, err)
	}

	return nil
}
