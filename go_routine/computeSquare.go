package main

import "fmt"

func computeSquare(j int, ch chan int) {
	result := j * j
	ch <- result
}

func computeSquareMain() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	ch := make(chan int)
	for _, j := range nums {
		go computeSquare(j, ch)
		result := <-ch
		fmt.Printf("Square of %d is %d\n", j, result)
	}
	go func() {
		close(ch)
	}()

	// for res := range ch {
	// 	fmt.Printf("Square of %d is %d\n", res[0], res[1])
	// }
}
