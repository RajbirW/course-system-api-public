package usecase

import (
	"RegisterApplication/internal/entity"
	"RegisterApplication/internal/repository"
)

type EnrollmentUsecase struct {
	repo *repository.EnrollmentRepository
}

func NewEnrollmentUsecase(r *repository.EnrollmentRepository) *EnrollmentUsecase {
	return &EnrollmentUsecase{repo: r}
}

func (u *EnrollmentUsecase) EnrollUser(userID, courseID uint) error {
	return u.repo.Enroll(userID, courseID)
}

func (u *EnrollmentUsecase) IsUserEnrolled(userID, courseID uint) (bool, error) {
	return u.repo.IsEnrolled(userID, courseID)
}

func (u *EnrollmentUsecase) GetUserCourses(userID uint) ([]entity.Course, error) {
	return u.repo.GetEnrolledCourses(userID)
}
