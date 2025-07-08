package main

import "fmt"

type Pizza struct {
	Name        string
	Base        string
	BaseCost    float64
	ToppingCost float64
}

type PizzaPricer struct {
}

func (pp PizzaPricer) CalculatePrice(p Pizza) float64 {
	return p.BaseCost + p.ToppingCost
}

type PizzaRepository struct {
}

func (pp PizzaRepository) Save(p Pizza) {
	fmt.Printf("This %s is getting saved in db.\n", p.Name)
}

type PizzaPrinter struct{}

func (pp PizzaPrinter) Print(p Pizza, total float64) {
	fmt.Printf("Pizza: %s (%s)\nTotal: ₹%.2f\n", p.Name, p.Base, total)
}

func main() {
	pizza := Pizza{Name: "Paneer Pizza", Base: "Pan", BaseCost: 200, ToppingCost: 50}
	pricer := PizzaPricer{}
	total := pricer.CalculatePrice(pizza)
	printer := PizzaPrinter{}
	printer.Print(pizza, total)
	repo := PizzaRepository{}
	repo.Save(pizza)
}
