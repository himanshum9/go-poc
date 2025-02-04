package main

import (
	"fmt"
	"time"
)

func sendRequest(i int, c chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("This is the", i, "th line")
	c <- i
}

func findMillionPrimeSequential() {
	start := time.Now()
	ch := make(chan int)
	for i := 0; i < 100000; i++ {
		go sendRequest(i, ch)
	}

	for i := 0; i < 100000; i++ {
		<-ch
	}

	// for channel := range ch {
	// 	fmt.Println(channel)
	// }
	end := time.Now()
	diff := end.Sub(start)
	fmt.Println(diff)
}
