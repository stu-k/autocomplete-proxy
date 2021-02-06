package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetUsers(term string) (string, error) {
	r, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080", nil)
	if err != nil {
		return "", fmt.Errorf("get users: error creating request: %v", err)
	}

	log.Printf("searching for term %q", term)

	q := r.URL.Query()
	q.Add("search", term)
	r.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return "", fmt.Errorf("get users: error making request: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("get users: error reading request body: %v", err)
	}

	return string(body), nil
}
