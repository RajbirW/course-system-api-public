package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Password string `json:"-"`
	Token    string `json:"-"`
	Role     string `json:"role"` // "admin" or "student"

	Enrollments []Enrollment `json:"-"`
}
