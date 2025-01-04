package main

import "fmt"

type Vehicle interface {
	Mileage() float64
	AverageSpeed() float64
}

type Car struct {
	fuelEfficieny float64
	speed         float64
}

func (c Car) Mileage() float64 {
	return c.fuelEfficieny
}

func (c Car) AverageSpeed() float64 {
	return c.speed
}

func PrintMileage(v Vehicle) {
	fmt.Printf("Mileage of car is %.2f\n", v.Mileage())
	fmt.Printf("Average Speed of car is %.2f\n", v.AverageSpeed())
}

func main() {
	mycar := Car{fuelEfficieny: 10.0, speed: 200.1}
	// mycar.fuelEfficieny = 10.0
	// mycar.speed = 200.2
	PrintMileage(mycar)
	message := "He said, \"Hello!\"\nWelcome to the program."
	fmt.Println(message)
}
