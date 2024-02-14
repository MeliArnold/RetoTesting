package repository

import (
	"app/internal"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/mock"
)

func TestSearchProducts(t *testing.T) {
	mockProductsMap := new(MockProductsMap)

	t.Run("success - query not set - return all products", func(t *testing.T) {
		mockProductsMap := new(MockProductsMap)
		mockProducts := map[int]internal.Product{
			1: {
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "This is product 1",
					Price:       9.99,
					SellerId:    123,
				},
			},
			2: {
				Id: 2,
				ProductAttributes: internal.ProductAttributes{
					Description: "This is product 2",
					Price:       19.99,
					SellerId:    234,
				},
			},
		}

		mockProductsMap.On("SearchProducts", internal.ProductQuery{}).Return(mockProducts, nil)

		result, err := mockProductsMap.SearchProducts(internal.ProductQuery{})

		mockProductsMap.AssertExpectations(t)

		assert.Len(t, result, 2)
		assert.Nil(t, err)
	})

	t.Run("success - query not set - return empty list if no products", func(t *testing.T) {
		mockProductsMap.On("SearchProducts", internal.ProductQuery{}).Return(make(map[int]internal.Product), nil)

		result, err := mockProductsMap.SearchProducts(internal.ProductQuery{})

		mockProductsMap.AssertExpectations(t)

		assert.Len(t, result, 0)
		assert.Nil(t, err)
	})

	t.Run("success - query set - return products that match the query", func(t *testing.T) {
		productQuery := internal.ProductQuery{Id: 1}

		mockProducts := map[int]internal.Product{
			1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "Product 1"}},
		}

		mockProductsMap.On("SearchProducts", productQuery).Return(mockProducts, nil)

		result, err := mockProductsMap.SearchProducts(productQuery)

		mockProductsMap.AssertExpectations(t)

		assert.Len(t, result, 1)
		assert.Nil(t, err)
	})

	t.Run("success - query set - return empty list if no products match the query", func(t *testing.T) {
		productQuery := internal.ProductQuery{Id: 3}

		mockProductsMap.On("SearchProducts", productQuery).Return(nil, nil)

		result, err := mockProductsMap.SearchProducts(productQuery)

		mockProductsMap.AssertExpectations(t)

		assert.Len(t, result, 0)
		assert.Nil(t, err)
	})
}
