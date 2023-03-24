package mocks

import (
	"context"
	"ryan-test-backend/model/domain"
	"ryan-test-backend/model/web"
	"testing"

	"github.com/stretchr/testify/mock"
	"gopkg.in/go-playground/assert.v1"
)

func TestMockProductService_Create(t *testing.T) {
	mockProduct := domain.Product{
		Id:          1,
		Name:        "Product 1",
		Price:       1000,
		Description: "Description 1",
		Quantity:    10,
	}
	mockProductService := new(MockProductService)
	mockProductService.On("Create", mock.Anything, mock.Anything).Return(mockProduct)

	ctx := context.Background()
	request := web.ProductCreateRequest{
		Name:        "Product 1",
		Price:       1000,
		Description: "Description 1",
		Quantity:    10,
	}
	product := mockProductService.Create(ctx, request)

	assert.Equal(t, mockProduct, product)
	mockProductService.AssertExpectations(t)
}

func TestMockProductService_GetProductsSorted(t *testing.T) {
	mockProducts := []domain.Product{
		{
			Id:          1,
			Name:        "Product 1",
			Price:       1000,
			Description: "Description 1",
			Quantity:    10,
		},
		{
			Id:          2,
			Name:        "Product 2",
			Price:       2000,
			Description: "Description 2",
			Quantity:    20,
		},
	}
	mockProductService := new(MockProductService)
	mockProductService.On("GetProductsSorted", mock.Anything, mock.Anything).Return(mockProducts)

	ctx := context.Background()
	products := mockProductService.GetProductsSorted(ctx, "price")

	assert.Equal(t, mockProducts, products)
	mockProductService.AssertExpectations(t)
}

func TestMockProductRepository_Save(t *testing.T) {
	repo := &MockProductRepository{}
	product := domain.Product{Name: "Product D", Price: 4000, Description: "Description D", Quantity: 40}
	savedProduct := repo.Save(context.Background(), nil, product)

	if savedProduct != product {
		t.Errorf("Save() = %v, want %v", savedProduct, product)
	}
}

func TestMockProductRepository_GetProductsSorted(t *testing.T) {
	repo := &MockProductRepository{}
	nameSort := "asc"
	expectedProducts := []domain.Product{
		{Name: "Product A", Price: 1000, Description: "Description A", Quantity: 10},
		{Name: "Product B", Price: 2000, Description: "Description B", Quantity: 20},
		{Name: "Product C", Price: 3000, Description: "Description C", Quantity: 30},
	}

	products := repo.GetProductsSorted(context.Background(), nil, nameSort)

	if len(products) != len(expectedProducts) {
		t.Errorf("GetProductsSorted() returned %d products, want %d", len(products), len(expectedProducts))
	}

	for i, p := range products {
		if p != expectedProducts[i] {
			t.Errorf("GetProductsSorted()[%d] = %v, want %v", i, p, expectedProducts[i])
		}
	}
}
