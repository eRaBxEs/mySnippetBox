package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from SnippetBox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Show a specific snippet ..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet ..."))
}

func main() {
	// Register the two new handler functions and corresponding URL patterns with
	// the servemux, in exactly the same way that we did before.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)                        // sub-tree path: ends with a trailing '/'
	mux.HandleFunc("/snippet", showSnippet)          // fixed path: does not ends with a trailing '/'
	mux.HandleFunc("/snippet/create", createSnippet) // fixed path: does not ends with a trailing '/'

	log.Println("Started server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
