package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

//Packages needed
func main() {
	//setting up a http request for you plus register it
	http.HandleFunc("/you", func(w http.ResponseWriter, r *http.Request) {

		// when request is made below message is printed
		fmt.Fprintf(w, "Hello person")

	})
	//printout of server with time shown
	log.Println("Starting server..")

	//server created at the localhost port usimg tcp
	listener, err := net.Listen("tcp", ":8080")

	//recieve error record and exit
	if err != nil {
		log.Fatal(err)
	}
	// setup HTTP connection for the listener of the server
	http.Serve(listener, nil)
}
