package main

import (
	"fmt"
	"sync"
)

func printNumbersWithChannel(n int, ch1 chan int, ch2 chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch1 {
		fmt.Printf("This is printed by %d worker | the number is:%d \n", n, i)
		ch2 <- true
	}

}

func printNNumbersWithChannel(n int) {
	const workerPool = 4
	ch1 := make(chan int, n)
	ch2 := make(chan bool)
	wg := &sync.WaitGroup{}

	for i := 0; i < workerPool; i++ {
		wg.Add(1)
		go printNumbersWithChannel(i, ch1, ch2, wg)
	}

	for i := 0; i < n; i++ {
		ch1 <- i
	}
	close(ch1)

	for i := 0; i < n; i++ {
		<-ch2
	}
	close(ch2)
	wg.Wait()

}
