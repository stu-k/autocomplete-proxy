package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type User struct {
	ID int `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
}

type Users []User

func (u Users) Refine(term string) (result Users) {
	otherMatches := Users{}

	termL := strings.ToLower(term)
	for _, user := range(u) {
		nameL := strings.ToLower(user.Name)
		for _, namePart := range(strings.Split(nameL, " ")) {
			if namePart == termL {
				result = append(result, user)
				break
			}
		}

		// Prevent a user with a name match from being added again
		if len(result) > 0 && result[len(result) - 1].ID == user.ID {
			continue
		}

		if strings.Contains(nameL, termL) || strings.Contains(user.Email, termL) {
			otherMatches = append(otherMatches, user)
		}
	}

	return append(result, otherMatches...)
}

func (u Users) JSON() (string, error) {
	json, err := json.Marshal(u)
	if err != nil {
		return "", fmt.Errorf("Users.JSON(): error marshaling JSON: %v", err)
	}
	return string(json), nil
}
