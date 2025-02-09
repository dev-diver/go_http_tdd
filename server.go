package main

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	if player == "Pepper" {
		fmt.Fprintf(w, "20")
		return
	}

	if player == "Floyd" {
		fmt.Fprintf(w, "10")
		return
	}

}
