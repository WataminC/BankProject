package models

import (
	"time"
)

type LoanRequest struct {
	ID       uint    `gorm:"primaryKey"`
	UserID   uint    `gorm:"unique;not null"`
	Amount   float64 `gorm:"type:numeric(15,2);default:0"`
	Interest float64 `gorm:"type:numeric(15,2);default:0"`
	Status   string  `gorm:"size:20;unique;not null"`

	CreatedAt time.Time `gorm:"autoCreateTime"` // 自动设置创建时间
	UpdatedAt time.Time `gorm:"autoUpdateTime"` // 自动更新更新时间
}

type Loan struct {
	ID          uint    `gorm:"primaryKey"`
	UserID      uint    `gorm:"unique;not null"`
	Amount      float64 `gorm:"type:numeric(15,2);default:0"`
	Interest    float64 `gorm:"type:numeric(15,2);default:0"`
	TotalAmount float64 `gorm:"type:numeric(15,2);default:0"`
	Status      string  `gorm:"size:20;not null"`

	CreatedAt time.Time `gorm:"autoCreateTime"` // 自动设置创建时间
	UpdatedAt time.Time `gorm:"autoUpdateTime"` // 自动更新更新时间
}
