package models

import "gorm.io/gorm"

type LoanRequest struct {
	gorm.Model
	UserID    uint    `gorm:"unique;not null"`
	Amount    float64 `gorm:"type:numeric(15,2);default:0"`
	Interest float64 `gorm:"type:numeric(15,2);default:0"`
	Status    string  `gorm:"size:20;unique;not null"`
}

type Loan struct {
	gorm.Model
	UserID    uint    `gorm:"unique;not null"`
	Amount    float64 `gorm:"type:numeric(15,2);default:0"`
	Interest float64 `gorm:"type:numeric(15,2);default:0"`
	Status    string  `gorm:"size:20;unique;not null"`
	LoanDate  string  `gorm:"size:20;unique;not null"`
	DueData   string  `gorm:"size:20;unique;not null"`
}