package repository

import "abc/internal/model"

type NetworkRepository interface {
	CreateNetwork(user *model.User) error
	GetNetworks() ([]*model.User, error)
}
