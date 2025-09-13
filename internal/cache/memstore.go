package cache

import (
	"log"
	"sync"
	"time"

	"github.com/wsb777/involta-test/internal/models"
)

type Value struct {
	Person     *models.Person
	Expiration int64
}

type MemStore struct {
	values map[int]Value
	mu     sync.RWMutex
	stop   chan struct{}
}

func NewMemStore(cleanupInterval time.Duration) *MemStore {
	m := &MemStore{
		values: make(map[int]Value),
		stop:   make(chan struct{}),
	}
	go m.cleanupLoop(cleanupInterval)
	return m
}

func (m *MemStore) cleanup() {
	m.mu.Lock()
	defer m.mu.Unlock()
	log.Print("[INFO] Start clean memstore")
	now := time.Now().UnixNano()
	for k, v := range m.values {
		if now > v.Expiration {
			delete(m.values, k)
		}
	}
}

func (m *MemStore) Set(key int, value *models.Person) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.values[key] = Value{
		Person:     value,
		Expiration: time.Now().Add(15 * time.Minute).UnixNano(),
	}
}

func (m *MemStore) Get(key int) (*models.Person, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, exist := m.values[key]
	if !exist || time.Now().UnixNano() > value.Expiration {
		return nil, false
	}
	return value.Person, true
}

func (m *MemStore) Delete(key int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, exist := m.values[key]
	if exist {
		delete(m.values, key)
		return true
	}
	return false
}

func (m *MemStore) cleanupLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			m.cleanup()
		case <-m.stop:
			return
		}
	}
}

func (m *MemStore) Stop() {
	close(m.stop)
}
