package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type UserAPI struct {
	Port, URL string
}

func NewUserAPI(url, port string) UserAPI {
	return UserAPI{
		Port: port,
		URL: url,
	}
}

func (api UserAPI) Search(term string) (string, error) {
	uri := fmt.Sprintf("%s:%s", api.URL, api.Port)
	r, err := http.NewRequest(http.MethodGet, uri, nil)
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
