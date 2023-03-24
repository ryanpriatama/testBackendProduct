package repository

import (
	"context"
	"database/sql"
	"ryan-test-backend/model/domain"
)

//Create data access layer to domain (repository)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	GetProductsSorted(ctx context.Context, tx *sql.Tx, nameSort string) []domain.Product
}
