// Worker always read from shared channel which means in my workerpool pattern my worker will always reads data from channel.
package main

import (
	"fmt"
	"sync"
)

func workers(i int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range ch {
		fmt.Printf("Reading from Worker %d and channel value is %d \n", i, j)
	}
}

func WorkerMain() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go workers(i, ch, &wg)
	}
	for i := 0; i < 1000; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
}
