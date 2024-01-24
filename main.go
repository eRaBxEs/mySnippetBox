package main

import (
	"log"
	"net/http"
)

// define a home handler function which writes a byte slice containing "Hello from snippet box"
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippet box"))
}

func main() {
	// use the http.NewServeMux() to initialize a new servemux
	mux := http.NewServeMux()
	// then register the home function as a handler for the URL pattern "/"
	mux.HandleFunc("/", home)

	// use the http.ListenAndServe() function to start a new server
	// we pass two parameters, the first is the tcp network address:4000
	// next we pass the servemux varaiable that we have initializeds
	// if http.ListenAndServe() returns an error then we log the error using log.fatal(err) which stops the execution of the program
	// note that error returned by http.ListenAndServe() is a non nil value else it is always nil
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
