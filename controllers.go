package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func StatusController(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ok")
}

func UsersController(w http.ResponseWriter, req *http.Request) {
	queries, ok := req.URL.Query()["search"]
	if !ok {
		queries = []string{""}
	}

	r, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080", nil)
	if err != nil {
		log.Printf("users controller: error creating request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("searching for term %q", queries[0])

	q := r.URL.Query()
	q.Add("search", queries[0])
	r.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Printf("users controller: error making request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("users controller: error reading request body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(body))
}
