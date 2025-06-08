package store

import (
	"sync"
	"time"
)

// Value represents a generic value stored in the Store along with its expiration time.
type Value[D any] struct {
	Value     D     // The actual value stored
	ExpiresAt int64 // Unix timestamp (in seconds) when the value expires
}

// Store is a generic, thread-safe in-memory key-value store
// that supports automatic expiration of entries based on TTL.
type Store[D any] struct {
	store map[string]Value[D] // Internal map to hold key-value pairs
	mutex sync.RWMutex        // Read-write mutex for concurrent access
}

// NewStore initializes and returns a new Store instance.
// The store is safe for concurrent use and ready for storing key-value pairs.
func NewStore[D any]() *Store[D] {
	return &Store[D]{store: make(map[string]Value[D])}
}

// Add inserts a new key-value pair into the store with a given TTL (time-to-live).
// The entry is automatically removed from the store after the TTL duration has passed.
func (s *Store[D]) Add(key string, value D, ttl time.Duration) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.store[key] = Value[D]{
		Value:     value,
		ExpiresAt: time.Now().Add(ttl).Unix(),
	}

	// Launch a goroutine to remove the item once TTL expires
	go func() {
		<-time.After(ttl)
		s.mutex.Lock()
		defer s.mutex.Unlock()
		delete(s.store, key)
	}()
}

// Get retrieves the stored Value associated with the given key.
// It returns the Value and a boolean indicating whether the key exists.
// Note: This does not check if the value is expired.
func (s *Store[D]) Get(key string) (Value[D], bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	value, ok := s.store[key]
	return value, ok
}

// Delete removes the key-value pair associated with the given key from the store.
func (s *Store[D]) Delete(key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.store, key)
}
