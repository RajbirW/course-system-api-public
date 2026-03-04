package repository

import (
	"RegisterApplication/internal/entity"
	"gorm.io/gorm"
	"errors"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByToken(token string) (entity.User, error) {
	if token == "" {
		return entity.User{}, errors.New("token empty")
	}
	var user entity.User
	result := r.db.Where("token = ?", token).First(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil
}


func (r *UserRepository) FindByCredentials(username, password string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("username = ? AND password = ?", username, password).First(&user).Error
	return &user, err
}

func (r *UserRepository) Save(user *entity.User) error {
	return r.db.Save(user).Error
}
