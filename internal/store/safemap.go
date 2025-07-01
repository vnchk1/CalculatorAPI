package store

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

var (
	storageInstance *SafeMap
	storageOnce     sync.Once
)

type SafeMap struct {
	mu sync.RWMutex
	m  map[uuid.UUID]int
}

type SafeMaps struct {
	storage *SafeMap
}

func NewSafeMap() *SafeMap {
	return &SafeMap{m: make(map[uuid.UUID]int)}
}

func GetStorage() *SafeMap {
	storageOnce.Do(func() {
		storageInstance = NewSafeMap()
	})
	return storageInstance
}

func (s *SafeMap) MapSet(key uuid.UUID, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func (s *SafeMap) MapGet(key uuid.UUID) (int, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.m[key]
	return value, ok
}

func (s *SafeMap) MapDelete(key uuid.UUID) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.m, key)
}

func (s *SafeMap) MapGetAll() map[uuid.UUID]int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	mapCopy := make(map[uuid.UUID]int, len(s.m))
	for key, value := range s.m {
		mapCopy[key] = value
	}
	return mapCopy
}

func (s *SafeMap) MapListAll() {
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Printf("List of all maps: \n")
	for key, value := range s.m {
		fmt.Printf("%v : %v \n", key, value)
	}
}
