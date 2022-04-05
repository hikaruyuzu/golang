package repository

import "testify-mock/entity"

// kontrak query
type CostumerRepository interface {
	GetById(id string) *entity.Costumer
}
