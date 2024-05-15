package step

import (
	"context"

	"github.com/adotomov/sf-sequence-api/internal/models"
	db "github.com/adotomov/sf-sequence-api/internal/repository/sql"
	"github.com/adotomov/sf-sequence-api/internal/types"
)

type StepService struct {
	env string
	d   *db.DB
}

func NewService(env string, db *db.DB) *StepService {
	return &StepService{
		env: env,
		d:   db,
	}
}

func (st *StepService) CreateStep(ctx context.Context, rb types.CreateStepRequest) (*models.Step, error) {
	res, err := st.d.CreateStep(ctx, rb)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (st *StepService) UpdateStep(ctx context.Context, rb types.UpdateStepRequest) (*models.Step, error) {
	res, err := st.d.UpdateStep(ctx, rb)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (st *StepService) DeleteStep(ctx context.Context, sid string) error {
	if err := st.d.DeleteStep(ctx, sid); err != nil {
		return err
	}
	return nil
}
