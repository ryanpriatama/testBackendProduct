package repository

import (
	"context"
	"database/sql"
	"ryan-test-backend/model/domain"
	"ryan-test-backend/test"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestProductRepositoryImpl_Save(t *testing.T) {
	repo := NewProductRepository()

	// setup database connection
	db, err := sql.Open("mysql", "root:rootadmin@tcp(localhost:3306)/testing_ryan_test_backend")
	test.TruncateCategory(db)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// create transaction
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	// create product to be saved
	product := domain.Product{
		Name:        "Test Product",
		Price:       1000,
		Description: "Test Product Description",
		Quantity:    10,
	}

	// save product
	ctx := context.Background()
	savedProduct := repo.Save(ctx, tx, product)

	// check if product was saved correctly
	if savedProduct.Id == 0 {
		t.Errorf("Expected saved product ID to be non-zero")
	}

	// get products sorted by name in ascending order
	productsByNameAsc := repo.GetProductsSorted(ctx, tx, "ascending_name")

	// check if products were retrieved correctly
	if len(productsByNameAsc) == 0 {
		t.Errorf("Expected non-zero number of products sorted by name in ascending order")
	}

	// get products sorted by name in descending order
	productsByNameDesc := repo.GetProductsSorted(ctx, tx, "descending_name")

	// check if products were retrieved correctly
	if len(productsByNameDesc) == 0 {
		t.Errorf("Expected non-zero number of products sorted by name in descending order")
	}

	// get products sorted by price in ascending order
	productsByPriceAsc := repo.GetProductsSorted(ctx, tx, "low_price_product")

	// check if products were retrieved correctly
	if len(productsByPriceAsc) == 0 {
		t.Errorf("Expected non-zero number of products sorted by price in ascending order")
	}

	// get products sorted by price in descending order
	productsByPriceDesc := repo.GetProductsSorted(ctx, tx, "high_price_product")

	// check if products were retrieved correctly
	if len(productsByPriceDesc) == 0 {
		t.Errorf("Expected non-zero number of products sorted by price in descending order")
	}
}
