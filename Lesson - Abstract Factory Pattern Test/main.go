package main

import "fmt"

type AgeGroupInterface interface {
	MyName()
	MyAge(age int) int
}

type Adolescence struct{}

func (a *Adolescence) MyName() {
	fmt.Println("Loreng")
}

func (a *Adolescence) MyAge(age int) int {
	return age
}

type Adult struct{}

func (ad *Adult) MyName() {
	fmt.Println("Anchol")
}

func (ad *Adult) MyAge(age int) int {
	return age
}

type AdolescenceFactory struct{}

func (af *AdolescenceFactory) New() AgeGroupInterface {
	return &Adolescence{}
}

type AdultFactory struct{}

func (adf *AdultFactory) New() AgeGroupInterface {
	return &Adult{}
}

func main() {
	var af AdolescenceFactory
	var adf AdultFactory

	loreng := af.New()
	loreng.MyName()
	fmt.Println("My age is: ", loreng.MyAge(6))
	anchol := adf.New()
	anchol.MyName()
	fmt.Println("My age is: ", anchol.MyAge(27))
}
