// Scenario: Build a dispatcher that assigns a list of jobs to workers concurrently.
// Requirements:
// Spawn N worker goroutines.
// Distribute M jobs through a channel.
// Each worker prints when it starts and finishes a job.
// Simulate processing time with time.Sleep.
// Use sync.WaitGroup to wait for all jobs to finish.
// Ensure no job runs more than once.
// Example Input:
// jobs := []Job{
//     {ID: 1, Duration: 2 * time.Second},
//     {ID: 2, Duration: 1 * time.Second},
//     {ID: 3, Duration: 3 * time.Second},
//     // ...
// }
// Typical Output (order may vary):
// Worker 1 started job 2
// Worker 3 started job 1
// Worker 2 started job 3
// Worker 1 finished job 2
// Worker 3 finished job 1
// Worker 2 finished job 3
// All jobs completed.

package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID       int
	Duration time.Duration
}

func Worker(i int, ch chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Printf("Worker %d started job %d \n", i, val.ID)
		time.Sleep(val.Duration)
		fmt.Printf("Worker %d finished job %d \n", i, val.ID)
	}
}

func main() {
	ch := make(chan Job)
	wg := sync.WaitGroup{}
	jobs := []Job{
		{ID: 1, Duration: 2 * time.Second},
		{ID: 2, Duration: 1 * time.Second},
		{ID: 3, Duration: 3 * time.Second},
		{ID: 4, Duration: 4 * time.Second},
		{ID: 5, Duration: 5 * time.Second},
		{ID: 6, Duration: 6 * time.Second},
	}
	for i := 1; i < 4; i++ {
		wg.Add(1)
		go Worker(i, ch, &wg)
	}
	for _, j := range jobs {
		ch <- j
	}
	close(ch)
	wg.Wait()
	fmt.Println("All jobs completed")
}
