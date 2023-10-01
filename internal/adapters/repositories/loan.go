package repositories

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/pnnguyen58/go-project-layout/internal/core/models"
	"github.com/pnnguyen58/go-project-layout/internal/core/ports/repositories"
	"github.com/pnnguyen58/go-project-layout/internal/utils/defined"
)

type loan struct {
	db *gorm.DB
}

func NewLoan(db *gorm.DB) repositories.Loan {
	return &loan{
		db: db,
	}
}

func (l *loan) GetByID(ctx context.Context, id uuid.UUID) (*models.Loan, error) {
	var result models.Loan
	err := l.db.WithContext(ctx).
		Model(models.Loan{}).
		Where("id = ?", id).
		First(&result).Error
	return &result, err
}

func (l *loan) Create(ctx context.Context, loan *models.Loan) error {
	err := l.db.Clauses(loan.InsertClause()...).
		Model(models.Loan{}).
		WithContext(ctx).
		Create(loan).Error
	return err
}

func (l *loan) UpdateState(ctx context.Context, id uuid.UUID, state defined.State) error {
	err := l.db.Model(models.Loan{}).
		WithContext(ctx).
		Where("id = ?", id).
		Update("state", state).Error
	return err
}

func (l *loan) Delete(ctx context.Context, id uuid.UUID) error {
	err := l.db.WithContext(ctx).
		Delete(&models.RepaymentPrepay{}, id).Error
	return err
}
