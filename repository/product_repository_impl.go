package repository

import (
	"context"
	"database/sql"
	"ryan-test-backend/helper"
	"ryan-test-backend/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

//Implementation from contract interface ProductRepository

func (f *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "INSERT INTO products (name,price,description,quantity) VALUES(?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.Description, product.Quantity)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int(id)
	return product
}

func (f *ProductRepositoryImpl) GetProductsSorted(ctx context.Context, tx *sql.Tx, nameSort string) []domain.Product {
	var SQL string
	if nameSort == helper.DescendingName {
		SQL = "SELECT id,name,price,description,quantity FROM products ORDER BY name DESC"
	} else if nameSort == helper.AscendingName {
		SQL = "SELECT id,name,price,description,quantity FROM products ORDER BY name ASC"
	} else if nameSort == helper.HighPriceProduct {
		SQL = "SELECT id,name,price,description,quantity FROM products ORDER BY price DESC"
	} else if nameSort == helper.LowPriceProduct {
		SQL = "SELECT id,name,price,description,quantity FROM products ORDER BY price ASC"
	} else {
		SQL = "SELECT id,name,price,description,quantity FROM products ORDER BY id DESC"
	}

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Description, &product.Quantity)
		helper.PanicIfError(err)
		products = append(products, product)
	}
	return products
}
