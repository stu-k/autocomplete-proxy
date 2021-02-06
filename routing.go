package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(userAPI userGetter) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/status", StatusController)

	router.HandleFunc("/users", UsersController(userAPI)).
		Methods(http.MethodGet).
		HeadersRegexp("Content-Type", "application/json")

	return router
}
