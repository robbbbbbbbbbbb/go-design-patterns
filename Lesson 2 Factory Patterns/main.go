package main

import (
	"factory-patterns/products"
	"fmt"
)

func main() {
	fp := products.Product{}

	product := fp.New()

	fmt.Println("my first product was created at:", product.CreatedAt.UTC())
}
