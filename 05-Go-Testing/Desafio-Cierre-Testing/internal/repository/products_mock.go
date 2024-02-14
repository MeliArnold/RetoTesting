package repository

import (
	"app/internal"
	"github.com/stretchr/testify/mock"
)

type MockProductsMap struct {
	mock.Mock
}

func (m *MockProductsMap) SearchProducts(q internal.ProductQuery) (map[int]internal.Product, error) {
	args := m.Called(q)

	if args.Get(0) == nil {
		// if the first argument is nil, return an empty map and the error
		return make(map[int]internal.Product), args.Error(1)
	}

	productMapArg, ok := args.Get(0).(map[int]internal.Product)

	if !ok {
		return nil, args.Error(1)
	}

	return productMapArg, args.Error(1)
}
