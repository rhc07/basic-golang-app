package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValue(100, 100)

	fmt.Fprintf(w, fmt.Sprintf("This is the about page and I am %d cm tall", sum))
}

func Euros(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world, it's coming home ")
}

func addValue(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {

	f, err := divideValue(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by zero")
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))

}

func divideValue(x, y float32) (float32, error) {
	if y <= x {
		err := errors.New("Cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil

}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/euros", Euros)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting on port number: %s", portNumber))

	http.ListenAndServe(portNumber, nil)

	// fmt.Println("Hello world!")

	// handle func sets up the routes
	// http.responsewriter needs to write resopnses to the user.
	// *pointer is an address in memory.
	// Handle func needs to send a request to an IP address/any address in memory for a successful connection

}