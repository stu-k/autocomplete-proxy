package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UserAPI struct {
	Port, URL string
}

func NewUserAPI(url, port string) UserAPI {
	return UserAPI{
		Port: port,
		URL:  url,
	}
}

func (api UserAPI) Search(term string) (users Users, err error) {
	uri := fmt.Sprintf("%s:%s", api.URL, api.Port)
	r, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("UserAPI.Search: error creating request: %v", err)
	}

	q := r.URL.Query()
	q.Add("search", term)
	r.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("UserAPI.Search: error making request: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("UserAPI.Search: error reading request body: %v", err)
	}

	if err := json.Unmarshal(body, &users); err != nil {
		return nil, fmt.Errorf("UserAPI.Search: error unmarshaling JSON: %v", err)
	}

	return users, nil
}
