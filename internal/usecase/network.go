package usecase

import (
	"abc/internal/model"
	"abc/internal/repository"
)

type NetworkUsecase struct {
	repo repository.NetworkRepository
	// network adaptor.NetworkAdapter
}

func NewNetworkUsecase(repo repository.NetworkRepository) *NetworkUsecase {
	return &NetworkUsecase{repo: repo}
}

func (u *NetworkUsecase) CreateNetwork(user *model.User) error {
	return u.repo.CreateNetwork(user)
}

func (u *NetworkUsecase) GetNetworks() ([]*model.User, error) {
	return u.repo.GetNetworks()
}
