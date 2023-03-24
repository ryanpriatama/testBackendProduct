package main

import (
	"database/sql"
	"net/http"
	"os"
	"ryan-test-backend/controller"
	"ryan-test-backend/helper"
	"ryan-test-backend/mocks"
	"ryan-test-backend/service"
	"testing"
	"time"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func TestNewServer(t *testing.T) {
	// create a mock db and product repository
	mockDB := &sql.DB{}
	mockProductRepository := &mocks.MockProductRepository{}

	// create a new product service with the mock dependencies
	productService := service.NewProductService(mockProductRepository, mockDB, validator.New())

	// create a new product controller with the product service
	productController := controller.NewProductController(productService)

	// create a new router and register the product controller routes
	router := httprouter.New()
	router.GET("/api/products", productController.GetProductsSorted)
	router.POST("/api/products", productController.Create)

	// create a new auth middleware with the router
	//authMiddleware := middleware.NewAuthMiddleware(router)

	// create a new server with the auth middleware
	server := NewServer()

	// assert that the server has the correct properties
	if server.Addr != "localhost:3000" {
		t.Errorf("Expected server address to be 'localhost:3000', but got '%s'", server.Addr)
	}
	// if server.Handler != authMiddleware {
	// 	t.Errorf("Expected server handler to be '%T', but got '%T'", authMiddleware, server.Handler)
	// }
}

func TestMain(m *testing.M) {
	go func() {
		// Run main() function
		main()
	}()
	// Wait for server to start
	time.Sleep(1 * time.Second)

	// Run tests
	code := m.Run()

	// Clean up
	_, err := http.Get("http://localhost:3000/shutdown")
	helper.PanicIfError(err)

	// Wait for server to shut down
	time.Sleep(1 * time.Second)

	// Exit with test code
	os.Exit(code)
}
