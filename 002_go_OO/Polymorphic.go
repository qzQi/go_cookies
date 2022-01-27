// package object
package main

import (
	"fmt"
	// "image/color"
)

type AnimalIF interface {
	Sleep()
	GetColor()
	GetType() string
}

type Cat struct {
	mycolor string
}

func (cat *Cat) Sleep() {
	fmt.Println("cat is sleeping")
}

func (cat *Cat) GetColor() {
	fmt.Printf("a cat with %s color\n", cat.mycolor)
}

type ChinaCat struct {
	Cat
	variety string
}

func (ccat *ChinaCat) GetType() string {
	return "cat from China"
}

func show(i AnimalIF) {
	fmt.Println(i.GetType())
	i.Sleep()
	i.GetColor()
}

func main() {
	var animal AnimalIF
	cBlack := Cat{"black"}
	animal = &ChinaCat{cBlack, "China"}

	show(animal)

	fmt.Println("outside print")

	// main结束后也不一定被调度，然后就终止了
	// 还是cpp里面比较好，countdown latch
	go fmt.Println("inner routine print")

	fmt.Println("in outside again")
}
