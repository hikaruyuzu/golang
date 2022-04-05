package service

import (
	"fmt"
	"testify-mock/entity"
	"testify-mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var costumerMockRepository = &repository.CostumerRepositoryMock{mock.Mock{}}
var costumerService = CostumerService{Repository: costumerMockRepository}

func TestGetDataCostumerNotFound(t *testing.T) {

	// parameter mock on, method, your param
	costumerMockRepository.On("GetById", "1").Return(nil)
	result, err := costumerService.Get("1")

	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func TestGetDataCostumerSuccess(t *testing.T) {
	costumer := entity.Costumer{
		Id:       "101",
		Name:     "Kaguya Shinomiya",
		Username: "kaguyachanuww",
		Address: entity.DetailsAddress{
			City:     "Semarang",
			Province: "Central java",
			Country:  "Indonesia",
		},
	}
	costumerMockRepository.On("GetById", costumer.Id).Return(costumer)
	result, err := costumerService.Get("101")

	assert.Nil(t, err)
	assert.NotNil(t, result)

	require.Equal(t, result.Name, costumer.Name)
	require.Equal(t, result.Username, costumer.Username)
	require.Equal(t, result.Address, costumer.Address)

	fmt.Println(result)

}
