package util

import (
	"sync"
)

// Item represents an element in the queue.
type Item[T any] struct {
	ID    string
	Value T
}

// Queue represents a thread-safe FILO queue.
type Queue[T any] struct {
	mu    sync.Mutex
	items []Item[T]
	index map[string]int
}

// NewQueue creates a new empty queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: make([]Item[T], 0),
		index: make(map[string]int),
	}
}

// Enqueue adds an item to the queue. If an item with the same ID exists, it is removed first.
func (q *Queue[T]) Enqueue(id string, value T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if pos, exists := q.index[id]; exists {
		q.removeAt(pos)
	}

	q.items = append(q.items, Item[T]{ID: id, Value: value})
	q.index[id] = len(q.items) - 1
}

// Dequeue removes and returns the last item in the queue (FILO behavior).
func (q *Queue[T]) Dequeue() (T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	var zero T
	if len(q.items) == 0 {
		return zero, false
	}

	lastIndex := len(q.items) - 1
	item := q.items[lastIndex]
	delete(q.index, item.ID)
	q.items = q.items[:lastIndex]

	return item.Value, true
}

// Remove deletes an item by ID from the queue if it exists.
func (q *Queue[T]) Remove(id string) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if pos, exists := q.index[id]; exists {
		q.removeAt(pos)
	}
}

// removeAt removes an item at a specific position.
func (q *Queue[T]) removeAt(pos int) {
	id := q.items[pos].ID
	delete(q.index, id)

	q.items = append(q.items[:pos], q.items[pos+1:]...)

	for i := pos; i < len(q.items); i++ {
		q.index[q.items[i].ID] = i
	}
}

// Size returns the number of items in the queue.
func (q *Queue[T]) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items)
}

// Contains checks if an item with the given ID exists in the queue.
func (q *Queue[T]) Contains(id string) bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	_, exists := q.index[id]
	return exists
}
