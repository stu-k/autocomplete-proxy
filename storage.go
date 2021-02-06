package main

import (
	"log"
	"sync"
)

type Storage struct {
	cache map[string]Users
	mu sync.Mutex
	fetcher userGetter
}

func NewStorage(fetcher userGetter) Storage {
	return Storage{
		cache: make(map[string]Users),
		fetcher: fetcher,
	}
}

func (s Storage) Search(term string) (Users, error) {
	if users, ok := s.cache[term]; ok {
		log.Printf("using cache for term %q", term)
		return users, nil
	}

	log.Printf("fetching data for term %q", term)
	users, err := s.fetcher.Search(term)
	if err != nil {
		return nil, err
	}

	if term != "" {
		users = users.Refine(term)
	}

	s.cache[term] = users

	return users, nil
}
