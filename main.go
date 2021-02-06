package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const port = "8081"

func statusController(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ok")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/status", statusController)

	log.Printf("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}
