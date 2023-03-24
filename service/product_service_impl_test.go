package service

import (
	"context"
	"database/sql"
	"reflect"
	"regexp"
	"ryan-test-backend/model/domain"
	"ryan-test-backend/model/web"
	"ryan-test-backend/repository"
	"ryan-test-backend/test"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	//"gopkg.in/go-playground/assert.v1"
)

func TestNewProductService(t *testing.T) {
	type args struct {
		productRepository repository.ProductRepository
		DB                *sql.DB
		validate          *validator.Validate
	}
	tests := []struct {
		name string
		args args
		want ProductService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductService(tt.args.productRepository, tt.args.DB, tt.args.validate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductServiceImpl_Create(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer mockDB.Close()

	ctx := context.Background()
	productRepo := repository.NewProductRepository()

	productService := NewProductService(productRepo, mockDB, validator.New())

	productCreateRequest := web.ProductCreateRequest{
		Name:        "Product Name",
		Price:       1000,
		Description: "Product Description",
		Quantity:    5,
	}

	// Set expectation for insert
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO products")).
		WithArgs(productCreateRequest.Name, productCreateRequest.Price, productCreateRequest.Description, productCreateRequest.Quantity).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	productResponse := productService.Create(ctx, productCreateRequest)

	assert.Equal(t, "Product Name", productResponse.Name)
	assert.Equal(t, int64(1000), productResponse.Price)
	assert.Equal(t, "Product Description", productResponse.Description)
	assert.Equal(t, 5, productResponse.Quantity)
}

func TestProductServiceImpl_GetProductsSorted(t *testing.T) {
	// Create a mock database connection
	db, err := sql.Open("mysql", "root:rootadmin@tcp(localhost:3306)/testing_ryan_test_backend")
	test.TruncateCategory(db)
	if err != nil {
		t.Fatalf("failed to open mock database connection: %v", err)
	}
	defer db.Close()

	// Create a mock product repository
	repo := &MockProductRepository{}

	// Create a product service with the mock dependencies
	service := &ProductServiceImpl{
		ProductRepostiroy: repo,
		DB:                db,
		Validate:          nil,
	}

	// Define the test cases
	testCases := []struct {
		nameSort   string
		products   []domain.Product
		expected   []web.ProductResponse
		expectFail bool
	}{
		{
			nameSort: "ascending_name",
			products: []domain.Product{
				{Id: 1, Name: "A", Price: 1000, Description: "Product A", Quantity: 5},
				{Id: 2, Name: "B", Price: 2000, Description: "Product B", Quantity: 10},
			},
			expected: []web.ProductResponse{
				{Id: 1, Name: "A", Price: 1000, Description: "Product A", Quantity: 5},
				{Id: 2, Name: "B", Price: 2000, Description: "Product B", Quantity: 10},
			},
			expectFail: false,
		},
		{
			nameSort: "descending_name",
			products: []domain.Product{
				{Id: 2, Name: "B", Price: 2000, Description: "Product B", Quantity: 10},
				{Id: 1, Name: "A", Price: 1000, Description: "Product A", Quantity: 5},
			},
			expected: []web.ProductResponse{
				{Id: 2, Name: "B", Price: 2000, Description: "Product B", Quantity: 10},
				{Id: 1, Name: "A", Price: 1000, Description: "Product A", Quantity: 5},
			},
			expectFail: false,
		},
		{
			nameSort:   "invalid",
			products:   nil,
			expected:   nil,
			expectFail: true,
		},
	}

	// Run the test cases
	for _, testCase := range testCases {
		t.Run(testCase.nameSort, func(t *testing.T) {
			// Set up the mock repository to return the test data
			repo.On("GetProductsSorted", mock.Anything, mock.Anything, testCase.nameSort).
				Return(testCase.products).Once()

			// Call the method being tested
			result := service.GetProductsSorted(context.Background(), testCase.nameSort)

			// Check the result
			if testCase.expectFail {
				assert.Nil(t, result)
			} else {
				assert.Equal(t, testCase.expected, result)
			}

			// Verify that the mock repository was called as expected
			repo.AssertExpectations(t)
		})
	}
}

// MockProductRepository is a mock implementation of the ProductRepository interface
type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) GetProductsSorted(ctx context.Context, tx *sql.Tx, nameSort string) []domain.Product {
	args := m.Called(ctx, tx, nameSort)
	return args.Get(0).([]domain.Product)
}

func (m *MockProductRepository) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	args := m.Called(ctx, tx, domain.Product{})
	return args.Get(0).(domain.Product)
}
