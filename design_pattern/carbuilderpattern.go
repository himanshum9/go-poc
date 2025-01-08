package main

type Car struct {
	Speed   int
	Mileage int
	Colour  string
	Brand   string
	Model   string
}

type CarBuilder struct {
	car Car
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

func (b *CarBuilder) SetColor(s string) *CarBuilder {
	b.car.Colour = s
	return b
}

func (b *CarBuilder) SetMileage(val int) *CarBuilder {
	b.car.Mileage = val
	return b
}

func (b *CarBuilder) SetBrand(val string) *CarBuilder {
	b.car.Brand = "test"
	return b
}

func (b *CarBuilder) build() Car {
	return b.car
}

// func main() {
// 	car := NewCarBuilder()
// 	car.SetColor("White").SetMileage(10).SetBrand("Toyota").build()
// 	fmt.Printf("Car: %+v\n", car)

// }
