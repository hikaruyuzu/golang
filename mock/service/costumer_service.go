package service

import (
	"errors"
	"testify-mock/entity"
	"testify-mock/repository"
)

type CostumerService struct {
	Repository repository.CostumerRepository
}

func (service CostumerService) Get(id string) (*entity.Costumer, error) {
	// memanggil kontrak dari package repository
	result := service.Repository.GetById(id)
	if result != nil {
		return result, nil
	} else {
		return nil, errors.New("Data costumer is not available")
	}

}
