package main

import (
	"fmt"
	"net/http"
)

type database map[string]int

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/ruiters":
		fmt.Fprintf(w, "Ruiters: %d\n", db["apple"])
	case "/washingons-post":
		fmt.Fprintf(w, "Bananas: %d\n", db["banana"])
	case "/twitter":
		fmt.Fprintf(w, "Oranges: %d\n", db["orange"])
	default:
		http.NotFound(w, r)
	}

}

func main() {
	db := database{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}

	http.ListenAndServe("localhost:8000", db)
}
