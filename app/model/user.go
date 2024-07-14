package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"size:50"`
	Email     string `gorm:"unique"`
	Password  string
	Age       uint
	Status    bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
