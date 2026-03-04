package repository

import (
	"RegisterApplication/internal/entity"
	"gorm.io/gorm"
)

type CourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{db}
}

func (r *CourseRepository) Create(course *entity.Course) error {
	return r.db.Create(course).Error
}

func (r *CourseRepository) FindAll() ([]entity.Course, error) {
	var courses []entity.Course
	err := r.db.Preload("Sections.Topics").Find(&courses).Error
	return courses, err
}

func (r *CourseRepository) FindByID(id uint) (*entity.Course, error) {
	var course entity.Course
	err := r.db.Preload("Sections.Topics").First(&course, id).Error
	return &course, err
}

func (r *CourseRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Course{}, id).Error
}
