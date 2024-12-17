package util

import (
	"sync"
)

type SyncMap[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}

// NewSyncMap creates a new generic sync map.
func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{
		m: make(map[K]V),
	}
}

// Load returns the value for a key, and a boolean indicating if the key was found.
func (sm *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, ok = sm.m[key]
	return
}

// Store sets the value for a key.
func (sm *SyncMap[K, V]) Store(key K, value V) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

// Delete removes a key from the map.
func (sm *SyncMap[K, V]) Delete(key K) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.m, key)
}

// Range calls the given function for each key and value present in the map.
func (sm *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	for k, v := range sm.m {
		if !f(k, v) {
			break
		}
	}
}

// Reset renew empty map.
func (sm *SyncMap[K, V]) Reset() {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m = make(map[K]V)
}

// Clear removes all keys from the map.
func (sm *SyncMap[K, V]) Clear() {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	for k, _ := range sm.m {
		delete(sm.m, k)
	}
}
