package usecase

import (
	"RegisterApplication/internal/entity"
	"RegisterApplication/internal/repository"
)

type SectionUsecase struct {
	repo *repository.SectionRepository
}

func NewSectionUsecase(r *repository.SectionRepository) *SectionUsecase {
	return &SectionUsecase{repo: r}
}

func (u *SectionUsecase) Create(section *entity.Section) error {
	return u.repo.Create(section)
}

func (u *SectionUsecase) GetByCourseID(courseID uint) ([]entity.Section, error) {
	return u.repo.FindByCourseID(courseID)
}
