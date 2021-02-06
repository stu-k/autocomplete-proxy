package main

import (
	"testing"
	"reflect"
)

func TestUsersRefine(t *testing.T) {
	tt := []struct{
		name, term string
		users, want Users
	}{
		{
			name: "removes irrelevant results",
			term: "Brad",
			users: Users{
				{ID: 1, Name: "Aldous Huxley", Email: "happyhappy@soma.org"},
				{ID: 2, Name: "Ray Bradbery", Email: "hotbooks@readingsux.edu"},
				{ID: 3, Name: "George Orwell", Email: "bbfan@peepme.gov"},
				{ID: 4, Name: "Marsha B", Email: "mbrady@jansux.com"},
			},
			want: Users{
				{ID: 2, Name: "Ray Bradbery", Email: "hotbooks@readingsux.edu"},
				{ID: 4, Name: "Marsha B", Email: "mbrady@jansux.com"},
			},
		},
	}

	for _, tc := range(tt) {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.users.Refine(tc.term)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
