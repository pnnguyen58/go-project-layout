package repositories

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/pnnguyen58/go-project-layout/internal/core/models"
	"github.com/pnnguyen58/go-project-layout/internal/core/ports/repositories"
	"github.com/pnnguyen58/go-project-layout/internal/utils/defined"
)

func NewRepayment(db *gorm.DB) repositories.Repayment {
	return &repayment{
		db: db,
	}
}

type repayment struct {
	db *gorm.DB
}

func (r *repayment) CreateLog(ctx context.Context, log *models.RepaymentLog) error {
	err := r.db.Clauses(log.InsertClause()...).
		Model(models.RepaymentLog{}).
		WithContext(ctx).
		Create(log).Error
	return err
}

func (r *repayment) CreatePrepay(ctx context.Context, prepay *models.RepaymentPrepay) error {
	err := r.db.Clauses(prepay.InsertClause()...).
		Model(models.RepaymentPrepay{}).
		WithContext(ctx).
		Create(prepay).Error
	return err
}

func (r *repayment) DeletePrepay(ctx context.Context, prepayID uuid.UUID) error {
	err := r.db.WithContext(ctx).
		Delete(&models.RepaymentPrepay{}, prepayID).Error
	return err
}

func (r *repayment) Create(ctx context.Context, repayment *models.Repayment) error {
	err := r.db.Clauses(repayment.InsertClause()...).
		Model(models.Repayment{}).
		WithContext(ctx).
		Create(&repayment).Error
	return err
}

func (r *repayment) CreateBatch(ctx context.Context, repayments []*models.Repayment, batchSize int) error {
	if len(repayments) == 0 {
		return fmt.Errorf("empty repayments")
	}
	if batchSize == 0 {
		return fmt.Errorf("invalid batchSize")
	}
	err := r.db.Clauses(repayments[0].InsertClause()...).
		Model(models.Repayment{}).
		WithContext(ctx).
		CreateInBatches(repayments, batchSize).Error
	return err
}

func (r *repayment) UpdateState(ctx context.Context, id uuid.UUID, state defined.State) error {
	err := r.db.WithContext(ctx).
		Model(models.Repayment{}).
		Where("id = ?", id).
		Update("state", state).Error
	return err
}

func (r *repayment) UpdateStateByLoanID(ctx context.Context, loanID uuid.UUID, state defined.State) error {
	err := r.db.WithContext(ctx).
		Model(models.Repayment{}).
		Where("loan_id = ?", loanID).
		Update("state", state).Error
	return err
}

func (r *repayment) GetByLoanID(ctx context.Context, loanID uuid.UUID) ([]*models.Repayment, error) {
	var result []*models.Repayment
	err := r.db.WithContext(ctx).
		Model(models.Repayment{}).
		Where("loan_id = ?", loanID).
		Find(&result).Error
	return result, err
}

func (r *repayment) GetUnpaidByLoanID(ctx context.Context, loanID uuid.UUID) ([]*models.Repayment, error) {
	var result []*models.Repayment
	err := r.db.WithContext(ctx).
		Model(models.Repayment{}).
		Where("loan_id = ?", loanID).
		Where("state = ?", defined.APPROVED).
		Order("due_date ASC").
		Find(&result).Error
	return result, err
}
