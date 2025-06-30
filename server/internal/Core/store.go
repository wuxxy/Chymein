package Core

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

type Store[T any] struct {
	mu      sync.RWMutex
	data    T
	path    string
	saveCh  chan struct{}
	closing chan struct{}
}

// Load initializes the store and starts the save worker
func Load[T any](path string, fallback T) (*Store[T], error) {
	store := &Store[T]{
		data:    fallback,
		path:    path,
		saveCh:  make(chan struct{}, 1),
		closing: make(chan struct{}),
	}

	if f, err := os.ReadFile(path); err == nil {
		_ = json.Unmarshal(f, &store.data)
	}

	go store.saveWorker()

	return store, nil
}

// View provides safe read-only access to the data
func (s *Store[T]) View(fn func(data *T)) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	fn(&s.data)
}

// GetPtr returns a pointer to the data (do not mutate unless you know what you're doing)
func (s *Store[T]) GetPtr() *T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return &s.data
}

// Update applies a mutation function and queues a save
func (s *Store[T]) Update(mutator func(*T)) {
	s.mu.Lock()
	mutator(&s.data)
	s.mu.Unlock()
	s.QueueSave()
}

// QueueSave schedules a background save (debounced)
func (s *Store[T]) QueueSave() {
	select {
	case s.saveCh <- struct{}{}:
	default:
		// already queued
	}
}

// Save immediately flushes to disk (blocking)
func (s *Store[T]) Save() error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	b, err := json.MarshalIndent(s.data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.path, b, 0644)
}

// saveWorker debounces and serializes save operations
func (s *Store[T]) saveWorker() {
	for {
		select {
		case <-s.saveCh:
			time.Sleep(50 * time.Millisecond) // debounce delay
			_ = s.Save()
		case <-s.closing:
			return
		}
	}
}

// Close flushes final state and shuts down the save worker
func (s *Store[T]) Close() error {
	close(s.closing)
	return s.Save()
}
