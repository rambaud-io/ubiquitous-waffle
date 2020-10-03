package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/hacktoberfest", func(w http.ResponseWriter, r *http.Request) {
		log.Print(r)
		b, err := fmt.Fprintf(w, "I am just here for the t-shirt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
		}
		log.Printf("%d bytes written", b)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r * http.Request) {
		log.Print(r)
		b, err := fmt.Fprintf(w, "OK")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
		}
		log.Printf("%d bytes written", b)
	})

	log.Print("Server is now listening on :8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}