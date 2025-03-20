package util

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue[string]()

	q.Enqueue("1", "first")
	q.Enqueue("2", "second")
	if q.Size() != 2 {
		t.Errorf("expected size 2, got %d", q.Size())
	}

	if !q.Contains("1") || !q.Contains("2") {
		t.Errorf("expected items to be in queue")
	}

	q.Enqueue("1", "updated-first")
	if q.Size() != 2 {
		t.Errorf("expected size 2 after re-enqueue, got %d", q.Size())
	}

	item, ok := q.Dequeue()
	if !ok || item != "updated-first" {
		t.Errorf("expected 'updated-first', got %v", item)
	}

	item, ok = q.Dequeue()
	if !ok || item != "second" {
		t.Errorf("expected 'second', got %v", item)
	}

	if q.Size() != 0 {
		t.Errorf("expected size 0 after dequeue, got %d", q.Size())
	}

	// Test removing an item
	q.Enqueue("3", "third")
	q.Enqueue("4", "fourth")
	q.Remove("3")

	if q.Contains("3") {
		t.Errorf("expected item '3' to be removed")
	}

	if q.Size() != 1 {
		t.Errorf("expected size 1 after removal, got %d", q.Size())
	}
}
