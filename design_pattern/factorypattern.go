package main

type Product interface {
	getName() string
}

type ProductA struct{}

func (a ProductA) getName() string {
	return "Product A"
}

type ProductB struct{}

func (b ProductB) getName() string {
	return "Product B"
}

func FactoryMethod(s string) Product {
	switch s {
	case "A":
		return ProductA{}
	case "B":
		return ProductB{}
	default:
		return nil
	}
}
