package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{store: map[string]int{}}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
