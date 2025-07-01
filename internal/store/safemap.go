package store

import (
	"fmt"
	"sync"
)

var (
	storageInstance *SafeMap
	storageOnce     sync.Once
)

type SafeMap struct {
	mu sync.RWMutex
	m  map[string]int
}

type SafeMaps struct {
	storage *SafeMap
}

func NewSafeMap() *SafeMap {
	return &SafeMap{m: make(map[string]int)}
}

func GetStorage() *SafeMap {
	storageOnce.Do(func() {
		storageInstance = NewSafeMap()
	})
	return storageInstance
}

func (s *SafeMap) MapSet(key string, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func (s *SafeMap) MapGet(key string) (int, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.m[key]
	return value, ok
}

func (s *SafeMap) MapDelete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.m, key)
}

func (s *SafeMap) MapGetAll() map[string]int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	mapCopy := make(map[string]int, len(s.m))
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
