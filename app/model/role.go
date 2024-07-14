package model

import "time"

// Role model
type Role struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"size:50"`
	Slug      string `gorm:"size:50;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
