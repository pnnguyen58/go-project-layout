package activities

import (
	"math"
	"time"

	"github.com/pnnguyen58/go-project-layout/internal/core/models"
	"github.com/pnnguyen58/go-project-layout/internal/utils/defined"
)

func makeRepayment(loan models.Loan) []*models.Repayment {
	var (
		repayments []*models.Repayment
		i          int32
	)
	amount := loan.Amount
	term := loan.Term
	termAmount := roundAmount(amount / float64(term))
	for i = 1; i <= term; i++ {
		r := &models.Repayment{
			LoanID:  loan.ID,
			DueDate: setRepaymentDueDate(loan.CreatedAt, i),
			State:   string(defined.PENDING),
		}
		if i == term {
			r.Amount = amount - termAmount*float64(term-1)
		} else {
			r.Amount = termAmount
		}
		repayments = append(repayments, r)
	}
	return repayments
}

func setRepaymentDueDate(createdAt time.Time, next int32) time.Time {
	return createdAt.AddDate(0, 0, int(next*7))
	// TODO: implement other types of repayment: monthly, daily, ...
}

func roundAmount(amount float64) float64 {
	ratio := math.Pow(10, float64(defined.PRECISION))
	return math.Round(amount*ratio) / ratio
}
