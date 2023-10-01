package models

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
	"time"
)

type Repayment struct {
	Base
	LoanID  uuid.UUID `json:"loanId,omitempty"  gorm:"<-:create;column:loan_id;type:uuid;NOT NULL"`
	Amount  float64   `json:"amount,omitempty" gorm:"<-:create;column:amount;NOT NULL"`
	DueDate time.Time `json:"dueDate,omitempty" gorm:"<-:create;column:due_date;NOT NULL"`
	State   string    `json:"state,omitempty" gorm:"column:state;NOT NULL"`
}

func (Repayment) TableName() string {
	return "repayments"
}

func (Repayment) InsertClause() []clause.Expression {
	return []clause.Expression{
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoNothing: true,
		},
	}
}

func (r Repayment) ValidateCreate() error {
	if r.Amount <= 0 {
		return fmt.Errorf("missing repayment amount")
	}
	if r.DueDate.IsZero() {
		return fmt.Errorf("missing repayment due date")
	}
	if r.LoanID == uuid.Nil {
		return fmt.Errorf("missing loan id")
	}
	return nil
}
