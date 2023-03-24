package helper

import (
	"ryan-test-backend/model/domain"
	"testing"
)

func TestToProductResponse(t *testing.T) {
	// Arrange
	product := domain.Product{
		Id:          1,
		Name:        "Product 1",
		Price:       1000,
		Description: "This is product 1",
		Quantity:    10,
	}

	// Act
	response := ToProductResponse(product)

	// Assert
	if response.Id != product.Id {
		t.Errorf("Incorrect response id, got: %d, want: %d", response.Id, product.Id)
	}
	if response.Name != product.Name {
		t.Errorf("Incorrect response name, got: %s, want: %s", response.Name, product.Name)
	}
	if response.Price != product.Price {
		t.Errorf("Incorrect response price, got: %d, want: %d", response.Price, product.Price)
	}
	if response.Description != product.Description {
		t.Errorf("Incorrect response description, got: %s, want: %s", response.Description, product.Description)
	}
	if response.Quantity != product.Quantity {
		t.Errorf("Incorrect response quantity, got: %d, want: %d", response.Quantity, product.Quantity)
	}
}

func TestToProductResponses(t *testing.T) {
	// Arrange
	products := []domain.Product{
		{
			Id:          1,
			Name:        "Product 1",
			Price:       1000,
			Description: "This is product 1",
			Quantity:    10,
		},
		{
			Id:          2,
			Name:        "Product 2",
			Price:       2000,
			Description: "This is product 2",
			Quantity:    20,
		},
	}

	// Act
	responses := ToProductResponses(products)

	// Assert
	if len(responses) != len(products) {
		t.Errorf("Incorrect number of responses, got: %d, want: %d", len(responses), len(products))
	}

	for i, response := range responses {
		if response.Id != products[i].Id {
			t.Errorf("Incorrect response id, got: %d, want: %d", response.Id, products[i].Id)
		}
		if response.Name != products[i].Name {
			t.Errorf("Incorrect response name, got: %s, want: %s", response.Name, products[i].Name)
		}
		if response.Price != products[i].Price {
			t.Errorf("Incorrect response price, got: %d, want: %d", response.Price, products[i].Price)
		}
		if response.Description != products[i].Description {
			t.Errorf("Incorrect response description, got: %s, want: %s", response.Description, products[i].Description)
		}
		if response.Quantity != products[i].Quantity {
			t.Errorf("Incorrect response quantity, got: %d, want: %d", response.Quantity, products[i].Quantity)
		}
	}
}
