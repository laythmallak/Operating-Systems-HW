// locked_queue.go - Two-Lock Concurrent Queue (Michael & Scott Figure 29.9)

package main

import (
	"sync"
)

// Node holds an integer value and points to the next node.
type Node struct {
	value int
	next  *Node
}

// LockedQueue uses two locks for concurrent enqueue and dequeue.
type LockedQueue struct {
	head     *Node
	tail     *Node
	headLock sync.Mutex
	tailLock sync.Mutex
}

// NewLockedQueue initializes the queue with a dummy node.
func NewLockedQueue() *LockedQueue {
	dummy := &Node{}
	return &LockedQueue{head: dummy, tail: dummy}
}

// Enqueue adds a new value to the end of the queue.
func (q *LockedQueue) Enqueue(value int) {
	node := &Node{value: value}

	// Lock tail, add node, update tail pointer
	q.tailLock.Lock()
	q.tail.next = node
	q.tail = node
	q.tailLock.Unlock()
}

// Dequeue removes and returns the first value from the queue.
func (q *LockedQueue) Dequeue() (int, bool) {
	q.headLock.Lock()
	defer q.headLock.Unlock()

	oldHead := q.head
	newHead := oldHead.next
	if newHead == nil {
		return 0, false // Empty queue
	}
	value := newHead.value
	q.head = newHead
	return value, true
}
