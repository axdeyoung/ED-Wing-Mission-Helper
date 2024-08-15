package main

import (
	"log"
	"net/http"
)

func main() {
	// Path to the React build directory
	fs := http.FileServer(http.Dir("../website/build"))

	// Handle requests to the root path
	http.Handle("/", fs)

	// Start the server on port 8080
	log.Println("Serving on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
