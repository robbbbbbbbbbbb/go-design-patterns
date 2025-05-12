package main

import "fmt"

type Animal interface {
	Says()
	LikesWater() bool
}

type Dog struct{}

func (d *Dog) Says() {
	fmt.Println("Woof")
}

func (d *Dog) LikesWater() bool {
	return true
}

type Cat struct{}

func (c *Cat) Says() {
	fmt.Println("Meow")
}

func (c *Cat) LikesWater() bool {
	return false
}

type DogFactory struct{}
type CatFactory struct{}

func (df *DogFactory) New() Animal {
	return &Dog{}
}

func (cf *CatFactory) New() Animal {
	return &Cat{}
}

func main() {
	dogFactory := DogFactory{}
	catFactory := CatFactory{}

	dog := dogFactory.New()
	dog.Says()
	fmt.Println("A dog likes water: ", dog.LikesWater())

	cat := catFactory.New()
	cat.Says()
	fmt.Println("A Cat likes water: ", cat.LikesWater())
}
