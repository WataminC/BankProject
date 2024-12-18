package models

import (
	"time"
)

// User 用户表
type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:100;not null"`
	Email     string    `gorm:"size:150;unique;not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"` // 自动设置创建时间
	UpdatedAt time.Time `gorm:"autoUpdateTime"` // 自动更新更新时间
}