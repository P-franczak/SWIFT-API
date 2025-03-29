package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Prosta funkcja obsługująca requesty na głównym endpointcie
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, SWIFT API is running!")
	})

	// Start serwera na porcie 8080
	fmt.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
