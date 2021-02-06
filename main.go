package main

import (
	"log"
	"net/http"
)

const port = "8081"

func main() {
	userAPI := NewUserAPI("http://127.0.0.1", "8080")
	router := NewRouter(userAPI)

	log.Printf("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}
