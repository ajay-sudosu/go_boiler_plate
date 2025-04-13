package repository

import "abc/internal/model"

type UserRepository interface {
	CreateUser(user *model.User) error
	GetAllUsers() ([]*model.User, error)
}
