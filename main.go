package main

import (
	"fmt"
	"net/http"

	"github.com/P-franczak/swift-api/parser"
)

func main() {

	parse_err := parser.ParseSwiftCSV("Interns_2025_SWIFT_CODES - Sheet1.csv") // Plik CSV z danymi
	if parse_err != nil {
		fmt.Println("Error parsing CSV:", parse_err)
	}

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
