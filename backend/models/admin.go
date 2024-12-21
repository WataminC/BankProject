package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	UserId    uint      `gorm:"unique;not null"`
}
