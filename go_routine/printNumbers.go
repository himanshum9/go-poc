package main

import (
	"fmt"
	"sync"
)

func printNumbers(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("printed by %d thread: %d \n", n, n)

}

func printNNumbers(n int) {
	wg := &sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go printNumbers(i, wg)
	}
	wg.Wait()
}
