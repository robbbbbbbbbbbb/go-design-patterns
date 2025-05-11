package product_histories

import "time"

type ProductInventoryMovement struct {
	ProductName  string
	SoldQuantity int
	CreatedAt    *time.Time
}

func (ph *ProductInventoryMovement) New() *ProductInventoryMovement {
	now := time.Now()
	return &ProductInventoryMovement{
		CreatedAt: &now,
	}
}
