package main

import "fmt"

// Pizza defines the structure of the pizza
type Pizza struct {
	Size       string
	Crust      string
	Cheese     string
	Toppings   []string
	ExtraSauce bool
}

// PizzaBuilder provides methods to build a pizza
type PizzaBuilder struct {
	size       string
	crust      string
	cheese     string
	toppings   []string
	extraSauce bool
}

// NewPizzaBuilder creates a new instance of PizzaBuilder
func NewPizzaBuilder() *PizzaBuilder {
	return &PizzaBuilder{}
}

// SetSize sets the size of the pizza
func (b *PizzaBuilder) SetSize(size string) *PizzaBuilder {
	b.size = size
	return b
}

// SetCrust sets the crust type
func (b *PizzaBuilder) SetCrust(crust string) *PizzaBuilder {
	b.crust = crust
	return b
}

// SetCheese sets the cheese type
func (b *PizzaBuilder) SetCheese(cheese string) *PizzaBuilder {
	b.cheese = cheese
	return b
}

// AddTopping adds a topping to the pizza
func (b *PizzaBuilder) AddTopping(topping string) *PizzaBuilder {
	b.toppings = append(b.toppings, topping)
	return b
}

// SetExtraSauce sets the extra sauce option
func (b *PizzaBuilder) SetExtraSauce(extraSauce bool) *PizzaBuilder {
	b.extraSauce = extraSauce
	return b
}

// Build assembles the pizza and returns it
func (b *PizzaBuilder) Build() Pizza {
	return Pizza{
		Size:       b.size,
		Crust:      b.crust,
		Cheese:     b.cheese,
		Toppings:   b.toppings,
		ExtraSauce: b.extraSauce,
	}
}

// Main function to demonstrate the Builder Pattern
func main() {
	// Create a new pizza using the builder
	pizza := NewPizzaBuilder().
		SetSize("Large").
		SetCrust("Thin").
		SetCheese("Mozzarella").
		AddTopping("Pepperoni").
		AddTopping("Olives").
		SetExtraSauce(true).
		Build()

	fmt.Println("Your Pizza Order:")
	fmt.Printf("Size: %s\n", pizza.Size)
	fmt.Printf("Crust: %s\n", pizza.Crust)
	fmt.Printf("Cheese: %s\n", pizza.Cheese)
	fmt.Printf("Toppings: %v\n", pizza.Toppings)
	fmt.Printf("Extra Sauce: %v\n", pizza.ExtraSauce)
}
