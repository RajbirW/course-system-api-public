package repository

import (
	"RegisterApplication/internal/entity"
	"gorm.io/gorm"
)

type DocumentRepository struct {
	db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) *DocumentRepository {
	return &DocumentRepository{db}
}

func (r *DocumentRepository) Create (doc *entity.Document) error {
	return r.db.Create(doc).Error
}

func (r *DocumentRepository) GetAll() ([]entity.Document, error) {
	var docs []entity.Document
	err := r.db.Where("deleted_at is NULL").Find(&docs).Error
	return docs, err
}

func (r *DocumentRepository) GetByID(id uint) (*entity.Document, error) {
	var doc entity.Document
	err := r.db.First(&doc, "id = ? AND deleted_at IS NULL", id).Error
	return &doc, err
}

func (r *DocumentRepository) SoftDelete(id uint) error {
	return r.db.Delete(&entity.Document{}, id).Error
}
