package usecase

import (
	"RegisterApplication/internal/entity"
	"RegisterApplication/internal/repository"
)

type DocumentUsecase struct {
	repo *repository.DocumentRepository
}

func NewDocumentUsecase(r *repository.DocumentRepository) *DocumentUsecase {
	return &DocumentUsecase{repo: r}
}

func (u *DocumentUsecase) Upload(doc *entity.Document) error {
	return u.repo.Create(doc)
}

func (u *DocumentUsecase) GetAll() ([]entity.Document, error) {
	return u.repo.GetAll()
}

func (u *DocumentUsecase) GetByID(id uint) (*entity.Document, error) {
	return u.repo.GetByID(id)
}

func (u *DocumentUsecase) Delete(id uint) error {
	return u.repo.SoftDelete(id)
}
