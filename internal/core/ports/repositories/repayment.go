package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/pnnguyen58/go-project-layout/internal/core/models"
	"github.com/pnnguyen58/go-project-layout/internal/utils/defined"
)

type Repayment interface {
	CreateLog(context.Context, *models.RepaymentLog) error

	Create(context.Context, *models.Repayment) error
	CreateBatch(context.Context, []*models.Repayment, int) error
	UpdateState(context.Context, uuid.UUID, defined.State) error
	UpdateStateByLoanID(context.Context, uuid.UUID, defined.State) error

	CreatePrepay(context.Context, *models.RepaymentPrepay) error
	DeletePrepay(context.Context, uuid.UUID) error

	GetByLoanID(context.Context, uuid.UUID) ([]*models.Repayment, error)
	GetUnpaidByLoanID(context.Context, uuid.UUID) ([]*models.Repayment, error)
}
