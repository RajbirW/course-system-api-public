package entity

import (
	"time"
	"gorm.io/gorm"
)
type Document struct {
	ID uint `gorm:"primary key"`
	Filename string `gorm:"not null"`
	Path string `gorm:"not null"`
	UserID     uint      // optional, if linked to a specific user
	ParentID   *uint     // optional, for parent-child document relationships
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}