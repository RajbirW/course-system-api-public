package repository

import (
	"RegisterApplication/internal/entity"
	"gorm.io/gorm"
)

type EnrollmentRepository struct {
	db *gorm.DB
}

func NewEnrollmentRepository(db *gorm.DB) *EnrollmentRepository {
	return &EnrollmentRepository{db}
}

func (r *EnrollmentRepository) Enroll(userID, courseID uint) error {
	enrollment := entity.Enrollment{
		UserID:   userID,
		CourseID: courseID,
	}
	return r.db.Create(&enrollment).Error
}

func (r *EnrollmentRepository) IsEnrolled(userID, courseID uint) (bool, error) {
	var count int64
	err := r.db.Model(&entity.Enrollment{}).
		Where("user_id = ? AND course_id = ?", userID, courseID).
		Count(&count).Error
	return count > 0, err
}

func (r *EnrollmentRepository) GetEnrolledCourses(userID uint) ([]entity.Course, error) {
	var courses []entity.Course
	err := r.db.Joins("JOIN enrollments ON enrollments.course_id = courses.id").
		Where("enrollments.user_id = ?", userID).
		Preload("Sections.Topics").
		Find(&courses).Error
	return courses, err
}
