package store

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestMemoryStore tests the asynchronous memory store
func TestMemoryStore(t *testing.T) {
	store := NewMemoryStore[string]()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Test Create
	err := store.Create(ctx, "key1", "Hello, World!")
	assert.NoError(t, err, "Create should not return an error")

	// Test Read
	value, err := store.Read(ctx, "key1")
	assert.NoError(t, err, "Read should not return an error")
	assert.Equal(t, "Hello, World!", value, "Read value should match created value")

	// Test Create Duplicate (should fail)
	err = store.Create(ctx, "key1", "Another Value")
	assert.Error(t, err, "Create with duplicate key should return an error")

	// Test Update
	err = store.Update(ctx, "key1", "Updated Value")
	assert.NoError(t, err, "Update should not return an error")

	// Test Read after Update
	value, err = store.Read(ctx, "key1")
	assert.NoError(t, err, "Read after update should not return an error")
	assert.Equal(t, "Updated Value", value, "Updated value should match")

	// Test Delete
	err = store.Delete(ctx, "key1")
	assert.NoError(t, err, "Delete should not return an error")

	// Test Read after Delete (should fail)
	_, err = store.Read(ctx, "key1")
	assert.Error(t, err, "Read after delete should return an error")
}
