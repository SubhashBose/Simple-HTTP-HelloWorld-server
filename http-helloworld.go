package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define port flag with default value of 8080
	text := flag.Int("text", "Hello, World!", "Text body to serve")
	port := flag.Int("port", 8080, "port to run the server on")
	flag.Parse()

	// Define handler for root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, *text)
	})

	// Start server on specified port
	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Server starting on port %d", *port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}