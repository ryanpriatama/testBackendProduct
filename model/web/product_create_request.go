package web

type ProductCreateRequest struct {
	Name        string `validate:"required,max=300,min=1" json:"name"`
	Price       int64  `validate:"required,min=1" json:"price"`
	Description string `validate:"required,max=2000,min=1" json:"description"`
	Quantity    int    `validate:"required" json:"quantity"`
}
