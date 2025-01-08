package main

import "fmt"

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
	fmt.Printf("Extra Sauce: %+v\n", pizza.ExtraSauce)

	car := NewCarBuilder()
	car.SetColor("White").SetMileage(10).SetBrand("Toyota").build()
	fmt.Printf("Car: %+v\n", car)
}
