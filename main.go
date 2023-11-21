package main

import "net/http"

// Create simple endpoint
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	http.ListenAndServe(":8080", nil)
}

// Run the server
// go run main.go
