package service

import (
	"camera-rent/entity"
	"camera-rent/input"
	"camera-rent/repository"
)

type ServiceTopUp interface {
	CreateTopUp(input input.InputTopUp, userID int) (*entity.TopUp, error)
	GetTopUp(ID int) (*entity.TopUp, error)
}

type serviceTopUp struct {
	repositoryTopUp repository.RepositoryTopUp
	repositoryUser  repository.RepositoryUser
}

func NewServiceTopUp(repositoryTopUp repository.RepositoryTopUp, repositoryUser repository.RepositoryUser) *serviceTopUp {
	return &serviceTopUp{repositoryTopUp, repositoryUser}
}

func (s *serviceTopUp) CreateTopUp(input input.InputTopUp, userID int) (*entity.TopUp, error) {

	findUser, err := s.repositoryUser.FindById(userID)
	if err != nil {
		return nil, err
	}

	topUp := &entity.TopUp{}

	topUp.Amount = input.Amount
	topUp.UserID = findUser.ID

	newtopUp, err := s.repositoryTopUp.Save(topUp)
	if err != nil {
		return newtopUp, err
	}
	return newtopUp, nil
}

func (s *serviceTopUp) GetTopUp(ID int) (*entity.TopUp, error) {

	product, err := s.repositoryTopUp.FindById(ID)

	if err != nil {
		return nil, err
	}

	if product.ID == 0 {
		return nil, err
	}

	return product, nil
}
