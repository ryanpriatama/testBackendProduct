package mocks

import (
	"context"
	"database/sql"
	"ryan-test-backend/model/domain"
	"ryan-test-backend/model/web"

	"github.com/stretchr/testify/mock"
)

type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) Create(ctx context.Context, request web.ProductCreateRequest) domain.Product {
	args := m.Called(ctx, request)
	return args.Get(0).(domain.Product)
}

func (m *MockProductService) GetProductsSorted(ctx context.Context, sort string) []domain.Product {
	args := m.Called(ctx, sort)
	return args.Get(0).([]domain.Product)
}

type MockProductRepository struct{}

func (m *MockProductRepository) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	return product
}

func (m *MockProductRepository) GetProductsSorted(ctx context.Context, tx *sql.Tx, nameSort string) []domain.Product {
	return []domain.Product{
		{Name: "Product A", Price: 1000, Description: "Description A", Quantity: 10},
		{Name: "Product B", Price: 2000, Description: "Description B", Quantity: 20},
		{Name: "Product C", Price: 3000, Description: "Description C", Quantity: 30},
	}
}
