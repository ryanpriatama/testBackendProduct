package web

type ProductResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
