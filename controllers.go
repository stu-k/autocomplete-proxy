package main

import (
	"fmt"
	"log"
	"net/http"
)

func StatusController(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ok")
}

func UsersController(userAPI userGetter) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		queries, ok := req.URL.Query()["search"]
		if !ok {
			queries = []string{""}
		}

		usersJSON, err := userAPI.Search(queries[0])
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, usersJSON)
	}
}

type userGetter interface {
	Search(term string) (string, error)
}
