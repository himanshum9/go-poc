package main

import "fmt"

func printOdd(num int, chanOdd chan int) {
	for i := 1; i < num; i += 2 {
		chanOdd <- i
	}
	close(chanOdd)
}

func printEven(num int, chanEven chan int) {
	for i := 0; i < num; i += 2 {
		chanEven <- i
	}
	close(chanEven)
}

func PrintOddEvenChannel(num int) {
	chanOdd := make(chan int)
	chanEven := make(chan int)
	go printOdd(num, chanOdd)
	go printEven(num, chanEven)

	for i := range chanOdd {
		fmt.Println(i)
	}

	for i := range chanEven {
		fmt.Println(i)
	}
}
