package usecase

import (
	"abc/internal/model"
	"abc/internal/repository"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) CreateUser(user *model.User) error {
	return u.repo.CreateUser(user)
}

func (u *UserUsecase) GetUsers() ([]*model.User, error) {
	return u.repo.GetAllUsers()
}
