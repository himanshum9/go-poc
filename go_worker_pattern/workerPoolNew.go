package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

const (
	Tasks = 10000000
)

func worker(i int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	valcount := 0
	primecount := 0
	start := time.Now()
	for j := range ch {
		flag := false
		valcount++
		if j < 2 {
			continue
		}
		for k := 2; k <= int(math.Sqrt(float64(j))); k++ {
			if j%k == 0 {
				flag = true
				break
			}
		}
		if flag != true {
			primecount++
		}
	}
	end := time.Since(start)
	fmt.Println("Worker", i, "total time", end, "valcount", valcount, "Prime Count", primecount)
}

func WorkerPoolNew() {
	workerPool := 10

	start := time.Now()
	ch := make(chan int, 100)
	wg := sync.WaitGroup{}
	for i := 0; i < workerPool; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	for i := 0; i < Tasks; i++ {
		ch <- i
	}

	close(ch)
	wg.Wait()

	end := time.Since(start)
	fmt.Println(end)
}
