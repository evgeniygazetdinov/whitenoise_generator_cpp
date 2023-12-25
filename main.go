package main

import (
	"fmt"
	"net/http"
)

const PORT = ":8080"

func homePage(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Go rulit")
}

func contactsPage(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "contacts page")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactsPage)
	fmt.Printf("Running on %s \n", PORT)
	http.ListenAndServe(PORT, nil)
}

func main() {
	handleRequest()
}
