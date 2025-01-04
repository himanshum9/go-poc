package main

import (
	"fmt"
)

func computeSquare(j int, ch chan [2]int) {
	result := j * j
	ch <- [2]int{j, result}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	ch := make(chan [2]int)
	for _, j := range nums {
		go computeSquare(j, ch)
		result := <-ch
		fmt.Printf("Square of %d is %d\n", result[0], result[1])
	}
	go func() {
		close(ch)
	}()

	// for res := range ch {
	// 	fmt.Printf("Square of %d is %d\n", res[0], res[1])
	// }
}
