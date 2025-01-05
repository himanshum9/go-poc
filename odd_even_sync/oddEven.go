package main

import (
	"fmt"
	"sync"
)

func PrintOdd(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	if num%2 != 0 {
		fmt.Println("This is printed by Odd Function:", num)
	}
}

func PrintEven(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	if num%2 == 0 {
		fmt.Println("This is printed by even Function:", num)
	}
}

func PrintOddEven(num int) {
	ws := sync.WaitGroup{}

	for i := 0; i < num; i++ {
		ws.Add(1)
		go PrintOdd(i, &ws)
		go PrintEven(i, &ws)
	}

	ws.Wait()

}
