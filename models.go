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
	termL := strings.ToLower(term)
	for _, user := range(u) {
		nameL := strings.ToLower(user.Name)
		if strings.Contains(nameL, termL) || strings.Contains(user.Email, termL) {
			result = append(result, user)
		}
	}
	return result
}

func (u Users) JSON() (string, error) {
	json, err := json.Marshal(u)
	if err != nil {
		return "", fmt.Errorf("Users.JSON(): error marshaling JSON: %v", err)
	}
	return string(json), nil
}
