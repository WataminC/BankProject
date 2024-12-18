package models

import (
	"time"
)

// Account 账户表
type Account struct {
	ID            uint      `gorm:"primaryKey"`
	UserID        uint      `gorm:"not null"`
	AccountNumber string    `gorm:"size:20;unique;not null"`
	Balance       float64   `gorm:"type:numeric(15,2);default:0"` // 使用 numeric 类型
	CreatedAt     time.Time `gorm:"autoCreateTime"`               // 自动设置创建时间
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`               // 自动更新更新时间
}