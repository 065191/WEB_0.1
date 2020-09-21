package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ViewData struct {
	Title string
	Users []User
}
type User struct {
	Name string
	Age  int
}

func main() {

	data := ViewData{
		Title: "Users List",
		Users: []User{
			User{Name: "Tom", Age: 21},
			User{Name: "Kate", Age: 23},
			User{Name: "Alice", Age: 25},
		},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl, _ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, data)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
