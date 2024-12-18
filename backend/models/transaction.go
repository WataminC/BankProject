package models

import (
	"time"
)

// Transaction 交易表
type Transaction struct {
	ID                  uint      `gorm:"primaryKey"`
	AccountID           uint      `gorm:"not null"`
	Type                string    `gorm:"size:20;not null"`
	Amount              float64   `gorm:"type:numeric(15,2);not null"` // 使用 numeric 类型
	TargetAccountNumber *string   `gorm:"size:20"`                     // 可为空的目标账户
	CreatedAt           time.Time `gorm:"autoCreateTime"`              // 自动设置创建时间
	UpdatedAt           time.Time `gorm:"autoUpdateTime"`              // 自动更新更新时间
}
