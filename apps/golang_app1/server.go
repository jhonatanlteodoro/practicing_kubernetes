package main

import (
	"log"
	"net/http"
)

func JsonHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(
		[]byte(`{"hello": "I'm a json response", "from": "app go 1"}`),
	)
}

func main() {
	http.HandleFunc("/", JsonHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
