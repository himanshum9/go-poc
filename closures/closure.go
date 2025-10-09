package main

//A closure in Go (or most languages) is a function that captures variables from its surrounding scope and can use or modify them even after that scope has ended.

import "fmt"

func counter() func() int {
	var count int = 0
	return func() int {
		count += 1
		return count
	}
}

func main() {
	inc := counter()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}
