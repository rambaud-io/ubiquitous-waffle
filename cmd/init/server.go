package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hacktoberfest", func(w http.ResponseWriter, r *http.Request) {
		// todo add request logging
		fmt.Fprintf(w, "I am just here for the t-shirt")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r * http.Request) {
		// todo add request logging
		fmt.Fprintf(w, "OK")
	})


	if err := http.ListenAndServe(":8080", nil); err != nil {
		// todo add logging if failed
	}
}