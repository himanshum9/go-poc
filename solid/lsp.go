package main

import "fmt"

type Bird struct {
	Name string
}

type Flyer interface {
	Fly() string
}

type Sparrow struct{}

func (s Sparrow) Fly() string {
	return "Sparrow is flying"
}

// Ostrich implements Flyer, but can't actually fly (violates LSP)
type Ostrich struct{}

func (o Ostrich) Fly() string {
	return "Ostrich can't fly!"
}

func BirdAction(b Bird, f Flyer) {
	fmt.Printf("%s: %s\n", b.Name, f.Fly())
}

func main() {
	sparrow := Bird{Name: "Sparrow"}
	ostrich := Bird{Name: "Ostrich"}

	fmt.Println("Bird actions:")
	BirdAction(sparrow, Sparrow{})
	BirdAction(ostrich, Ostrich{})
}
