package main

import (
	"log"
	"net/http"
	"time"
)

const port = "8081"

func main() {
	userAPI := NewUserAPI("http://127.0.0.1", "8080")
	storage := NewStorage(userAPI, 15 * time.Second)
	router := NewRouter(storage)

	log.Printf("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}
