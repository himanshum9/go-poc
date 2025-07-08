package main

import "fmt"

type Pizza struct {
	Name  string
	Price float64
}

type Discount interface {
	Apply(price float64) float64
}

type NoDiscount struct {
}

func (nd NoDiscount) Apply(price float64) float64 {
	return price
}

type FlatDiscount struct {
}

func (fd FlatDiscount) Apply(price float64) float64 {
	return price - 50
}

func FinalPrice(p Pizza, d Discount) float64 {
	return d.Apply(p.Price)
}

func main() {
	pizza := Pizza{Name: "Panner", Price: 200}
	fmt.Println("The discount on this pizza is:", FinalPrice(pizza, NoDiscount{}))
	fmt.Println("The discount on this pizza is:", FinalPrice(pizza, FlatDiscount{}))
}
