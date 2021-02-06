package main

import (
	"log"
	"sync"
	"time"
)

type UserCache struct {
	createdAt time.Time
	users Users
}

func (uc UserCache) Expired(now time.Time, expiry time.Duration) bool {
	return uc.createdAt.Add(expiry).Before(now)
}

type Storage struct {
	cache map[string]UserCache
	cacheExpiry time.Duration
	mu sync.Mutex
	fetcher userGetter
}

func NewStorage(fetcher userGetter, cacheExpiry time.Duration) Storage {
	return Storage{
		cache: make(map[string]UserCache),
		cacheExpiry: cacheExpiry,
		fetcher: fetcher,
	}
}

func (s Storage) Search(term string) (Users, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	userCache, ok := s.cache[term]
	if ok && !userCache.Expired(time.Now(), s.cacheExpiry) {
		log.Printf("using cache for term %q", term)
		return userCache.users, nil
	}

	log.Printf("fetching data for term %q", term)
	users, err := s.fetcher.Search(term)
	if err != nil {
		return nil, err
	}

	if term != "" {
		users = users.Refine(term)
	}

	s.cache[term] = UserCache{
		createdAt: time.Now(),
		users: users,
	}

	return users, nil
}
