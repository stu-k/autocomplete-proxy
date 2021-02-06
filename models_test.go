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
			},
			want: Users{
				{ID: 2, Name: "Ray Bradbery", Email: "hotbooks@readingsux.edu"},
			},
		},
		{
			name: "prioritizes full name matches first",
			term: "Brad",
			users: Users{
				{ID: 1, Name: "Ray Bradbery", Email: "hotbooks@readingsux.edu"},
				{ID: 2, Name: "Brad Pitt", Email: "joeblack@reepr.gov"},
			},
			want: Users{
				{ID: 2, Name: "Brad Pitt", Email: "joeblack@reepr.gov"},
				{ID: 1, Name: "Ray Bradbery", Email: "hotbooks@readingsux.edu"},
			},
		},
		{
			name: "prioritizes email matches second",
			term: "Brad",
			users: Users{
				{ID: 1, Name: "Ray Bradbery", Email: "hotbooks@readingsux.edu"},
				{ID: 2, Name: "Brad Pitt", Email: "joeblack@reepr.gov"},
				{ID: 3, Name: "Marsha B", Email: "mbrady@boojan.com"},
			},
			want: Users{
				{ID: 2, Name: "Brad Pitt", Email: "joeblack@reepr.gov"},
				{ID: 3, Name: "Marsha B", Email: "mbrady@boojan.com"},
				{ID: 1, Name: "Ray Bradbery", Email: "hotbooks@readingsux.edu"},
			},
		},
		{
			name: "removes and sorts correctly",
			term: "Brad",
			users: Users{
				{ID: 1, Name: "Aldous Huxley", Email: "happyhappy@soma.org"},
				{ID: 2, Name: "Ray Bradbery", Email: "hotbooks@readingsux.edu"},
				{ID: 3, Name: "George Orwell", Email: "bbfan@peepme.gov"},
				{ID: 4, Name: "Brad Pitt", Email: "joeblack@reepr.gov"},
				{ID: 5, Name: "Marsha B", Email: "mbrady@boojan.com"},
			},
			want: Users{
				{ID: 4, Name: "Brad Pitt", Email: "joeblack@reepr.gov"},
				{ID: 5, Name: "Marsha B", Email: "mbrady@boojan.com"},
				{ID: 2, Name: "Ray Bradbery", Email: "hotbooks@readingsux.edu"},
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
