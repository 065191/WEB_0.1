package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {

		name := r.URL.Query().Get("name")
		age := r.URL.Query().Get("age")
		nick := r.URL.Query().Get("nick")
		fmt.Fprintf(w, "Имя: %s \nВозраст: %s \nnick: %s\n", name, age, nick)
	})
	fmt.Println("Server is listening...")
	http.ListenAndServe(":88", nil)
}
