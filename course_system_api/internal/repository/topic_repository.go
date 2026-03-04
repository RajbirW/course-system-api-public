package repository

import (
	"RegisterApplication/internal/entity"
	"gorm.io/gorm"
)

type TopicRepository struct {
	db *gorm.DB
}

func NewTopicRepository(db *gorm.DB) *TopicRepository {
	return &TopicRepository{db}
}

func (r *TopicRepository) Create(topic *entity.Topic) error {
	return r.db.Create(topic).Error
}

func (r *TopicRepository) FindBySectionID(sectionID uint) ([]entity.Topic, error) {
	var topics []entity.Topic
	err := r.db.Where("section_id = ?", sectionID).Find(&topics).Error
	return topics, err
}
