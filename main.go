package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/euros", Euros)

	fmt.Println(fmt.Sprintf("Starting on port number: %s", portNumber))
	http.ListenAndServe(portNumber, nil)

	// fmt.Println("Hello world!")

	// handle func sets up the routes
	// http.responsewriter needs to write resopnses to the user.
	// *pointer is an address in memory.
	// Handle func needs to send a request to an IP address/any address in memory for a successful connection

}
