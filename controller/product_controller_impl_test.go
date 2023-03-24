package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"ryan-test-backend/mocks"
	"ryan-test-backend/model/web"
	"ryan-test-backend/repository"
	"ryan-test-backend/service"
	"ryan-test-backend/test"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/mock"
	"gopkg.in/go-playground/assert.v1"
)

func TestProductController_Create(t *testing.T) {
	// Prepare mock ProductService
	mockProductService := &mocks.MockProductService{}
	mockProduct := web.ProductResponse{
		Id:          1,
		Name:        "test",
		Price:       100,
		Description: "test product",
		Quantity:    10,
	}
	mockProductService.On("Create", mock.Anything, mock.Anything).Return(mockProduct)

	// Prepare request body
	requestBody := web.ProductCreateRequest{
		Name:        "test",
		Price:       100,
		Description: "test product",
		Quantity:    10,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	// Prepare request
	req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	db := test.SetupTestDB()
	test.TruncateCategory(db)
	validate := validator.New()
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)

	// Prepare controller and router
	controller := &ProductControllerImpl{
		ProductService: productService,
	}
	router := httprouter.New()
	router.POST("/products", controller.Create)

	// Make request and check response
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var response web.WebResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	respDataMap := response.Data.(map[string]interface{})

	ResponseData := web.ProductResponse{
		Id:          int(respDataMap["id"].(float64)),
		Name:        respDataMap["name"].(string),
		Price:       int64(respDataMap["price"].(float64)),
		Description: respDataMap["description"].(string),
		Quantity:    int(respDataMap["quantity"].(float64)),
	}
	assert.Equal(t, mockProduct, ResponseData)
}

func TestProductControllerImpl_GetProductsSorted(t *testing.T) {
	db := test.SetupTestDB()
	test.TruncateCategory(db)
	validate := validator.New()
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)

	// Create a new product controller
	controller := &ProductControllerImpl{
		ProductService: productService,
	}

	// Create a new HTTP request with a sort query parameter
	req := httptest.NewRequest("GET", "/products?sort=name", nil)

	// Create a new HTTP response recorder
	rr := httptest.NewRecorder()

	// Create a new httprouter Params object
	params := httprouter.Params{}

	// Call the GetProductsSorted function
	controller.GetProductsSorted(rr, req, params)

	// Check the HTTP status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Check the HTTP response body
	expected := `{"code":200,"status":"OK","data":null}`
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("Expected body %s, but got %s", expected, rr.Body.String())
	}
}

func TestNewProductController(t *testing.T) {
	db := test.SetupTestDB()
	test.TruncateCategory(db)
	validate := validator.New()
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := NewProductController(productService)

	type args struct {
		productService service.ProductService
	}
	tests := []struct {
		name string
		args args
		want ProductController
	}{
		{
			name: "default",
			args: args{
				productService: productService,
			},
			want: productController,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductController(tt.args.productService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductController() = %v, want %v", got, tt.want)
			}
		})
	}
}
