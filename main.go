package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// define a home handler function which writes a byte slice containing "Hello from snippet box"
func home(w http.ResponseWriter, r *http.Request) {
	// add check to catch 404 missing page which manages the organic behavior of servemux that interprets "/" as a wild card
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from snippet box"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	// extract the value of the id parameter from the query string
	// convert the string id to integer and check that the value is a positive integer else return not found
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// use the fmt.Fprintf() to interpolate the id with our response as below
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// check if the http Method is POST or not
	if r.Method != http.MethodPost {
		// if not use the http.Error() function to pass the the error message Method not allowed, status code 405 via http.StatusMethodNotAllowed
		// then return the function to stop further execution of the function
		w.Header().Set("Allow", http.MethodPost) // customize header to set which header are allowed for this particular URL
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// use the http.NewServeMux() to initialize a new servemux
	mux := http.NewServeMux()
	// then register the home function as a handler for the URL pattern "/"
	mux.HandleFunc("/", home)
	// register the two new handlers
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// use the http.ListenAndServe() function to start a new server
	// we pass two parameters, the first is the tcp network address:4000
	// next we pass the servemux varaiable that we have initializeds
	// if http.ListenAndServe() returns an error then we log the error using log.Fatal(err) which stops the execution of the program
	// note that error returned by http.ListenAndServe() is a non nil value incase of an error else it is nil
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
