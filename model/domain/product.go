package domain

//Domain or entity is representation field on dabase

type Product struct {
	Id          int
	Name        string
	Price       int64
	Description string
	Quantity    int
}
