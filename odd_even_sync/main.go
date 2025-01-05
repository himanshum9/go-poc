package main

import "time"

func main() {
	// PrintOddEven(20)
	// PrintOddEvenChannel(20)
	SyncOddEvenChannel(20)
	time.Sleep(9000)
}
