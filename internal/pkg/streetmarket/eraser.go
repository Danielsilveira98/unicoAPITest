package streetmarket

import (
	"context"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
)

type repositoryEraser interface {
	DeleteByID(ctx context.Context, ID string) *domain.Error
}

type StreetMarketEraser struct {
	repo repositoryEraser
}

func NewEraser(repo repositoryEraser) *StreetMarketEraser {
	return &StreetMarketEraser{repo}
}

func (s *StreetMarketEraser) Delete(ctx context.Context, ID domain.SMID) *domain.Error {
	if err := ID.Validate(); err != nil {
		return &domain.Error{
			Kind:     domain.InpValidationErrKd,
			Msg:      "Invalid ID",
			Previous: err,
		}
	}

	if err := s.repo.DeleteByID(ctx, string(ID)); err != nil {
		switch err.Kind {
		case domain.NothingDeletedErrKd:
			return &domain.Error{Kind: domain.SMNotFoundErrKd, Msg: "Entity not exists", Previous: err}
		default:
			return &domain.Error{Kind: domain.UnexpectedErrKd, Msg: "Unexpected error when delete", Previous: err}
		}
	}

	return nil
}
