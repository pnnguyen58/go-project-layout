package repositories

import (
	"context"
	"github.com/google/uuid"

	"github.com/pnnguyen58/go-project-layout/internal/core/models"
	"github.com/pnnguyen58/go-project-layout/internal/utils/defined"
)

type Loan interface {
	GetByID(context.Context, uuid.UUID) (*models.Loan, error)
	Create(context.Context, *models.Loan) error
	UpdateState(context.Context, uuid.UUID, defined.State) error
	Delete(context.Context, uuid.UUID) error
}
