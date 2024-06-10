package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	/*
		> go run main.go --> will start the server
		> curl -v localhost:9090 --> will match path with "/"
		> curl -v localhost:9090/goodbye --> will match path with "/goodbye"
		> curl -v localhost:9090/anything --> will match path with "/" because it's a greedy match
		> curl -v -d 'sadi' localhost:9090 --> pass data 'sadi' to the server
	*/

	//vim ~/zshrc
	//go.dev servemux link
	//https://go.dev/src/net/http/server.go?s=61509%3A61556#L2378

	//registers a function to a path "on a thing" called default servemux
	//defualt servemux is an http handler
	//Servemux will determine which handler will be executed by analyzing pattern of the path
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World!!!")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading body", err)
			http.Error(rw, "Ooops, wrong window may be!", http.StatusBadRequest)
			return
		}
		log.Printf("Data %s\n", d)

		fmt.Fprintf(rw, "Hello %s !!! \n ", d)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World!!!")
	})

	//constructs an http server and registers a defualt handler to it
	//second param is a --> handler.
	//If nil, then it uses "default servemux"
	//which means, if u do not give any handler, it will call HandleFunc() function
	//(I need more clarification on this, for now, I am good)
	//":9090" --> port 9090 of every IP

	// ServeMux is an HTTP request multiplexer.
	// It matches the URL of each incoming request against a list of registered
	// patterns and calls the handler for the pattern that
	// most closely matches the URL.

	// Listen for connections on all ip addresses (0.0.0.0)
	// port 9090
	log.Println("Starting Server")
	err := http.ListenAndServe(":9090", nil)
	log.Fatal(err)
}
