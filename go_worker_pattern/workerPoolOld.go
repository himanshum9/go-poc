package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task struct {
	ID   int
	Name string
}

type Result struct {
	Taskid int
	Output string
}

func workerPoolOld() {
	fmt.Println("Starting of the main function")
	tasks := []Task{
		{1, "1st task"},
		{2, "2nd Task"},
		{3, "3rd Task"},
		// {4, "4th Task"},
		// {5, "5th Task"},
	}
	taskChan := make(chan Task, len(tasks))
	resChan := make(chan Result, len(tasks))
	numWorker := 3

	wg := sync.WaitGroup{}

	for i := 0; i < numWorker; i++ {
		wg.Add(1)
		go workerOld(i, taskChan, resChan, &wg)
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	for _, task := range tasks {
		taskChan <- task
	}
	close(taskChan)

	for result := range resChan {
		fmt.Printf("Task %d is completed:%s\n", result.Taskid, result.Output)
	}

	fmt.Println("All tasks are completed now,End of Project.")

}

func workerOld(id int, tasks chan Task, results chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worked %d started task:%s\n", id, task.Name)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		output := fmt.Sprintf("Worker %d finished its task:%s\n", id, task.Name)
		results <- Result{task.ID, output}
	}
}
