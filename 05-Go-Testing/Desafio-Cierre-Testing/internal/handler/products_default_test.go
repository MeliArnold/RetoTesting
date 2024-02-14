package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/repository"
	_ "app/internal/repository"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetProducts(t *testing.T) {
	tests := []struct {
		name           string
		mockArgs       []interface{}
		url            string
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "success search products - no query - full",
			mockArgs: []interface{}{
				map[int]internal.Product{
					1: {
						Id: 1,
						ProductAttributes: internal.ProductAttributes{
							Description: "product 1",
							Price:       1.1,
							SellerId:    1,
						},
					},
				},
				nil,
			},
			url:            "/products",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"success", "data": {"1": {"id":1,"description":"product 1","price":1.1,"seller_id":1}}}`,
		},
		{
			name: "success search products - no query - empty",
			mockArgs: []interface{}{
				map[int]internal.Product{},
				nil,
			},
			url:            "/products",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"success", "data":{}}`,
		},
		{
			name: "success search products - id query - full",
			mockArgs: []interface{}{
				map[int]internal.Product{
					1: {
						Id: 1,
						ProductAttributes: internal.ProductAttributes{
							Description: "product 1",
							Price:       1.1,
							SellerId:    1,
						},
					},
				},
				nil,
			},
			url:            "/products?id=1",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"success", "data":{"1":{"id":1,"description":"product 1","price":1.1,"seller_id":1}}}`,
		},
		{
			name: "success search products - id query - empty",
			mockArgs: []interface{}{
				map[int]internal.Product{},
				nil,
			},
			url:            "/products?id=1",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"success", "data":{}}`,
		},
		{
			name: "error search products",
			mockArgs: []interface{}{
				nil,
				errors.New("internal error"),
			},
			url:            "/products",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"message":"internal error", "status":"Internal Server Error"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/products", nil)

			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			mockRepository := new(repository.MockProductsMap)
			mockRepository.On("SearchProducts", mock.Anything).Return(test.mockArgs...)

			handler := handler.NewProductsDefault(mockRepository)

			handler.Get().ServeHTTP(rr, req)

			assert.Equal(t, test.expectedStatus, rr.Code)
			assert.JSONEq(t, test.expectedBody, strings.TrimSpace(rr.Body.String()))

			mockRepository.AssertExpectations(t)
		})
	}
}
