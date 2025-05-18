package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	// Define port flag with default value of 8080
	text := flag.String("text", "Hello, World!", "Text body to serve")
	addr := flag.String("addr", "8080", "Address to run the server on. This can a full address in IP:PORT format or just a PORT listening to all adresses.")
	logging := flag.Bool("log", false, "Enable logging of all incoming hits.")
	flag.Parse()

	// Define handler for root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if(logging) {
			log.Printf("Received hit from %s", r.RemoteAddr)
		}
		w.Header().Set("Server", "HTTP HelloWorld Server by S. Bose")
		fmt.Fprintf(w, *text)
	})

	// Start server on specified port
	if(strings.Index(*addr, ":") < 0){
		*addr = ":" + *addr
	}
	log.Printf("Server starting on address %s", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}