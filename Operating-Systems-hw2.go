package main


import "fmt"
import "runtime"
import "sync"
import "sync/atomic"
import "time"


// Ticket Lock
// Assigns a unique num to each thread attempting to access the lock
// turn keeps track of which threads turn to enter critical section of the code
type TicketLock struct {
	ticket int32
	turn   int32
}

// Lock acquires lock by getting unique ticket num
func (lock *TicketLock) Lock() {
	// Get and increment the ticket counter
	myTurn := atomic.AddInt32(&lock.ticket, 1) - 1

	// Wait until its threads turn to enter the critical section
	for atomic.LoadInt32(&lock.turn) != myTurn {
		runtime.Gosched() // Avoid waiting too long
	}
}

// Unlock releases lock by incrementing turn val allowing next waiting thread to enter
func (lock *TicketLock) Unlock() {
	atomic.AddInt32(&lock.turn, 1)
}

// Constructor function to create a new TicketLock
func NewTicketLock() *TicketLock {
	return &TicketLock{}
}

// CAS Spin Lock
// CASSpinLock struct contains flag that indicates if the lock is accessed
type CASSpinLock struct {
	flag int32
}

// Lock accesses the lock using atomic CAS
func (lock *CASSpinLock) Lock() {
	// For loop to keep trying to acquire the lock until successful
	for !atomic.CompareAndSwapInt32(&lock.flag, 0, 1) {
		runtime.Gosched() // Yield CPU to avoid waiting too long
	}
}

// Unlock releases the lock by setting the flag back to 0.
func (lock *CASSpinLock) Unlock() {
	atomic.StoreInt32(&lock.flag, 0)
}

// Constructor function to create a new CASSpinLock
func NewCASSpinLock() *CASSpinLock {
	return &CASSpinLock{}
}

// Benchmarking Function

// benchmarkLock tests the lock performance by measuring time taken to
// Acquire and release the lock by by the goroutines
// lock instance to be tested 
// name identifies the lock type
// numThreads is num of goroutines/threads
// iterations is number of lock/unlocks per goroutine
func benchmarkLock(lock interface {
	Lock()
	Unlock()
}, name string, numThreads, iterations int) {
	var wg sync.WaitGroup // WaitGroup to sync goroutines
	start := time.Now()   // record start time of the benchmark

	// Start goroutines to acquire and release the lock
	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < iterations; j++ {
				lock.Lock()
				lock.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()                             // Wait for all goroutines to finish
	elapsed := time.Since(start)          // Calculate total time taken
	fmt.Printf("%s: %v\n", name, elapsed) // Print the time taken
}

// Main Function

func main() {
	numThreads := 10   // Num of goroutines working simultaneously
	iterations := 1000 // Num of lock/unlock per goroutine

	fmt.Println("Benchmarking locks with", numThreads, "goroutines and", iterations, "iterations per goroutine.")

	// Create instances of both locks
	ticketLock := NewTicketLock()
	casSpinLock := NewCASSpinLock()

	// Run benchmarks on each lock type
	benchmarkLock(ticketLock, "Ticket Lock", numThreads, iterations)
	benchmarkLock(casSpinLock, "CAS Spin Lock", numThreads, iterations)
}
