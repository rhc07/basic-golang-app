package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home_page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about_page.html")
}

func Euros(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "euros_page.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}

}

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
