package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID int
}

type Pool struct {
	PrgChannel chan Task
	PoolSize   int
	Wg         *sync.WaitGroup
}

func (pool *Pool) Worker() {
	for task := range pool.PrgChannel {
		task.Run()
	}
	pool.Wg.Done()
}

func (task *Task) Run() {
	time.Sleep(time.Second * 1)
	fmt.Printf("Processing task %d\n", task.ID)
}

func main() {
	start := time.Now()
	numberOfTasks := 10
	// init Porgram
	pool := Pool{PoolSize: 5}
	pool.Wg = &sync.WaitGroup{}
	pool.Wg.Add(pool.PoolSize)
	pool.PrgChannel = make(chan Task, numberOfTasks)
	for range pool.PoolSize {
		go pool.Worker()
	}
	// Load tasks
	for i := range numberOfTasks {
		pool.PrgChannel <- Task{ID: i}
	}
	close(pool.PrgChannel)
	pool.Wg.Wait()
	end := time.Now()
	timeTaken := end.Sub(start).Seconds()
	fmt.Println("Time Taken: ", timeTaken)
}
