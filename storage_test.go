package main

import (
	"testing"
	"time"
)

var userSet1 = Users{
	{ID: 1, Name: "Aldous Huxley", Email: "happyhappy@soma.org"},
	{ID: 2, Name: "Ray Bradbery", Email: "hotbooks@readingsux.edu"},
	{ID: 3, Name: "George Orwell", Email: "bbfan@peepme.gov"},
}

var userSet2 = Users{
	{ID: 4, Name: "Brad Pitt", Email: "joeblack@reepr.gov"},
	{ID: 5, Name: "Marsha B", Email: "mbrady@boojan.com"},
}

type mockUserGetter struct {
	used bool
	userSet Users
}

func NewMockUserGetter() *mockUserGetter {
	return &mockUserGetter{
		userSet: userSet1,
	}
}

func (mock *mockUserGetter) Search(_ string) (Users, error) {
	mock.used = true
	return mock.userSet, nil
}

func TestStorageSearch(t *testing.T) {
	tt := []struct{
		name, term string
		cacheExpiry time.Duration
		users Users
		fetcherUsed, injectCache bool
	}{
		{
			name: "fetches when data not cached",
			term: "test",
			cacheExpiry: 0 * time.Second,
			users: userSet1,
			fetcherUsed: true,
			injectCache: false,
		},
		{
			name: "uses cached data when present",
			term: "test",
			cacheExpiry: 9999 * time.Hour,
			users: userSet1,
			fetcherUsed: false,
			injectCache: true,
		},
	}

	for _, tc := range(tt) {
		t.Run(tc.name, func(t *testing.T) {
			mockFetcher := NewMockUserGetter()
			storage := NewStorage(mockFetcher, tc.cacheExpiry)
			if tc.injectCache {
				storage.cache[tc.term] = UserCache{
					createdAt: time.Now(),
					users: tc.users,
				}
			}

			_, err := storage.Search(tc.term)
			if err != nil {
				t.Errorf("wanted no error, got %v", err)
			}

			if mockFetcher.used != tc.fetcherUsed {
				t.Errorf("wanted fetcher use to == %v, got %v", tc.fetcherUsed, mockFetcher.used)
			}
		})
	}
}
