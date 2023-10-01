package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type RepaymentLog struct {
	Base
	LoanID uuid.UUID `json:"loanId,omitempty"  gorm:"<-:create;column:loan_id;type:uuid;NOT NULL"`
	Amount float64   `json:"amount,omitempty" gorm:"column:amount;NOT NULL"`
}

func (RepaymentLog) TableName() string {
	return "repayment_logs"
}

func (RepaymentLog) InsertClause() []clause.Expression {
	return []clause.Expression{
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoNothing: true,
		},
	}
}
