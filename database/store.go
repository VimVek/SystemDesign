package database

import (
	"errors"
	"sync"

	"github.com/vimvek/urlshortner/models"
)

// Store represents an in-memory database.
type Store struct {
	mu   sync.RWMutex
	data map[string]string
}

// NewStore creates a new in-memory store.
func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

// SaveURL saves a URL in the store.
func (s *Store) SaveURL(url models.URL) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[url.ShortURL] = url.LongURL
}

// GetURL retrieves a URL from the store.
func (s *Store) GetURL(shortURL string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	longURL, ok := s.data[shortURL]
	if !ok {
		return "", errors.New("url not found")
	}

	return longURL, nil
}
func (s *Store) GetAllURLs() map[string]string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	allURLs := make(map[string]string)
	for shortURL, longURL := range s.data {
		allURLs[shortURL] = longURL
	}

	return allURLs
}
