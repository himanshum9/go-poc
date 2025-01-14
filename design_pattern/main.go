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

	car := NewCarBuilder().SetColor("White").SetMileage(10).SetBrand("Toyota").build()
	// car.SetColor("Black").build()
	fmt.Printf("Car: %+v\n", car)

	p := FactoryMethod("A")
	fmt.Println(p.getName())

	instance := GetInstance(20)
	fmt.Println(instance)

	instance = GetInstance(40)
	fmt.Println(instance)
}
