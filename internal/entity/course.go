package entity

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Instructor  string     `json:"instructor"`
	Sections    []Section  `gorm:"foreignKey:CourseID" json:"sections"`
	Enrollments []Enrollment `json:"-"`
}
