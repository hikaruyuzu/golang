package repository

import (
	"testify-mock/entity"

	"github.com/stretchr/testify/mock"
)

// struct untuk menggunakan mock
type CostumerRepositoryMock struct {
	mock.Mock
}

// logic query, mock
func (costumerMock *CostumerRepositoryMock) GetById(id string) *entity.Costumer {
	args := costumerMock.Mock.Called(id)
	// karena parameternya 1 saja maka index get index 0
	if args.Get(0) == nil {
		return nil
	} else {
		data := args.Get(0).(entity.Costumer)
		return &data
	}

}
