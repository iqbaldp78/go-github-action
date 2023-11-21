package main

import "net/http"

// Create simple endpoint
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello, world!"))
		if err != nil {
			panic(err)
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

// Run the server
// go run main.go
