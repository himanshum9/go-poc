package main

import (
	"fmt"
	"sync"
)

func printOddSync(num int, chanOdd chan string, chanEven chan string, printChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= num; i += 2 {
		_, ok := <-chanOdd
		if !ok {
			return
		}
		printChan <- i
		if i == num {
			continue
		}
		chanEven <- "now"
	}
}

func printEvenSync(num int, chanOdd chan string, chanEven chan string, printChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= num; i += 2 {
		_, ok := <-chanEven
		if !ok {
			return
		}
		printChan <- i
		if i == num {
			continue
		}
		chanOdd <- "now"
	}
}

func SyncOddEvenChannel(num int) {
	chanOdd := make(chan string)  // Buffer size 1
	chanEven := make(chan string) // Buffer size 1
	printChan := make(chan int)   // Buffer size num

	wg := sync.WaitGroup{}
	wg.Add(2)

	go printOddSync(num, chanOdd, chanEven, printChan, &wg)
	go printEvenSync(num, chanOdd, chanEven, printChan, &wg)

	// Start with odd number
	chanOdd <- "now"

	// Wait for completion in a separate goroutine
	go func() {
		wg.Wait()
		close(printChan)
	}()

	// Print the numbers as they come
	for i := range printChan {
		fmt.Println(i)
	}

	// Clean up
	close(chanOdd)
	close(chanEven)
}

// func main() {
//     Syncoddevenchannel(10)
// }
