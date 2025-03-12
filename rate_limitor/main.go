package main

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	tokens chan struct{}
}

func NewRateLimiter(rate int, burst int) *RateLimiter {
	rl := &RateLimiter{tokens: make(chan struct{}, burst)}

	// Fill the bucket initially
	for i := 0; i < burst; i++ {
		rl.tokens <- struct{}{}
	}

	// Add tokens at the given rate
	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(rate))
		defer ticker.Stop()
		for range ticker.C {
			select {
			case rl.tokens <- struct{}{}:
			default: // If full, drop the new token
			}
		}
	}()
	return rl
}

func (rl *RateLimiter) Allow() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

func main() {
	limiter := NewRateLimiter(5, 10) // 5 requests/sec, burst of 10

	for i := 0; i < 20; i++ {
		if limiter.Allow() {
			fmt.Println("Request allowed", i)
		} else {
			fmt.Println("Rate limited", i)
		}
		time.Sleep(200 * time.Millisecond)
	}
}
