package storage

import (
	"errors"
	"sync"
)

// MapURLStorage uses a standard Go map to
// store the URL.
type MapURLStorage struct {
	storage map[string]string
	lock    sync.RWMutex
}

func MakeMapURLStorage() *MapURLStorage {
	return &MapURLStorage{
		storage: make(map[string]string),
		lock:    sync.RWMutex{},
	}
}

func (s *MapURLStorage) Add(identifier, url string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, exists := s.storage[identifier]; !exists {
		s.storage[identifier] = url
		return nil
	}
	return errors.New("Identifier already exists")
}

func (s *MapURLStorage) Delete(identifier string) error {
	s.lock.Lock()
	delete(s.storage, identifier)
	s.lock.Unlock()
	return nil
}

func (s *MapURLStorage) Get(identifier string) (string, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if value, ok := s.storage[identifier]; ok {
		return value, nil
	}
	return "", errors.New("Identifier does not exist")

}
