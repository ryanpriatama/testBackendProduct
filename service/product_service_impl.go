package service

import (
	"context"
	"database/sql"
	"ryan-test-backend/helper"
	"ryan-test-backend/model/domain"
	"ryan-test-backend/model/web"
	"ryan-test-backend/repository"

	"github.com/go-playground/validator"
)

//Implementation Business Logic or Service

type ProductServiceImpl struct {
	ProductRepostiroy repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepostiroy: productRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product := domain.Product{
		Name:        request.Name,
		Price:       request.Price,
		Description: request.Description,
		Quantity:    request.Quantity,
	}

	product = service.ProductRepostiroy.Save(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) GetProductsSorted(ctx context.Context, nameSort string) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepostiroy.GetProductsSorted(ctx, tx, nameSort)

	return helper.ToProductResponses(products)
}
