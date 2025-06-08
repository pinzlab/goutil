package store

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddAndGet(t *testing.T) {
	store := NewStore[string]()

	key := "123456"
	value := "test_value"
	ttl := 2 * time.Second

	store.Add(key, value, ttl)
	got, ok := store.Get(key)

	assert.True(t, ok, "Expected value to be present, but it was not found.")
	assert.Equal(t, value, got.Value, "Returned value does not match expected value.")
	assert.Greater(t, got.ExpiresAt, time.Now().Unix(), "ExpiresAt should be in the future.")
}

func TestAddWithTTL(t *testing.T) {
	store := NewStore[string]()

	key := "expired_key"
	value := "temp_value"
	ttl := 1 * time.Second

	store.Add(key, value, ttl)
	time.Sleep(2 * time.Second) // Wait for the key to expire

	_, ok := store.Get(key)
	assert.False(t, ok, "Expected value to be expired and removed, but it was still found.")
}

func TestDelete(t *testing.T) {
	store := NewStore[string]()

	key := "to_be_deleted"
	value := "some_value"
	ttl := 10 * time.Second

	store.Add(key, value, ttl)
	store.Delete(key)

	_, ok := store.Get(key)
	assert.False(t, ok, "Expected value to be deleted, but it was still found.")
}

func TestConcurrentAddAndGet(t *testing.T) {
	store := NewStore[string]()

	baseKey := "concurrent_key"
	value := "thread_safe_value"
	ttl := 3 * time.Second

	// Add 100 items concurrently
	for i := 0; i < 100; i++ {
		go store.Add(fmt.Sprintf("%s_%d", baseKey, i), value, ttl)
	}

	// Allow time for all goroutines to complete
	time.Sleep(500 * time.Millisecond)

	// Check if all keys are present and correct
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("%s_%d", baseKey, i)
		got, ok := store.Get(key)
		assert.True(t, ok, "Expected value for key %s, but it was not found.", key)
		assert.Equal(t, value, got.Value, "Unexpected value for key %s.", key)
	}
}
