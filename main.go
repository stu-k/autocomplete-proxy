package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const port = "8081"

func main() {
	userAPI := NewUserAPI("http://127.0.0.1", "8080")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/status", StatusController)

	router.HandleFunc("/users", UsersController(userAPI)).
		Methods(http.MethodGet)

	log.Printf("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}
