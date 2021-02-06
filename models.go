package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID int `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
}

type Users []User

func (u Users) JSON() (string, error) {
	json, err := json.Marshal(u)
	if err != nil {
		return "", fmt.Errorf("Users.JSON(): error marshaling JSON: %v", err)
	}
	return string(json), nil
}
