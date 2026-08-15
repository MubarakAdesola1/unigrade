package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", showForm)
	http.HandleFunc("/calculate", calculate)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("UniGrade running on port:", port)
	http.ListenAndServe(":"+port, nil)
}
