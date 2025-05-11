package products

import "time"

type Product struct {
	Name        string
	Quantity    int
	Price       float64
	Description string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

func (p *Product) New() *Product {
	now := time.Now()
	return &Product{
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}
