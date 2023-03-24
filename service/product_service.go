package service

import (
	"context"
	"ryan-test-backend/model/web"
)

//contract interface business logic / service

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse
	GetProductsSorted(ctx context.Context, nameSort string) []web.ProductResponse
}
