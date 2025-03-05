// benchmark.go - Compare performance of Locked vs Lock-Free Queue

package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numProducers = 4
	numConsumers = 4
	operations   = 100000
)

func benchmarkQueue(name string, enqueue func(int), dequeue func() (int, bool)) {
	var wg sync.WaitGroup
	start := time.Now()

	// Start producers
	for i := 0; i < numProducers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < operations; j++ {
				enqueue(id*operations + j)
			}
		}(i)
	}

	// Start consumers
	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < operations; j++ {
				for {
					if _, ok := dequeue(); ok {
						break
					}
					time.Sleep(time.Microsecond)
				}
			}
		}()
	}

	wg.Wait()
	duration := time.Since(start)
	totalOps := numProducers*operations + numConsumers*operations
	fmt.Printf("[%s] Total ops: %d, Time: %v, Throughput: %.2f ops/sec\n",
		name, totalOps, duration, float64(totalOps)/duration.Seconds())
}

func main() {
	// Locked Queue Benchmark
	lockedQueue := NewLockedQueue()
	benchmarkQueue("LockedQueue", lockedQueue.Enqueue, lockedQueue.Dequeue)

	// Lock-Free Queue Benchmark
	lockFreeQueue := NewLockFreeQueue()
	benchmarkQueue("LockFreeQueue", lockFreeQueue.Enqueue, lockFreeQueue.Dequeue)
}
