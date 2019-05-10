package main

import (
	"net/http"
	"fmt"
	// "html"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s", r.URL.Path[1:])
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error");
	}
}