package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/catmullet/go-workers"
)

var (
	count = make(map[string]int)
	mut   = sync.RWMutex{}
)

func main() {
	ctx := context.Background()
	t := time.Now()
	workerOne := workers.NewRunner(ctx, NewWorkerOne(), 50000).Start()
	workerTwo := workers.NewRunner(ctx, NewWorkerTwo(), 50000).InFrom(workerOne).Start()

	go func() {
		for i := 0; i < 100000; i++ {
			workerOne.Send(rand.Intn(100))
		}
		if err := workerOne.Wait(); err != nil {
			fmt.Println(err)
		}
	}()

	if err := workerTwo.Wait(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("worker_one", count["worker_one"])
	fmt.Println("worker_two", count["worker_two"])
	fmt.Println("finished")

	totalTime := time.Since(t).Milliseconds()
	fmt.Printf("total time %dms\n", totalTime)
}

type WorkerOne struct {
}
type WorkerTwo struct {
}

func NewWorkerOne() workers.Worker {
	return &WorkerOne{}
}

func NewWorkerTwo() workers.Worker {
	return &WorkerTwo{}
}

func (wo *WorkerOne) Work(in interface{}, out chan<- interface{}) error {
	var workerOne = "worker_one"
	mut.Lock()
	if val, ok := count[workerOne]; ok {
		count[workerOne] = val + 1
	} else {
		count[workerOne] = 1
	}
	mut.Unlock()

	total := in.(int) * 2
	fmt.Println("worker1", fmt.Sprintf("%d * 2 = %d", in.(int), total))
	out <- total
	return nil
}
func (wt *WorkerTwo) Work(in interface{}, out chan<- interface{}) error {
	var workerTwo = "worker_two"
	mut.Lock()
	if val, ok := count[workerTwo]; ok {
		count[workerTwo] = val + 1
	} else {
		count[workerTwo] = 1
	}
	mut.Unlock()

	totalFromWorkerOne := in.(int)
	fmt.Println("worker2", fmt.Sprintf("%d * 4 = %d", totalFromWorkerOne, totalFromWorkerOne*4))
	return nil
}
