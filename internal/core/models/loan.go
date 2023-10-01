package models

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type Loan struct {
	Base
	CustomerID    uuid.UUID `json:"customerId,omitempty"  gorm:"<-:create;column:customer_id;type:uuid;NOT NULL"`
	RepaymentType string    `json:"repaymentType,omitempty"  gorm:"<-:create;column:repayment_type;NOT NULL"`
	Amount        float64   `json:"amount,omitempty" gorm:"<-:create;column:amount;NOT NULL"`
	Term          int32     `json:"term,omitempty" gorm:"<-:create;column:term;NOT NULL"`
	State         string    `json:"state,omitempty" gorm:"column:state;NOT NULL"`
}

func (Loan) TableName() string {
	return "loans"
}

func (Loan) InsertClause() []clause.Expression {
	return []clause.Expression{
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoNothing: true,
		},
	}
}

func (l Loan) ValidateCreate() error {
	if l.Amount <= 0 {
		return fmt.Errorf("missing loan amount")
	}
	if l.Term <= 0 {
		return fmt.Errorf("missing loan term")
	}
	return nil
}
