package product_histories

import "time"

type ProductHistory struct {
	OldProductName string
	NewProductName string
	OldQuantity    int
	NewQuantity    int
	OldPrice       float64
	NewPrice       float64
	OldDescription string
	NewDescription string
	CreatedAt      *time.Time
}

func (ph *ProductHistory) New() *ProductHistory {
	now := time.Now()
	return &ProductHistory{
		CreatedAt: &now,
	}
}
