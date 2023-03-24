package main

import (
	"net/http"
	"ryan-test-backend/app"
	"ryan-test-backend/controller"
	"ryan-test-backend/helper"
	"ryan-test-backend/middleware"
	"ryan-test-backend/repository"
	"ryan-test-backend/service"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	server := NewServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

func NewServer() *http.Server {
	db := app.NewDB()
	validate := validator.New()
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)

	router := httprouter.New()
	router.GET("/api/products", productController.GetProductsSorted)
	router.POST("/api/products", productController.Create)

	authMiddleware := middleware.NewAuthMiddleware(router)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}

	return &server
}
