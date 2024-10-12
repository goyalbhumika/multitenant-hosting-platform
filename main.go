package main

import (
	"fmt"
	"net/http"
)

// helloHandler responds with "Hello, World!" to HTTP requests.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	// Register the helloHandler function for the root path "/"
	http.HandleFunc("/", helloHandler)

	// Start the HTTP server on port 8080
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
