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

type UserLogin struct {
	Email    string `json:"email" binding:"required,email,lte=150"`
	Password string `json:"password" binding:"required,gte=10,lte=150"`
}

type UserRequest struct {
	Name     string `json:"name" binding:"required,lte=150"`
	Email    string `json:"email" binding:"required,email,lte=150"`
	Password string `json:"password" binding:"required,gte=10,lte=150"`
	Age      uint   `json:"age" binding:"required,gte=1,lte=150"`
}
