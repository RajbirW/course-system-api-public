package repository

import (
	"RegisterApplication/internal/entity"
	"gorm.io/gorm"
)

type SectionRepository struct {
	db *gorm.DB
}

func NewSectionRepository(db *gorm.DB) *SectionRepository {
	return &SectionRepository{db}
}

func (r *SectionRepository) Create(section *entity.Section) error {
	return r.db.Create(section).Error
}

func (r *SectionRepository) FindByCourseID(courseID uint) ([]entity.Section, error) {
	var sections []entity.Section
	err := r.db.Preload("Topics").Where("course_id = ?", courseID).Find(&sections).Error
	return sections, err
}
