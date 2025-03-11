package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Timeout occur before processing")
				return
			default:
				fmt.Println("Everything is right")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(4 * time.Second)
	fmt.Println("Main function executing")
}
