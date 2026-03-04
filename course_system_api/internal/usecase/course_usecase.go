package usecase

import (
	"RegisterApplication/internal/entity"
	"RegisterApplication/internal/repository"
)

type CourseUsecase struct {
	repo *repository.CourseRepository
}

func NewCourseUsecase(r *repository.CourseRepository) *CourseUsecase {
	return &CourseUsecase{repo: r}
}

func (u *CourseUsecase) Create(course *entity.Course) error {
	return u.repo.Create(course)
}

func (u *CourseUsecase) GetAll() ([]entity.Course, error) {
	return u.repo.FindAll()
}

func (u *CourseUsecase) GetByID(id uint) (*entity.Course, error) {
	return u.repo.FindByID(id)
}

func (u *CourseUsecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
