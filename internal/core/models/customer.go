package models

import (
	"gorm.io/gorm/clause"
)

type Customer struct {
	Base
	Name string `json:"name,omitempty"  gorm:"unique;column:name;NOT NULL"`
}

func (Customer) TableName() string {
	return "customers"
}

func (Customer) InsertClause() []clause.Expression {
	return []clause.Expression{
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}, {Name: "name"}},
			DoNothing: true,
		},
	}
}
