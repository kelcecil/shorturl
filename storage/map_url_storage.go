package storage

import (
	"errors"
	"sync"
)

// MapURLStorage uses a standard Go map to
// store the URL.
type MapURLStorage struct {
	storage           map[int]string
	lock              sync.RWMutex
	currentIdentifier int
}

func MakeMapURLStorage() *MapURLStorage {
	return &MapURLStorage{
		storage:           make(map[int]string),
		lock:              sync.RWMutex{},
		currentIdentifier: 0,
	}
}

func (s *MapURLStorage) Add(url string) (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	identifier := s.currentIdentifier

	if _, exists := s.storage[identifier]; !exists {
		s.storage[identifier] = url
		s.currentIdentifier = s.currentIdentifier + 1
		return identifier, nil
	}
	return 0, errors.New("Identifier already exists")
}

func (s *MapURLStorage) Delete(identifier int) error {
	s.lock.Lock()
	delete(s.storage, identifier)
	s.lock.Unlock()
	return nil
}

func (s *MapURLStorage) Get(identifier int) (string, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if value, ok := s.storage[identifier]; ok {
		return value, nil
	}
	return "", errors.New("Identifier does not exist")

}
