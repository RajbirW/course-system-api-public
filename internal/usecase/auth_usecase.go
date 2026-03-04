package usecase

import (
	"github.com/google/uuid"
	"RegisterApplication/internal/entity"
	"RegisterApplication/internal/repository"
)

type AuthUsecase struct {
	repo *repository.UserRepository
}

func NewAuthUsecase(repo *repository.UserRepository) *AuthUsecase {
	return &AuthUsecase{repo: repo}
}

func (u *AuthUsecase) Register(username, password, role string) (*entity.User, error) {
	user := &entity.User{
		Username: username,
		Password: password,
		Role:     role,
	}
	err := u.repo.Create(user)
	return user, err
}

func (u *AuthUsecase) Login(username, password string) (string, error) {
	user, err := u.repo.FindByCredentials(username, password)
	if err != nil {
		return "", err
	}
	user.Token = uuid.New().String()
	err = u.repo.Save(user)
	return user.Token, err
}
