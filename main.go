package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Flux 333 ")
	})
	log.Fatalf("error: %s", http.ListenAndServe(":8080", nil))
}
