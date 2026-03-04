package entity

import "gorm.io/gorm"

type Section struct {
	gorm.Model
	Title    string   `json:"title"`
	CourseID uint     `json:"course_id"`
	Topics   []Topic  `gorm:"foreignKey:SectionID" json:"topics"`
}
