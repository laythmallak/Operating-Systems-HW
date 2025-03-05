// lockfree_queue.go - Lock-Free Concurrent Queue (Michael & Scott 1996 PODC)

package main

import (
	"sync/atomic"
)

// LFNode holds an integer value and a pointer to the next node.
type LFNode struct {
	value int
	next  atomic.Pointer[LFNode]
}

// LockFreeQueue is a concurrent queue with atomic head/tail pointers.
type LockFreeQueue struct {
	head atomic.Pointer[LFNode]
	tail atomic.Pointer[LFNode]
}

// NewLockFreeQueue initializes the queue with a dummy node.
func NewLockFreeQueue() *LockFreeQueue {
	dummy := &LFNode{}
	q := &LockFreeQueue{}
	q.head.Store(dummy)
	q.tail.Store(dummy)
	return q
}

// Enqueue adds a new value to the queue.
func (q *LockFreeQueue) Enqueue(value int) {
	node := &LFNode{value: value}
	for {
		tail := q.tail.Load()
		next := tail.next.Load()

		if tail == q.tail.Load() { // Consistency check
			if next == nil { // Tail points to the last node
				if tail.next.CompareAndSwap(nil, node) { // Link new node
					q.tail.CompareAndSwap(tail, node) // Move tail forward
					return
				}
			} else {
				// Tail is lagging — move it forward
				q.tail.CompareAndSwap(tail, next)
			}
		}
	}
}

// Dequeue removes and returns a value from the queue.
func (q *LockFreeQueue) Dequeue() (int, bool) {
	for {
		head := q.head.Load()
		tail := q.tail.Load()
		next := head.next.Load()

		if head == q.head.Load() { // Consistency check
			if head == tail { // Empty queue or lagging tail
				if next == nil {
					return 0, false // Empty
				}
				q.tail.CompareAndSwap(tail, next) // Fix tail
			} else {
				value := next.value
				if q.head.CompareAndSwap(head, next) { // Move head forward
					return value, true
				}
			}
		}
	}
}
