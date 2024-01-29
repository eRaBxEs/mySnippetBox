package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	// create a file server which serves files out of the "./ui/static/" directory
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// due to the behaviour of th http.FileServer() which removes the leading slash from the given URL path
	// this will lead to searching for ./ui/static which is a path that does not exist
	// so, we use http.StripPrefix() to then strip the leading "/static" before registering it to mux
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// register the other application routes as normal
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
