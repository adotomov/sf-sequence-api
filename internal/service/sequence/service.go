package sequence

import (
	"context"

	"github.com/adotomov/sf-sequence-api/internal/models"
	db "github.com/adotomov/sf-sequence-api/internal/repository/sql"
	"github.com/adotomov/sf-sequence-api/internal/types"
)

type SequenceService struct {
	env string
	d   *db.DB
}

func NewService(env string, db *db.DB) *SequenceService {
	return &SequenceService{
		env: env,
		d:   db,
	}
}

func (sq *SequenceService) GetSequences(ctx context.Context) ([]*models.Sequence, error) {
	res, err := sq.d.GetSequences(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (sq *SequenceService) CreateSequence(ctx context.Context, rb types.CreateSequenceRequest) (*models.Sequence, error) {
	res, err := sq.d.CreateSequence(ctx, rb)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (sq *SequenceService) GetSequenceByID(ctx context.Context, id string) (*models.Sequence, error) {
	res, err := sq.d.GetSequenceByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (sq *SequenceService) UpdateSequence(ctx context.Context, rb types.UpdateSequenceRequest) (*models.Sequence, error) {
	res, err := sq.d.UpdateSequence(ctx, rb)
	if err != nil {
		return nil, err
	}

	return res, nil
}
