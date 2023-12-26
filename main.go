package main

import (
	"fmt"
	"net/http"
)

const PORT = ":8080"

type User struct {
	name                string
	age                 uint16
	money               int16
	avgGrades, happines float64
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s He is %d and he"+
		"has money: %d", u.name, u.age, u.money)
}

func (u *User) setNewName(newName string) {
	u.name = newName
}

func homePage(page http.ResponseWriter, r *http.Request) {
	bob := User{"boh", 25, -50, 3.0, 5}
	bob.setNewName("Alex")
	fmt.Fprintf(page, bob.getAllInfo())
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
