package usecase

import (
	"RegisterApplication/internal/entity"
	"RegisterApplication/internal/repository"
)

type TopicUsecase struct {
	repo *repository.TopicRepository
}

func NewTopicUsecase(r *repository.TopicRepository) *TopicUsecase {
	return &TopicUsecase{repo: r}
}

func (u *TopicUsecase) Create(topic *entity.Topic) error {
	return u.repo.Create(topic)
}

func (u *TopicUsecase) GetBySectionID(sectionID uint) ([]entity.Topic, error) {
	return u.repo.FindBySectionID(sectionID)
}
