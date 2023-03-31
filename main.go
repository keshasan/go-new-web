package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/users", usersPage)

	port := ":8080"
	println("Server listens on port : ", port)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal("Http error", err)
	}
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsFired   bool
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	//user := User{"Vasya", "Ivanov"}

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func usersPage(w http.ResponseWriter, r *http.Request) {
	users := []User{User{"John", "Doe", true}, User{"Jane", "Doe", false}}

	tmpl, err := template.ParseFiles("static/users.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	if err := tmpl.Execute(w, users); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}
