package store

import (
	"context"
	"fmt"
	"sync"
)

// AsyncStore defines an interface for an asynchronous store
type AsyncStore[T any] interface {
	Create(ctx context.Context, key string, value T) error
	Read(ctx context.Context, key string) (T, error)
	Update(ctx context.Context, key string, value T) error
	Delete(ctx context.Context, key string) error
}

// MemoryStore is a concurrent in-memory implementation of AsyncStore
type MemoryStore[T any] struct {
	data sync.Map
}

// NewMemoryStore creates a new MemoryStore instance
func NewMemoryStore[T any]() *MemoryStore[T] {
	return &MemoryStore[T]{}
}

// Create inserts a new value asynchronously
func (s *MemoryStore[T]) Create(ctx context.Context, key string, value T) error {
	done := make(chan error, 1)
	go func() {
		if _, exists := s.data.Load(key); exists {
			done <- fmt.Errorf("key already exists")
			return
		}
		s.data.Store(key, value)
		done <- nil
	}()
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Read retrieves a value asynchronously
func (s *MemoryStore[T]) Read(ctx context.Context, key string) (T, error) {
	done := make(chan struct {
		value T
		err   error
	}, 1)

	go func() {
		val, exists := s.data.Load(key)
		if !exists {
			done <- struct {
				value T
				err   error
			}{err: fmt.Errorf("key not found")}
			return
		}
		done <- struct {
			value T
			err   error
		}{value: val.(T)}
	}()

	select {
	case result := <-done:
		return result.value, result.err
	case <-ctx.Done():
		var zero T
		return zero, ctx.Err()
	}
}

// Update modifies an existing value asynchronously
func (s *MemoryStore[T]) Update(ctx context.Context, key string, value T) error {
	done := make(chan error, 1)
	go func() {
		if _, exists := s.data.Load(key); !exists {
			done <- fmt.Errorf("key not found")
			return
		}
		s.data.Store(key, value)
		done <- nil
	}()
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Delete removes a value asynchronously
func (s *MemoryStore[T]) Delete(ctx context.Context, key string) error {
	done := make(chan error, 1)
	go func() {
		if _, exists := s.data.Load(key); !exists {
			done <- fmt.Errorf("key not found")
			return
		}
		s.data.Delete(key)
		done <- nil
	}()
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
